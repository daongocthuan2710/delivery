package service

import (
	"context"

	"delivery/internal/repo"
)

func NewLocation(wardRepo repo.IWard) *LocationSvc {
	return &LocationSvc{
		wardRepo: wardRepo,
	}
}

type LocationSvc struct {
	wardRepo repo.IWard
}

func (s *LocationSvc) GetDetail(ctx context.Context, loc *LocationIDs) (*LocationDetail, error) {
	data, err := s.wardRepo.GetDetailLocation(ctx, repo.DetailLocQuery{WardID: int(loc.Ward)})
	if err != nil {
		return nil, err
	}
	return &LocationDetail{
		Province: data.Province,
		District: data.District,
		Ward:     data.Ward,
	}, nil
}
