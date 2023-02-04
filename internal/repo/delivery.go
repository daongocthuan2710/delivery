package repo

import (
	"context"
	"database/sql"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"delivery/internal/model/entity"
)

func NewDeliveryRepo(db *sql.DB) *DeliveryRepo {
	return &DeliveryRepo{
		DB: db,
	}
}

type DeliveryRepo struct {
	// logger logger.Interface
	DB *sql.DB
}

func (r *DeliveryRepo) Count(ctx context.Context, query DeliveryQuery) (int64, error) {
	q := []qm.QueryMod{
		entity.DeliveryWhere.Code.EQ(null.StringFrom(query.OrderCode)),
	}
	return entity.Deliveries(q...).Count(ctx, r.DB)
}

func (r *DeliveryRepo) Insert(ctx context.Context, delivery entity.Delivery) error {
	return delivery.Insert(ctx, r.DB, boil.Infer())
}

func (r *DeliveryRepo) UpdateBlacklist(ctx context.Context, delivery entity.Delivery, blacklist ...string) error {
	_, err := delivery.Update(ctx, r.DB, boil.Blacklist(blacklist...))
	return err
}

func (r *DeliveryRepo) UpdateWhitelist(ctx context.Context, delivery entity.Delivery, whitelist ...string) error {
	_, err := delivery.Update(ctx, r.DB, boil.Whitelist(whitelist...))
	return err
}
