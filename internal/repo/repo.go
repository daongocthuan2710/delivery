package repo

import "database/sql"

type Repositories struct {
	Delivery IDelivery
	Ward     IWard
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Delivery: NewDeliveryRepo(db),
		Ward:     NewWardRepo(db),
	}
}
