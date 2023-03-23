package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Responses struct {
	Code int
	Data any
	Msg  string
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		//下边这种方式不区分大小写
		fmt.Println(context.GetHeader("Content-Type"))
		fmt.Println(context.GetHeader("User-Agent"))

		//下边这种方式区分大小写
		fmt.Println(context.Request.Header.Get("Content-Type"))
		fmt.Println(context.Request.Header["User-Agent"])
	})
	router.GET("/index", func(context *gin.Context) {
		head := context.GetHeader("Content-Type")
		have := strings.Contains(head, "python")
		if have {
			context.JSON(200, Responses{
				Code: 200,
				Data: nil,
				Msg:  "你是爬虫",
			})
			return
		}
		context.JSON(200, Responses{
			Code: 200,
			Data: nil,
			Msg:  "你是用户",
		})
	})
	router.GET("/res", func(context *gin.Context) {
		context.Header("Context-Type", "application/text")
		context.JSON(200, gin.H{"data": "jode"})
	})
	router.Run(":8080")
}
