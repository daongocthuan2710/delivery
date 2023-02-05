package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewConnection(add string) (*grpc.ClientConn, error) {
	return grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
