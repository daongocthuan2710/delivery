package repo

import "database/sql"

type Repositories struct {
	Delivery        IDelivery
	DeliveryHistory IDeliveryHistory
	Ward            IWard
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Delivery:        NewDeliveryRepo(db),
		Ward:            NewWardRepo(db),
		DeliveryHistory: NewDeliveryHistoryRepo(db),
	}
}
