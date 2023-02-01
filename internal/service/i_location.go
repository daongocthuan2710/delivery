package service

import "context"

type LocationIDs struct {
	Province int32
	District int32
	Ward     int32
}

type ILocation interface {
	GetDetail(ctx context.Context, loc *LocationIDs) (*LocationDetail, error)
}
