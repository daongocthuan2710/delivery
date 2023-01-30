package service

import "delivery/internal/repo"

type Deps struct {
	Repos *repo.Repositories
}

type Services struct {
	Delivery IDelivery
}

func NewServices(deps *Deps) *Services {
	return &Services{
		Delivery: NewDeliveryService(deps.Repos.Delivery),
	}
}
