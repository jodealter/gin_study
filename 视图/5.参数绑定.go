package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  bool   `json:"sex" form:"sex" uri:"sex"`
}

func main() {
	router := gin.Default()
	var user UserInfo

	//参数绑定，指的是前端（用户）传进来的参数，与自己的结构体进行结合
	router.POST("/", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{"data": "你错了"})
			return
		}
		c.JSON(200, user)
	})
	router.POST("/query", func(context *gin.Context) {
		var user UserInfo
		err := context.ShouldBindQuery(&user)
		if err != nil {
			context.JSON(200, gin.H{"data": "nicuole"})
			return
		}
		context.JSON(200, user)
	})

	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBindUri(&user)
		if err != nil {
			c.JSON(200, gin.H{"data": "nicuole"})
			return
		}
		c.JSON(200, user)
	})

	//很多默认的会走form，如果上边格式不匹配的话就会走form 比如form-data
	router.POST("/form", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(200, gin.H{"data": "nicuole"})
			return
		}
		c.JSON(200, user)
	})
	router.Run(":80")

}
