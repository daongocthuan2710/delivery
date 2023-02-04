package repo

import (
	"context"

	"delivery/internal/model/entity"
)

type DeliveryQuery struct {
	OrderCode string
}

type IDelivery interface {
	Insert(ctx context.Context, delivery entity.Delivery) error
	Count(ctx context.Context, query DeliveryQuery) (int64, error)
	UpdateBlacklist(ctx context.Context, delivery entity.Delivery, blacklist ...string) error
	UpdateWhitelist(ctx context.Context, delivery entity.Delivery, whitelist ...string) error
}
