package req

type DeliveryCreate struct {
	OrderCode   string        `json:"orderCode,omitempty"`
	Note        string        `json:"note,omitempty"`
	OrderValue  int64         `json:"orderValue,omitempty"`
	Cod         int64         `json:"cod,omitempty"`
	Weight      int64         `json:"weight,omitempty"`
	ServiceCode string        `json:"serviceCode,omitempty"`
	From        *DeliveryInfo `json:"from,omitempty"`
	To          *DeliveryInfo `json:"to,omitempty"`
}

type DeliveryInfo struct {
	Name       string `json:"name,omitempty"`
	Phone      string `json:"phone,omitempty"`
	Address    string `json:"address,omitempty"`
	ProvinceId int32  `json:"provinceId,omitempty"`
	DistrictId int32  `json:"districtId,omitempty"`
	WardId     int32  `json:"wardId,omitempty"`
}

type EstimateFee struct {
	OrderValue int64
	Cod        int64
	Weight     int64
	From       *DeliveryInfo
	To         *DeliveryInfo
}
