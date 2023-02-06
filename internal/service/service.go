package service

import (
	"delivery/internal/config"
	"delivery/internal/repo"
)

type Deps struct {
	Repos *repo.Repositories
	Cfg   *config.Config
}

type Services struct {
	Delivery IDelivery
	Location ILocation
}

func NewServices(deps *Deps) *Services {
	loc := NewLocation(deps.Repos.Ward)
	return &Services{
		Location: loc,
		Delivery: NewDeliveryService(deps.Repos.Delivery, deps.Repos.DeliveryHistory, deps.Cfg, loc),
	}
}
