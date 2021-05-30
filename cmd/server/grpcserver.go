package main

import (
	opb "blog-service/api/order"
	"blog-service/global"
	"blog-service/internal/order/rpcservice/order"
	"google.golang.org/grpc"
)

const endpoint = "localhost:8081"

func StartAllGrpc() {
	startOrderGrpc()
}

func startOrderGrpc() {
	go func() {
		err := global.RunGrpcServer(&global.GRPCConfig{
			Name: "order",
			Addr: endpoint,
			RegisterFunc: func(server *grpc.Server) {
				opb.RegisterOrderServiceServer(server, &order.OrderServer{})
			},
		})
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		conn, _ := global.CreateConn(&global.ClientConfig{
			Addr: endpoint,
		})
		global.OrderClient = opb.NewOrderServiceClient(conn)
	}()
}
