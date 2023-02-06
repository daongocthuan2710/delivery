package repo

import (
	"context"

	"delivery/internal/model/entity"
)

type DeliveryQuery struct {
	OrderCode    string
	TrackingCode string
}

type IDelivery interface {
	FindOne(ctx context.Context, query DeliveryQuery) (*entity.Delivery, error)
	Count(ctx context.Context, query DeliveryQuery) (int64, error)

	Insert(ctx context.Context, delivery *entity.Delivery) error
	UpdateBlacklist(ctx context.Context, delivery *entity.Delivery, blacklist ...string) error
	UpdateWhitelist(ctx context.Context, delivery *entity.Delivery, whitelist ...string) error
}
