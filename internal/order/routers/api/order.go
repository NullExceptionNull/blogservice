package api

import (
	opb "blog-service/api/order"
	"blog-service/global"
	"blog-service/internal/order/rpcservice/order"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Order struct {
	srv order.OrderServer
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) DealAutoOrder(c *gin.Context) {
	//接收token参数
	_ = c.DefaultQuery("token", "")

	orderParam := opb.OrderReq{}
	//绑定请求参数
	err := c.ShouldBindJSON(&orderParam)

	if err != nil {
		//参数格式不正确
		fmt.Printf("error %v", err)
		app.NewResponse(c).ToErrorResponse(errcode.OrderParamsError)
		return
	}
	//检查Token

	client := global.OrderClient

	autoOrder, err := client.AutoOrder(c, &orderParam)

	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
		return
	}

	app.NewResponse(c).ToResponse(autoOrder)

	return
}
