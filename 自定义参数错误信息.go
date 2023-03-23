package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"reflect"
)

func GetVaildMsg(err error, user any) string {
	getObj := reflect.TypeOf(user)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {

			//这个err.Field返回实际字段名称，并不知道是哪个结构体的，只是字段的名称如Name

			//返回结构体中的字段,比如Name
			if f, exit := getObj.Elem().FieldByName(err.Field()); exit {

				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}
func main() {
	router := gin.Default()
	router.POST("/", func(context *gin.Context) {
		type User struct {
			Name string `json:"name"  binding:"required" msg:"年龄不对"`
		}
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			logrus.Info(err)
			context.JSON(200, gin.H{"msg": GetVaildMsg(err, &user)})
			return
		}
		context.JSON(200, user)
	})

	router.Run(":8080")
}
