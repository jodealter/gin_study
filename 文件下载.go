package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.Header("Content-Type", "application/octet-stream")
		context.Header("Content-Disposition", "attachment; filename=jode.png")
		context.File("./upload/jode.png")
	})

	err := router.Run(":8080")
	if err != nil {
		logrus.Info(err)
		return
	}
}
