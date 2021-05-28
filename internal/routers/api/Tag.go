package api

import (
	"blog-service/internal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}
func (t *Tag) List(c *gin.Context) {
	s := service.New(c)
	params := service.ListTagReq{}
	err := c.ShouldBind(&params)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
	}
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	//get count
	count, err := s.CountTag(&service.CountTagReq{Name: params.Name, State: params.State})

	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	}

	list, err := s.ListTag(&params, pager)

	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	}

	app.NewResponse(c).ToResponseList(list, count)

	return
}

func (t *Tag) Create(c *gin.Context) {
	s := service.New(c)
	req := service.CreateTagReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		app.NewResponse(c).ToErrorResponse(errcode.InvalidParams)
	}
	s.CreateTag(&req)
}
func (t *Tag) Update(c *gin.Context) {

}
func (t *Tag) Delete(c *gin.Context) {

}
