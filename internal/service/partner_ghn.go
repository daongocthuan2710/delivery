package service

import (
	"context"

	"delivery/internal/config"
	"delivery/internal/constant"
	"delivery/internal/model/req"
	"delivery/internal/model/res"
	"delivery/pkg/tpl/ghn"
	"delivery/pkg/util/timeutil"
)

const (
	GHNServiceCodeExpress  = "GHN_EXP"
	GHNServiceCodeStandard = "GHN_STD"
	GHNServiceCodeSaving   = "GHN_SAVING"

	GHNServiceName = "Giao hàng nhanh"
)

func NewGHN(cfg config.GHN, isProd bool) *PartnerGHN {
	return &PartnerGHN{
		Cfg:    cfg,
		IsProd: isProd,
	}
}

type PartnerGHN struct {
	Cfg    config.GHN
	IsProd bool
}

func (p *PartnerGHN) EstimateFee(ctx context.Context, payload req.EstimateFee, from, to *LocationDetail) ([]res.DeliveryService, error) {
	c := ghn.New(p.Cfg.Token, p.Cfg.ShopID, p.IsProd)
	ghnPayload := ghn.PreviewOrderReq{
		PaymentTypeID:   ghn.PaymentTypeSeller,
		Note:            "",
		RequiredNote:    ghn.RequiredNoteKHONGCHOXEMHANG,
		ClientOrderCode: "",
		CODAmount:       payload.COD,
		FromName:        payload.From.Name,
		FromPhone:       payload.From.Phone,
		FromAddress:     payload.From.Address,
		FromDistrictID:  from.District.GHNID.Int,
		FromWardCode:    from.Ward.GHNCode.String,
		ToName:          payload.To.Name,
		ToPhone:         payload.To.Phone,
		ToAddress:       payload.To.Address,
		ToWardCode:      to.Ward.GHNCode.String,
		ToDistrictID:    to.District.GHNID.Int,
		Content:         "Motasanpham",
		Weight:          payload.Weight,
		InsuranceValue:  payload.OrderValue,
		ServiceTypeID:   ghn.ServiceTypeStandard,
	}
	data, err := c.PreviewOrder(ctx, ghnPayload)
	if err != nil {
		return nil, err
	}

	estimateTime := timeutil.FormatInLocation(constant.TimeLayoutDDMMYYYY, data.Data.ExpectedDeliveryTime, constant.TimezoneHCM)
	return []res.DeliveryService{
		{
			Name:         GHNServiceName,
			Code:         GHNServiceCodeExpress,
			TplCode:      TPLCodeGHN,
			TotalFee:     data.Data.TotalFee,
			EstimateTime: estimateTime,
		},
	}, nil
}

func (p *PartnerGHN) CreateOrder(ctx context.Context, payload req.DeliveryCreate, from, to *LocationDetail) (*res.DeliveryCreate, error) {
	return nil, nil
}
