package service

import (
	"context"
	"sync"

	"github.com/volatiletech/null/v8"
	"golang.org/x/sync/errgroup"

	"delivery/internal/apperr"
	"delivery/internal/config"
	"delivery/internal/constant"
	"delivery/internal/model/req"
	"delivery/internal/model/res"
	"delivery/internal/repo"
)

func NewDeliveryService(repo repo.IDelivery, cfg *config.Config, loc ILocation) *DeliverySvc {
	isProd := cfg.App.ENV == constant.ENVProd
	return &DeliverySvc{
		repo:     repo,
		cfg:      cfg,
		location: loc,
		ghn:      NewGHN(cfg.GHN, isProd),
	}
}

type DeliverySvc struct {
	repo repo.IDelivery
	cfg  *config.Config

	location ILocation
	ghn      *PartnerGHN
}

func (s *DeliverySvc) EstimateFee(ctx context.Context, payload req.EstimateFee) (*res.EstimateFee, error) {
	type data struct {
		list []res.DeliveryService
		err  error
	}
	var (
		result = &res.EstimateFee{Services: make([]res.DeliveryService, 0)}
		errs   []error
		ch     = make(chan data)
		done   = make(chan bool)
	)
	partners := s.getPartners()
	if len(partners) == 0 {
		return result, nil
	}
	fromLoc, toLoc, err := s.getLocation(ctx, payload.From, payload.To)
	if err != nil {
		return nil, err
	}
	// read result via channel
	go func() {
		for value := range ch {
			if value.err != nil {
				errs = append(errs, value.err)
			} else {
				result.Services = append(result.Services, value.list...)
			}
		}
		done <- true
	}()
	var wg sync.WaitGroup
	wg.Add(len(partners))
	for _, partner := range partners {
		go func(partner IPartner) {
			defer wg.Done()
			list, err := partner.EstimateFee(ctx, payload, fromLoc, toLoc)
			ch <- data{
				list: list,
				err:  err,
			}
		}(partner)
	}
	wg.Wait()
	close(ch)
	<-done

	if len(result.Services) == 0 && len(errs) > 0 {
		return nil, errs[0]
	}

	return result, nil
}

func (s *DeliverySvc) Create(ctx context.Context, payload req.DeliveryCreate) (*res.DeliveryCreate, error) {
	// check exists first
	total, err := s.repo.Count(ctx, repo.DeliveryQuery{OrderCode: payload.OrderCode})
	if err != nil {
		return nil, err
	}
	if total > 0 {
		return nil, apperr.OrderCodeExisted
	}
	record, err := payload.ToEntity()
	if err != nil {
		return nil, err
	}
	p, err := s.getPartnerByIdentityCode(record.PartnerIdentityCode.String)
	if err != nil {
		return nil, err
	}
	from, to, err := s.getLocation(ctx, payload.From, payload.To)
	if err != nil {
		return nil, err
	}
	result, err := p.CreateOrder(ctx, payload, from, to)
	if err != nil {
		return nil, err
	}
	record.TrackingCode = null.StringFrom(result.TrackingCode)
	record.Status = null.StringFrom(result.Status)
	record.TotalFee = null.Int64From(result.TotalFee)

	if err = s.repo.Insert(ctx, *record); err != nil {
		return nil, err
	}

	return &res.DeliveryCreate{
		OrderCode:    payload.OrderCode,
		TrackingCode: result.TrackingCode,
		Status:       result.Status,
		TotalFee:     result.TotalFee,
	}, nil
}

func (s *DeliverySvc) getPartnerByIdentityCode(code string) (IPartner, error) {
	switch code {
	case constant.TPLCodeGHN:
		return s.ghn, nil
	}
	return nil, apperr.PartnerNotFound
}

func (s *DeliverySvc) getLocation(ctx context.Context, from *req.DeliveryInfo, to *req.DeliveryInfo) (
	*LocationDetail, *LocationDetail, error) {
	var (
		fromLoc, toLoc *LocationDetail
		g              errgroup.Group
	)
	// sender location
	g.Go(func() error {
		var err error
		fromLoc, err = s.location.GetDetail(ctx, &LocationIDs{
			Province: from.ProvinceId,
			District: from.DistrictId,
			Ward:     from.WardId,
		})
		return err
	})

	// receiver location
	g.Go(func() error {
		var err error
		toLoc, err = s.location.GetDetail(ctx, &LocationIDs{
			Province: to.ProvinceId,
			District: to.DistrictId,
			Ward:     to.WardId,
		})
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, nil, err
	}
	return fromLoc, toLoc, nil
}

func (s *DeliverySvc) getPartners() []IPartner {

	return []IPartner{
		s.ghn,
	}
}
