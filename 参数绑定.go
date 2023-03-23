package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  bool   `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	var user UserInfo
	router := gin.Default()

	//因为是用户填写表单数据，所以用post
	router.POST("/json", func(context *gin.Context) {
		err := context.ShouldBindJSON(&user)
		if err != nil {
			logrus.Info(err)
			context.JSON(200, gin.H{"data": "错了"})
			return
		}
		context.JSON(200, gin.H{"data": "yes"})
	})
	router.POST("/query", func(context *gin.Context) {
		err := context.ShouldBindQuery(&user)
		if err != nil {
			logrus.Info("解析错误")
			context.JSON(200, gin.H{"msg": "error"})
			return
		}
		fmt.Println(user)
		return
	})

	router.POST("/uri/:name/:age/:sex", func(context *gin.Context) {
		err := context.ShouldBindUri(&user)
		if err != nil {
			logrus.Info(err)
			context.String(200, err.Error())
			return
		}
		fmt.Println(user)

	})
	router.POST("/form", func(context *gin.Context) {
		err := context.ShouldBind(&user)
		if err != nil {
			logrus.Info(err)
			context.String(200, err.Error())
			return
		}
		fmt.Println(user)
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
