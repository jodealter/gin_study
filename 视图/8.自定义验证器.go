package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

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
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
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
	router.Run(":80")
}
