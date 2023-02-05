package grpc

import (
	"delivery/internal/model/proto"
	"delivery/internal/model/req"
	"delivery/internal/model/res"
)

func createDeliveryPBToReq(orderReq *proto.CreateDeliveryReq) req.DeliveryCreate {
	payload := req.DeliveryCreate{
		OrderCode:   orderReq.GetOrderCode(),
		Note:        orderReq.GetNote(),
		OrderValue:  orderReq.GetOrderValue(),
		COD:         orderReq.GetCod(),
		Weight:      orderReq.GetWeight(),
		ServiceCode: orderReq.GetServiceCode(),
		From:        deliveryInfoPBToReq(orderReq.GetFrom()),
		To:          deliveryInfoPBToReq(orderReq.GetTo()),
	}
	items := make([]req.Item, len(orderReq.GetItems()))
	for i, item := range orderReq.GetItems() {
		items[i] = itemPBToReq(item)
	}
	return payload
}

func itemPBToReq(item *proto.Item) req.Item {
	return req.Item{
		Name:     item.GetName(),
		Quantity: int(item.GetQuantity()),
		Weight:   item.GetWeight(),
	}
}

func estimateFeePBToReq(payload *proto.EstimateFeeReq) req.EstimateFee {
	return req.EstimateFee{
		OrderValue: payload.GetOrderValue(),
		COD:        payload.GetCod(),
		Weight:     payload.GetWeight(),
		From:       deliveryInfoPBToReq(payload.GetFrom()),
		To:         deliveryInfoPBToReq(payload.GetTo()),
	}
}

func estimateFeeResToPB(info *res.EstimateFee) *proto.EstimateFeeRes {
	services := make([]*proto.DeliveryService, len(info.Services))
	for i, item := range info.Services {
		services[i] = &proto.DeliveryService{
			Name:         item.Name,
			Code:         item.Code,
			TplCode:      item.TplCode,
			TotalFee:     item.TotalFee,
			EstimateTime: item.EstimateTime,
		}
	}
	return &proto.EstimateFeeRes{Services: services}
}

func deliveryInfoPBToReq(info *proto.DeliveryInfo) *req.DeliveryInfo {
	return &req.DeliveryInfo{
		Name:       info.GetName(),
		Phone:      info.GetPhone(),
		Address:    info.GetAddress(),
		ProvinceId: info.GetProvinceId(),
		DistrictId: info.GetDistrictId(),
		WardId:     info.GetWardId(),
	}
}
