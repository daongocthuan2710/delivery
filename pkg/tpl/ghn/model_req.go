package ghn

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Weight   int64  `json:"weight"`
	Code     string `json:"code,omitempty"`
	Price    int64  `json:"price,omitempty"`
	Length   int    `json:"length,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type CreateOrderReq struct {
	PaymentTypeID   int     `json:"payment_type_id"`
	Note            string  `json:"note"`
	RequiredNote    string  `json:"required_note"`
	ClientOrderCode string  `json:"client_order_code"`
	CODAmount       float64 `json:"cod_amount"`

	FromName       string `json:"from_name"`
	FromPhone      string `json:"from_phone"`
	FromAddress    string `json:"from_address"`
	FromDistrictID int    `json:"from_district_id"`
	FromWardCode   string `json:"from_ward_code"`

	ToName       string `json:"to_name"`
	ToPhone      string `json:"to_phone"`
	ToAddress    string `json:"to_address"`
	ToWardCode   string `json:"to_ward_code"`
	ToDistrictID int    `json:"to_district_id"`

	Content string `json:"content"`
	Weight  int    `json:"weight"`
	Length  int    `json:"length,omitempty"`
	Width   int    `json:"width,omitempty"`
	Height  int    `json:"height,omitempty"`

	InsuranceValue float64 `json:"insurance_value"`
	ServiceTypeID  int     `json:"service_type_id"`
	Items          []*Item `json:"items"`
}

type PreviewOrderReq struct {
	PaymentTypeID   int     `json:"payment_type_id"`
	Note            string  `json:"note"`
	RequiredNote    string  `json:"required_note"`
	ClientOrderCode string  `json:"client_order_code"`
	CODAmount       float64 `json:"cod_amount"`

	FromName       string `json:"from_name"`
	FromPhone      string `json:"from_phone"`
	FromAddress    string `json:"from_address"`
	FromDistrictID int    `json:"from_district_id"`
	FromWardCode   string `json:"from_ward_code"`

	ToName       string `json:"to_name"`
	ToPhone      string `json:"to_phone"`
	ToAddress    string `json:"to_address"`
	ToWardCode   string `json:"to_ward_code"`
	ToDistrictID int    `json:"to_district_id"`

	Content string `json:"content"`
	Weight  int    `json:"weight"`
	Length  int    `json:"length,omitempty"`
	Width   int    `json:"width,omitempty"`
	Height  int    `json:"height,omitempty"`

	InsuranceValue float64 `json:"insurance_value"`
	ServiceTypeID  int     `json:"service_type_id"`
	Items          []*Item `json:"items"`
}
