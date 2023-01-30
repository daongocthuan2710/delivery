package service

import (
	"context"

	"delivery/internal/model/req"
	"delivery/internal/model/res"
	"delivery/internal/repo"
)

func NewDeliveryService(repo repo.IDelivery) *DeliverySvc {
	return &DeliverySvc{
		repo: repo,
	}
}

type DeliverySvc struct {
	repo repo.IDelivery
}

func (s *DeliverySvc) EstimateFee(ctx context.Context, payload req.EstimateFee) (*res.EstimateFee, error) {
	return &res.EstimateFee{
		Services: []res.DeliveryService{
			{
				Name:         "GHN - gói tiêu chuẩn",
				Code:         "GHN_STD",
				TplCode:      "GHN",
				TotalFee:     20000,
				EstimateTime: "3 Ngày",
			},
		},
	}, nil
}

func (s *DeliverySvc) Create(ctx context.Context, payload req.DeliveryCreate) (*res.DeliveryCreate, error) {
	return &res.DeliveryCreate{
		OrderCode:    "123131",
		TrackingCode: "ssdjfhnj",
		Status:       "waiting_to_pick",
		TotalFee:     10000,
	}, nil
}
