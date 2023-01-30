package repo

import "database/sql"

type Repositories struct {
	Delivery IDelivery
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Delivery: NewDeliveryRepo(db),
	}
}
