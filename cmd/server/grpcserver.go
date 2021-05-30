package main

import (
	opb "blog-service/api/order"
	"blog-service/global"
	"blog-service/internal/order/rpcservice/order"
	"google.golang.org/grpc"
)

func StartAllGrpc() {
	startOrderGrpc()
}

func startOrderGrpc() {
	go func() {
		err := global.RunGrpcServer(&global.GRPCConfig{
			Name: "order",
			Addr: global.HostSettings.OrderHost,
			RegisterFunc: func(server *grpc.Server) {
				opb.RegisterOrderServiceServer(server, &order.OrderServer{})
			},
		})
		if err != nil {
			panic(err)
		}
	}()
}
