package main

import (
	"github.com/gin-gonic/gin"
)

type SignUserInfo struct {

	//这个binding:"required"指定了这个参数不能为空，空字串也不可以
	//binding后边的参数可以多个用“,”进行分割，但是不可以用两个binding写，检测不到
	//min 与 max是针对字符串的
	Name string `json:"name" binding:"required,min=4,max=6"`

	//针对数字有大小限制范围，gt是最小，lt是最大 eq是等于
	Age int `json:"age" binding:"gt=10,lt=20"`

	//eqfield是针对同级字段是否相等，比如密码，和重复密码,=后边的是自定义的不是json的
	Password string `json:"password"`

	//RePassword string `json:"re_password" binding:"eqfield=Password"`
	//json 后边加“-”表示忽略这个字段不进行绑定
	//虽然binding后加“-”说的跟json后加这个东西一样，但是感觉没有效果
	RePassword string `json:"re_password" binding:"-"`
}

func main() {
	router := gin.Default()

	//json自带类型检测，所以可以不用特意进行类型匹配
	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {

			c.JSON(200, gin.H{"data": err.Error()})
			return
		}
		c.JSON(200, user)
	})
	router.Run(":80")
}
