package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"delivery/internal/model/entity"
)

func NewWardRepo(db *sql.DB) *WardRepo {
	return &WardRepo{DB: db}
}

type WardRepo struct {
	DB *sql.DB
}

func (r *WardRepo) GetDetailLocation(ctx context.Context, query DetailLocQuery) (res *DetailLocationResult, err error) {
	res = &DetailLocationResult{}
	slt := []string{
		entity.WardTableColumns.ID,
		entity.WardTableColumns.Name,
		entity.WardTableColumns.GHNCode,
		entity.DistrictTableColumns.ID,
		entity.DistrictTableColumns.Name,
		entity.DistrictTableColumns.GHNID,
		entity.ProvinceTableColumns.ID,
		entity.ProvinceTableColumns.Name,
		entity.ProvinceTableColumns.GHNID,
	}
	cond := []qm.QueryMod{
		qm.Select(slt...),
		entity.WardWhere.ID.EQ(null.IntFrom(query.WardID)),
		qm.InnerJoin(fmt.Sprintf("%s on %s = %s", entity.TableNames.Districts, entity.DistrictTableColumns.ID, entity.WardTableColumns.DistrictID)),
		qm.InnerJoin(fmt.Sprintf("%s on %s = %s", entity.TableNames.Provinces, entity.ProvinceTableColumns.ID, entity.DistrictTableColumns.ID)),
	}
	err = entity.Wards(cond...).Bind(ctx, r.DB, res)
	return res, err
}
