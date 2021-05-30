package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		t2 := time.Since(t)
		logrus.Infof("this request is cost: %s", t2)
	}
}
