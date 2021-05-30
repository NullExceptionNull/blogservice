package order

import (
	opb "blog-service/api/order"
	"context"
	"math/rand"
	"strconv"
)

type OrderServer struct {
	opb.UnimplementedOrderServiceServer
}

func (o *OrderServer) AutoOrder(context.Context, *opb.OrderReq) (*opb.OrderResp, error) {

	return &opb.OrderResp{
		Stat: 0,
		Data: &opb.OrderResultInfo{
			OrderNum: strconv.FormatInt(rand.Int63n(100000), 10),
		},
	}, nil
}
