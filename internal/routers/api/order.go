package api

import (
	opb "blog-service/api/order"
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Order struct {
	srv service.Order
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
	app.NewResponse(c).ToResponse(opb.OrderResp{})
	return
}
