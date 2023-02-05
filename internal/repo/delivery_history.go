package repo

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"delivery/internal/model/entity"
)

func NewDeliveryHistoryRepo(db *sql.DB) *DeliveryHistoryRepo {
	return &DeliveryHistoryRepo{DB: db}
}

type DeliveryHistoryRepo struct {
	DB *sql.DB
}

func (r *DeliveryHistoryRepo) Insert(ctx context.Context, h *entity.DeliveryHistory) error {
	return h.Insert(ctx, r.DB, boil.Infer())
}
