package main

import (
	"github.com/gin-gonic/gin"
)

type SignUserInfo2 struct {
	//Name string `json:"name" binding:"required,min=4,max=6"`

	//contains是包含某字符串
	//Name string `json:"name" binding:"contains=kds"`

	//excludes 是不许包含某一字符串
	//Name string `json:"name" binding:"excludes=jode,excludes=kds"`

	//startswith是以规定字符串开头
	//endswith是以规定字符串结尾
	Name string `json:"name" binding:"startswith=kds,endswith=jode"`

	Age int `json:"age" binding:"gt=10,lt=20"`

	//oneof是在几个给定选项中选择一个
	Sex        string `json:"sex" binding:"oneof=男 女"`
	Password   string `json:"password"`
	RePassword string `json:"re_password" binding:"eqfield=Password"`

	//dive 是针对每一个，这个字段后边的是针对数组每一个进行限制，注意各个字段之间不能有空格
	LikeList []string `json:"likelist" binding:"required,dive,startswith=jode" `

	//ip与url字段是对ip与url进行限制的
	Ip string `json:"ip" binding:"required,ip"`
	//url与uri的区别是uri是url后边的资源定位，uri是url的子集，类似于baidu.com/123 的/123
	//但是可以放在url里的绝对可以放在uri中，但是uri中的有时不能放在url中
	Url string `json:"url" binding:"required,url"`
	Uri string `json:"uri" binding:"required,uri"`

	//这个是时间的 这个时间就是提供的一个格式
	Date string `json:"date" binding:"datetime=2006-01-04 15:04:05"`
}

func main() {
	router := gin.Default()

	//json自带类型检测，所以可以不用特意进行类型匹配
	router.POST("/", func(c *gin.Context) {
		var user SignUserInfo2
		err := c.ShouldBindJSON(&user)
		if err != nil {

			c.JSON(200, gin.H{"data": err.Error()})
			return
		}
		c.JSON(200, user)
	})
	router.Run(":80")
}
