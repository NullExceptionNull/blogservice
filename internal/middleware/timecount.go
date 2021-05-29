package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		log.Info("中间件开始执行了")
		c.Next()
		fmt.Println("中间件执行完毕")
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
