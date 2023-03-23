package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Ser struct {
	Name   string `json:"name" form:"name" binding:"required"`
	Baoju  string `json:"baoju" form:"baoju" binding:"gt=1,lt=10"`
	Baoju2 string `json:"baoju2" form:"baoju2" binding:"eqfield=Baoju"`
}

func main() {
	var serv Ser
	router := gin.Default()
	router.POST("/", func(context *gin.Context) {
		err := context.ShouldBind(&serv)
		if err != nil {
			logrus.Info(err)
			context.String(200, err.Error())
			return
		}
		fmt.Println(serv)
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}

}
