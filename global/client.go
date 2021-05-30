package global

import (
	opb "blog-service/api/order"
	"google.golang.org/grpc"
)

var OrderClient opb.OrderServiceClient

type ClientConfig struct {
	Addr string
}

func CreateConn(c *ClientConfig) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(c.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
		return nil, err
	}
	return conn, nil
}
