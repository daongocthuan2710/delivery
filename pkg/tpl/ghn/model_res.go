package ghn

import "time"

type CreateOrderRes struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    OrderData `json:"data"`
}

type PreviewOrderRes struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    OrderData `json:"data"`
}

type OrderData struct {
	OrderCode            string    `json:"order_code"`
	SortCode             string    `json:"sort_code"`
	TransType            string    `json:"trans_type"`
	WardEncode           string    `json:"ward_encode"`
	DistrictEncode       string    `json:"district_encode"`
	Fee                  Fee       `json:"fee"`
	TotalFee             int64     `json:"total_fee"`
	ExpectedDeliveryTime time.Time `json:"expected_delivery_time"`
}

type Fee struct {
	MainService int64 `json:"main_service"`
	Insurance   int64 `json:"insurance"`
	StationDo   int64 `json:"station_do"`
	StationPu   int64 `json:"station_pu"`
	Return      int64 `json:"return"`
	R2S         int64 `json:"r2s"`
	Coupon      int64 `json:"coupon"`
}

type OrderWebhook struct {
	CODAmount       int         `json:"CODAmount"`
	CODTransferDate interface{} `json:"CODTransferDate"`
	ClientOrderCode string      `json:"ClientOrderCode"`
	ConvertedWeight int         `json:"ConvertedWeight"`
	Description     string      `json:"Description"`
	Fee             struct {
		Coupon      int `json:"Coupon"`
		Insurance   int `json:"Insurance"`
		MainService int `json:"MainService"`
		R2S         int `json:"R2S"`
		Return      int `json:"Return"`
		StationDO   int `json:"StationDO"`
		StationPU   int `json:"StationPU"`
	} `json:"Fee"`
	Height       int    `json:"Height"`
	Length       int    `json:"Length"`
	OrderCode    string `json:"OrderCode"`
	Reason       string `json:"Reason"`
	ReasonCode   string `json:"ReasonCode"`
	ShipperName  string `json:"ShipperName"`
	ShipperPhone string `json:"ShipperPhone"`
	Status       string `json:"Status"`
	Time         string `json:"Time"`
	TotalFee     int    `json:"TotalFee"`
	Type         string `json:"Type"`
	Warehouse    string `json:"Warehouse"`
	Weight       int    `json:"Weight"`
	Width        int    `json:"Width"`
}

type ProvinceRes struct {
	Data    []Province `json:"data"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
}

type DistrictRes struct {
	Data    []District `json:"data"`
	Code    int        `json:"code"`
	Message string     `json:"message"`
}

type WardRes struct {
	Data    []Ward `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Province struct {
	ProvinceID   int    `json:"ProvinceID"`
	ProvinceName string `json:"ProvinceName"`
	CountryID    int    `json:"CountryID"`
	Code         string `json:"Code"`
}

type District struct {
	DistrictID   int    `json:"DistrictID"`
	ProvinceID   int    `json:"ProvinceID"`
	DistrictName string `json:"DistrictName"`
	Code         string `json:"Code"`
}

type Ward struct {
	WardCode   string `json:"WardCode"`
	DistrictID int    `json:"DistrictID"`
	WardName   string `json:"WardName"`
}
