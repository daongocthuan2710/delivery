package service

import "context"

func NewLocation() *LocationSvc {
	return &LocationSvc{}
}

type LocationSvc struct {
}

func (s *LocationSvc) GetDetail(ctx context.Context, loc *LocationIDs) (*LocationDetail, error) {
	// TODO implement me
	return nil, nil
}
