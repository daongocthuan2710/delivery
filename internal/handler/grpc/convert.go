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
		Cod:         orderReq.GetCod(),
		Weight:      orderReq.GetWeight(),
		ServiceCode: orderReq.GetServiceCode(),
		From:        deliveryInfoPBToReq(orderReq.GetFrom()),
		To:          deliveryInfoPBToReq(orderReq.GetTo()),
	}
	return payload
}

func estimateFeePBToReq(payload *proto.EstimateFeeReq) req.EstimateFee {
	return req.EstimateFee{
		OrderValue: payload.GetOrderValue(),
		Cod:        payload.GetCod(),
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
