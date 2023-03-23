package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	//请求参数是？后边的
	router.GET("/query", _query)

	//路径参数
	router.GET("/param/:name/:ser", _param)

	//post提交表单
	router.POST("/post", _post)

	//raw
	router.GET("/raw", _raw)

	//RowWithBind
	router.GET("/RawBind", _RawWithBind)

	_ = router.Run(":8080")

}

func RawBind(context *gin.Context, v any) error {
	body, _ := context.GetRawData()
	switch context.GetHeader("Content-Type") {

	case "application/json":
		err := json.Unmarshal(body, v)
		if err != nil {
			logrus.Info("error", err)
			return err
		}
	}
	return nil
}

func _RawWithBind(context *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	_ = RawBind(context, &user)
	fmt.Println(user)
}

func _raw(context *gin.Context) {

	//现获取原始数据
	data, err := context.GetRawData()
	fmt.Println(data, err)
	//判断类型
	switch context.GetHeader("Content-Type") {
	case "application/json":
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		//在进行解码
		err := json.Unmarshal(data, &user)
		if err != nil {
			logrus.Info(err)
		}
		fmt.Println(user)
	}
}

func _post(context *gin.Context) {
	fmt.Println(context.PostForm("name"))
	fmt.Println(context.PostFormArray("name"))

	//这个map只有在content-tyoe是application/x-www-form-urlencoded是才管用，否则返回空的map
	fmt.Println(context.PostFormMap("name"))

	//这个是在用户不给key的时候，自己设置一个value，比如用户不给my_name,我们就自己使用默认的"kds“
	fmt.Println(context.DefaultPostForm("my_name", "kds"))

	//返回所有的键值，记住mul的就是多个，
	fmt.Println(context.MultipartForm())
}

func _param(context *gin.Context) {
	fmt.Println(context.Param("name"))
	fmt.Println(context.Param("ser"))

}

func _query(context *gin.Context) {
	fmt.Println(context.Query("name"))

	//带get的与不带get的区别就是，带get会返回一个bool值，这个值代表是否存在这样的query
	fmt.Println(context.GetQuery("name"))
	fmt.Println(context.QueryArray("name"))
	fmt.Println(context.GetQueryArray("name"))

}
