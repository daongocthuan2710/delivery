package repo

import (
	"context"

	"delivery/internal/model/entity"
)

type IDelivery interface {
	Insert(ctx context.Context, delivery entity.Delivery) error
	UpdateBlacklist(ctx context.Context, delivery entity.Delivery, blacklist ...string) error
	UpdateWhitelist(ctx context.Context, delivery entity.Delivery, whitelist ...string) error
}
