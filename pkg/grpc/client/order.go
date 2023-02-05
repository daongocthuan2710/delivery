package client

import (
	"google.golang.org/grpc"

	"delivery/internal/model/proto"
)

func NewOrderClient(add string) (*grpc.ClientConn, proto.OrderClient, error) {
	conn, err := NewConnection(add)
	if err != nil {
		return nil, nil, err
	}
	c := proto.NewOrderClient(conn)
	return conn, c, nil
}
