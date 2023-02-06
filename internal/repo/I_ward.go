package repo

import (
	"context"

	"delivery/internal/model/entity"
)

type IWard interface {
	GetDetailLocation(ctx context.Context, query DetailLocQuery) (*DetailLocationResult, error)
}

type DetailLocQuery struct {
	WardID int
}

type DetailLocationResult struct {
	Ward     entity.Ward     `boil:"ward,bind"`
	District entity.District `boil:"district,bind"`
	Province entity.Province `boil:"province,bind"`
}
