package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		//注意这里是不区分大小写的
		fmt.Println(c.GetHeader("Content-Type"))
		fmt.Println(c.GetHeader("User-Agent"))

		//获取全部的头部
		//Header 实质是一个map[string][]string
		//所以如果使用get，会取切片的第一个（前提是有多个）
		fmt.Println(c.Request.Header)
		fmt.Println(c.Request.Header.Get("User-Agent"))
		fmt.Println(c.Request.Header["User-Agent"])
		//fmt.Println(c.Request.Header["User-agent"])//这里是不行的，因为是一个map，所以需要匹配，上边可以是因为底层调用了header，再加上一系列处理就不区分大小写了
		//头部可以自定义，也同样遵守上边的规定

	})
	router.GET("/index", func(c *gin.Context) {
		header := c.GetHeader("User-Agent")
		if strings.Contains(header, "python") {
			c.JSON(200, gin.H{"data": "你是爬虫"})
			return
		}
		c.JSON(200, gin.H{"data": "你是用户"})
	})
	router.GET("/res", func(c *gin.Context) {
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(200, gin.H{"serv": "jode"})
	})
	err := router.Run(":80")
	if err != nil {
		return
	}
}
