package global

import (
	opb "blog-service/api/order"
	"google.golang.org/grpc"
)

var OrderClient opb.OrderServiceClient

type ClientConfig struct {
	OrderAddr string
}

func GetOrderClient() opb.OrderServiceClient {
	conn, err := grpc.Dial(HostSettings.OrderHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	OrderClient = opb.NewOrderServiceClient(conn)
	return OrderClient
}
