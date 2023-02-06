package repo

import (
	"context"

	"delivery/internal/model/entity"
)

type IDeliveryHistory interface {
	Insert(ctx context.Context, h *entity.DeliveryHistory) error
}
