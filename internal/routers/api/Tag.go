package api

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct{}

func NewTag() *Tag {
	return &Tag{}
}
func (t *Tag) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}

func (t *Tag) Create(c *gin.Context) {

}
func (t *Tag) Update(c *gin.Context) {

}
func (t *Tag) Delete(c *gin.Context) {

}
