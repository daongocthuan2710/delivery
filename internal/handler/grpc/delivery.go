package grpc

import (
	"context"

	"delivery/internal/model/proto"
	"delivery/internal/service"
)

func NewDeliveryServer(services *service.Services) *DeliveryServer {
	return &DeliveryServer{
		UnimplementedDeliveryServer: proto.UnimplementedDeliveryServer{},
		services:                    services,
	}
}

type DeliveryServer struct {
	proto.UnimplementedDeliveryServer
	services *service.Services
}

func (d *DeliveryServer) EstimateFee(ctx context.Context, req *proto.EstimateFeeReq) (*proto.EstimateFeeRes, error) {
	payload := estimateFeePBToReq(req)
	res, err := d.services.Delivery.EstimateFee(ctx, payload)
	if err != nil {
		return nil, err
	}

	return estimateFeeResToPB(res), nil
}

func (d *DeliveryServer) CreateDelivery(ctx context.Context, orderReq *proto.CreateDeliveryReq) (*proto.CreateDeliveryRes, error) {
	payload := createDeliveryPBToReq(orderReq)
	res, err := d.services.Delivery.Create(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &proto.CreateDeliveryRes{
		OrderCode:    res.OrderCode,
		TrackingCode: res.TrackingCode,
		Status:       res.Status,
		TotalFee:     res.TotalFee,
	}, nil
}
