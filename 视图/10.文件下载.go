package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/download", func(c *gin.Context) {
		//指定头部，是一个数据流
		c.Header("Content-Type", "application/octet-stream")

		//设置头部字段指定名字
		c.Header("Content-Disposition", "attachment; filename="+"天草2.png")
		c.File("./upload/12.png")
	})
	router.Run(":8080")
}
