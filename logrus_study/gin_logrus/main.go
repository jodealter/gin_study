package main

import (
	"gin_study/logrus_study/gin_logrus/log"
	"gin_study/logrus_study/gin_logrus/middleware"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	log.InitFile("logrus_study/gin_logrus/logs", "fate")
	router := gin.New()
	router.Use(middleware.LogMiddleware())
	router.GET("/", func(c *gin.Context) {
		logrus.Info("来了")
		c.JSON(200, gin.H{"msg": "你好"})
	})
	router.Run(":8081")
}
