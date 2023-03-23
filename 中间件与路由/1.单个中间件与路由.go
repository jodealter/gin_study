package main

import (
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {

	// 这里可以阻断后续的中间件
	//但是这个中间件的操作会完成
	c.Abort()

	//这里可以直接进行下一个中间件
	c.Next()
	c.JSON(200, gin.H{"msg": "m1...."})

}

func main() {
	router := gin.Default()
	router.GET("/", m1, func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "m2..."})
	})
	router.Run(":80")
}
