package service

import (
	"context"
	"time"

	"delivery/internal/model/entity"
	"delivery/internal/model/req"
	"delivery/internal/model/res"
)

type LocationDetail struct {
	entity.Province
	entity.District
	entity.Ward
}

type WebhookData struct {
	TrackingCode      string
	OrderCode         string
	Status            string
	PartnerStatus     string
	PartnerStatusName string
	Reason            string
	UpdatedTime       time.Time
}

type IPartner interface {
	EstimateFee(ctx context.Context, payload req.EstimateFee, from, to *LocationDetail) ([]res.DeliveryService, error)
	CreateOrder(ctx context.Context, payload req.DeliveryCreate, from, to *LocationDetail) (*res.DeliveryCreate, error)
	GetWebhookData(b []byte) (*WebhookData, error)
}
