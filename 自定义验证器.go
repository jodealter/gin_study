package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"reflect"
)

// 返回false就说明不符合自定义的验证器去
func signVaild(fl validator.FieldLevel) bool {
	var nameList []string = []string{"kds", "jode"}
	for _, na := range nameList {
		name := fl.Field().Interface().(string)
		if name == na {
			return false
		}
	}
	return true
}

// 获取msg字段
func GetMsg(err error, user any) string {
	getObj := reflect.TypeOf(user)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {

			if v, exit := getObj.Elem().FieldByName(e.Field()); exit {
				return v.Tag.Get("msg")
			}
		}
	}
	return ""
}
func main() {
	router := gin.Default()

	//注册一个验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("sign", signVaild)

		if err != nil {
			logrus.Info(err)
			return
		}
	}
	router.POST("/", func(context *gin.Context) {
		type User struct {
			Name string `json:"name" binding:"required,sign" msg:"名字格式不对"`
			Age  int    `json:"age" binding:"required" msg:"年龄格式不对"`
		}
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			logrus.Info(err)
			context.JSON(200, GetMsg(err, &user))
			return
		}
		context.JSON(200, gin.H{"data": user})
	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
