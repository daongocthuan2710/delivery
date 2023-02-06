package service

import (
	"context"

	"delivery/internal/model/req"
	"delivery/internal/model/res"
)

type IDelivery interface {
	Create(ctx context.Context, payload req.DeliveryCreate) (*res.DeliveryCreate, error)
	EstimateFee(ctx context.Context, payload req.EstimateFee) (*res.EstimateFee, error)
	UpdateStatusFromWebhook(b []byte, partnerCode, ip string) error
}
