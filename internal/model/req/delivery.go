package req

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"

	"delivery/internal/constant"
	"delivery/internal/model/entity"
)

type DeliveryCreate struct {
	OrderCode   string        `json:"orderCode,omitempty"`
	Note        string        `json:"note,omitempty"`
	OrderValue  int64         `json:"orderValue,omitempty"`
	COD         int64         `json:"cod,omitempty"`
	Weight      int64         `json:"weight,omitempty"`
	ServiceCode string        `json:"serviceCode,omitempty"`
	From        *DeliveryInfo `json:"from,omitempty"`
	To          *DeliveryInfo `json:"to,omitempty"`
	Items       []Item        `json:"items"`
}

func (d DeliveryCreate) GetContent() string {
	arr := make([]string, len(d.Items))
	for i, item := range d.Items {
		s := fmt.Sprintf("%s x %d", item.Name, item.Quantity)
		arr[i] = s
	}
	return strings.Join(arr, ",")
}

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Weight   int64  `json:"weight"`
}

func (d DeliveryCreate) ToEntity() (*entity.Delivery, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	tpCode := constant.ServiceCodeMapping[d.ServiceCode]
	return &entity.Delivery{
		ID:                  null.StringFrom(id.String()),
		Code:                null.StringFrom(d.OrderCode),
		TrackingCode:        null.String{},
		Note:                null.StringFrom(d.Note),
		Status:              null.String{},
		CreatedAt:           null.TimeFrom(now),
		UpdatedAt:           null.TimeFrom(now),
		Value:               null.Int64From(d.OrderValue),
		Cod:                 null.Int64From(d.COD),
		Weight:              null.Int64From(d.Weight),
		ServiceCode:         null.StringFrom(d.ServiceCode),
		PartnerStatus:       null.String{},
		PartnerIdentityCode: null.StringFrom(tpCode),
		FromName:            null.StringFrom(d.From.Name),
		FromPhone:           null.StringFrom(d.From.Phone),
		FromAddress:         null.StringFrom(d.From.Address),
		FromProvinceCode:    null.IntFrom(int(d.From.ProvinceId)),
		FromDistrictCode:    null.IntFrom(int(d.From.DistrictId)),
		FromWardCode:        null.IntFrom(int(d.From.WardId)),
		ToName:              null.StringFrom(d.To.Name),
		ToPhone:             null.StringFrom(d.To.Phone),
		ToAddress:           null.StringFrom(d.To.Address),
		ToProvinceCode:      null.IntFrom(int(d.To.ProvinceId)),
		ToDistrictCode:      null.IntFrom(int(d.To.DistrictId)),
		ToWardCode:          null.IntFrom(int(d.To.WardId)),
	}, nil
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
	COD        int64
	Weight     int64
	From       *DeliveryInfo
	To         *DeliveryInfo
}
