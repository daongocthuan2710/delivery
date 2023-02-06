package res

type DeliveryCreate struct {
	OrderCode    string `json:"orderCode"`
	TrackingCode string `json:"trackingCode"`
	Status       string `json:"status"`
	TotalFee     int64  `json:"totalFee"`
}

type EstimateFee struct {
	Services []DeliveryService `json:"services"`
}

type DeliveryService struct {
	Name         string `json:"name,omitempty"`
	Code         string `json:"code,omitempty"`
	TplCode      string `json:"tplCode,omitempty"`
	TotalFee     int64  `json:"totalFee,omitempty"`
	EstimateTime string `json:"estimateTime,omitempty"`
}
