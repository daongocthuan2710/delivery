package service

import (
	"context"

	"delivery/internal/model/entity"
	"delivery/internal/model/req"
	"delivery/internal/model/res"
)

type LocationDetail struct {
	entity.Province
	entity.District
	entity.Ward
}

type IPartner interface {
	EstimateFee(ctx context.Context, payload req.EstimateFee, from, to *LocationDetail) ([]res.DeliveryService, error)
	CreateOrder(ctx context.Context, payload req.DeliveryCreate, from, to *LocationDetail) (*res.DeliveryCreate, error)
}
