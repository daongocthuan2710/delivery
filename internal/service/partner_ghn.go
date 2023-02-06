package service

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/thoas/go-funk"

	"delivery/internal/config"
	"delivery/internal/constant"
	"delivery/internal/model/req"
	"delivery/internal/model/res"
	"delivery/pkg/tpl/ghn"
	"delivery/pkg/util/timeutil"
)

var (
	ghnServiceTypeMapping = map[string]int{
		constant.GHNServiceCodeStandard: ghn.ServiceTypeStandard,
	}
	ghnMappingStatus = map[string]string{
		"ready_to_pick":         constant.DStatusWaitingToPick,
		"picking":               constant.DStatusPicking,
		"money_collect_picking": constant.DStatusPicking,
		"picked":                constant.DStatusPicked,

		"storing":                  constant.DStatusDelivering,
		"transporting":             constant.DStatusDelivering,
		"sorting":                  constant.DStatusDelivering,
		"delivering":               constant.DStatusDelivering,
		"money_collect_delivering": constant.DStatusDelivering,
		"delivered":                constant.DStatusDelivered,
		"delivery_fail":            constant.DStatusDeliveryFailed,
		"waiting_to_return":        constant.DStatusWaitingToReturn,

		"return":              constant.DStatusReturning,
		"return_transporting": constant.DStatusReturning,
		"return_sorting":      constant.DStatusReturning,
		"returning":           constant.DStatusReturning,
		"returned":            constant.DStatusReturned,
		"return_fail":         constant.DStatusDelayReturn,

		"cancel": constant.DStatusCancelled,
		"damage": constant.DStatusCompensation,
		"lost":   constant.DStatusCompensation,
	}
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

func (p *PartnerGHN) GetWebhookData(b []byte) (*WebhookData, error) {
	var wh ghn.OrderWebhook
	if err := json.Unmarshal(b, &wh); err != nil {
		return nil, err
	}
	types := []string{"switch_status", "create"}
	if !funk.ContainsString(types, strings.ToLower(wh.Type)) {
		return nil, nil
	}
	t := timeutil.ParseISODate(wh.Time)
	return &WebhookData{
		TrackingCode:      wh.OrderCode,
		Status:            ghnMappingStatus[wh.Status],
		OrderCode:         wh.ClientOrderCode,
		PartnerStatus:     ghnMappingStatus[wh.Status],
		PartnerStatusName: wh.Status,
		Reason:            wh.Reason,
		UpdatedTime:       t,
	}, nil
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
			Name:         constant.GHNServiceName,
			Code:         constant.GHNServiceCodeExpress,
			TplCode:      constant.TPLCodeGHN,
			TotalFee:     data.Data.TotalFee,
			EstimateTime: estimateTime,
		},
	}, nil
}

func (p *PartnerGHN) CreateOrder(ctx context.Context, payload req.DeliveryCreate, from, to *LocationDetail) (*res.DeliveryCreate, error) {
	c := ghn.New(p.Cfg.Token, p.Cfg.ShopID, p.IsProd)
	ghnPayload := ghn.CreateOrderReq{
		PaymentTypeID:   ghn.PaymentTypeSeller,
		Note:            payload.Note,
		RequiredNote:    ghn.RequiredNoteKHONGCHOXEMHANG,
		ClientOrderCode: payload.OrderCode,
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
		Content:         payload.GetContent(),
		Weight:          payload.Weight,
		InsuranceValue:  payload.OrderValue,
		ServiceTypeID:   ghnServiceTypeMapping[payload.ServiceCode],
		Items:           make([]*ghn.Item, 0),
	}
	result, err := c.CreateOrder(ctx, ghnPayload)
	if err != nil {
		return nil, err
	}

	return &res.DeliveryCreate{
		OrderCode:    payload.OrderCode,
		TrackingCode: result.Data.OrderCode,
		Status:       constant.DStatusWaitingToPick,
		TotalFee:     result.Data.TotalFee,
	}, nil
}
