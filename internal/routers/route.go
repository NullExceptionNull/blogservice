package routers

import (
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger())
	g.Use(middleware.MiddleWare())
	g.Use(middleware.Recovery())
	//
	orderGroup := g.Group("/sale")
	{
		order := api.NewOrder()
		orderGroup.POST("/auto/order", order.DealAutoOrder)
		//v1.POST("/tags", tag.Create)
		//v1.PUT("/tags/:id", tag.Update)
		//v1.DELETE("/tags/:id", tag.Delete)
	}
	return g
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.HTML(500, "error.tmpl", gin.H{
		"title": "Error",
		"err":   err,
	})
}
