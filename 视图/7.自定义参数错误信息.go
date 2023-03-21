package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func GetVaildMsg(err error, user any) string {
	getObj := reflect.TypeOf(user) //将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//类型断言成功
		for _, e := range errs {

			//为什么使用typeof，因为typeof中有tag字段，可以通过这个获取其中的tag
			//typeof 返回一个type(一个内置类型)类型 valueof返回一个value（内置类型）类型
			//返回的这两个类型中有许多字段和方法
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required" msg:"请输入用户名"`
			Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {

			c.JSON(200, gin.H{"msg": GetVaildMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})

	router.Run(":80")
}
