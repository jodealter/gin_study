package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

//参数错误信息指的是这个字段出错的时候，放出来的，验证器是限制这个字段的东西

func _GetVaildMsg(err error, user any) string {
	getObj := reflect.TypeOf(user) //将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//类型断言成功
		for _, e := range errs {
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}

type User struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名输入错误"`
	Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
}

// 返回true 就代表验证通过
func signValid(fl validator.FieldLevel) bool {
	var nameList []string = []string{"kds", "jode"}
	for _, nameStr := range nameList {
		name := fl.Field().Interface().(string)
		if name == nameStr {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	//这个ok检验的不是engine这个函数，而是他的返回值被断言成*validator.Validate是否会成功
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("sign", signValid)
		if err != nil {
			return
		}
	}
	router.POST("/", func(c *gin.Context) {
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": _GetVaildMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})
	err := router.Run(":80")
	if err != nil {
		return
	}
}
