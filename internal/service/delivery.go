package service

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"

	"delivery/internal/config"
	"delivery/internal/constant"
	"delivery/internal/model/req"
	"delivery/internal/model/res"
	"delivery/internal/repo"
)

const (
	TPLCodeGHN = "GHN"
)

func NewDeliveryService(repo repo.IDelivery, cfg *config.Config, loc ILocation) *DeliverySvc {
	return &DeliverySvc{
		repo:     repo,
		cfg:      cfg,
		location: loc,
	}
}

type DeliverySvc struct {
	repo repo.IDelivery
	cfg  *config.Config

	location ILocation
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
	return &res.DeliveryCreate{
		OrderCode:    "123131",
		TrackingCode: "code",
		Status:       "waiting_to_pick",
		TotalFee:     10000,
	}, nil
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
	isProd := s.cfg.App.ENV == constant.ENVProd
	return []IPartner{
		NewGHN(s.cfg.GHN, isProd),
	}
}
