package routers

import (
	"blog-service/internal/routers/api"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger())
	g.Use(gin.Recovery())

	v1 := g.Group("/api/v1")
	{
		tag := api.NewTag()
		v1.GET("/tags", tag.List)
		v1.POST("/tags", tag.Create)
		v1.PUT("/tags/:id", tag.Update)
		v1.DELETE("/tags/:id", tag.Delete)
	}
	return g
}
