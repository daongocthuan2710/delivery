package constant

const (
	DStatusWaitingToPick = "waiting_to_pick"
	DStatusPicking       = "picking"
	DStatusPicked        = "picked"
	DStatusDelayPickup   = "delay_pickup"
	DStatusPickupFailed  = "pickup_failed"

	DStatusDelivering     = "delivering"
	DStatusDelivered      = "delivered"
	DStatusDelayDelivery  = "delay_delivery"
	DStatusDeliveryFailed = "delivery_failed"

	DStatusWaitingToReturn = "waiting_to_return"
	DStatusReturning       = "returning"
	DStatusReturned        = "returned"
	DStatusReturnFailed    = "return_failed"
	DStatusDelayReturn     = "delay_return"

	DStatusCancelled    = "cancelled"
	DStatusCompensation = "compensation"
)

const (
	TPLCodeGHN = "GHN"
)

const (
	GHNServiceCodeExpress  = "GHN_EXP"
	GHNServiceCodeStandard = "GHN_STD"
	GHNServiceCodeSaving   = "GHN_SAVING"

	GHNServiceName = "Giao hàng nhanh"
)

var (
	ServiceCodeMapping = map[string]string{
		GHNServiceCodeExpress: TPLCodeGHN,
	}
)
