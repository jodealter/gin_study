package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 动态参数
func _param(c *gin.Context) {
	fmt.Println(c.Param("name"))
	fmt.Println(c.Param("serv"))
}

// 查询参数
func _query(c *gin.Context) {
	fmt.Println(c.Query("name"))
	fmt.Println(c.GetQuery("name"))
	fmt.Println(c.GetQueryArray("name"))

	//这个键默认的value
	fmt.Println(c.DefaultQuery("serv", "jodealter"))

}

func _form(c *gin.Context) {
	//获得key对应的value
	fmt.Println(c.PostForm("name"))
	//获得这个键的对应的所有value
	fmt.Println(c.PostFormArray("name"))

	//这个同上，是在客户不提供的时候，给一个默认值
	fmt.Println(c.DefaultPostForm("serv", "jodealter"))

	//这个是返回所有的键值，以map（键是string，value是[]string）的形式返回，还包括error
	fmt.Println(c.MultipartForm())
}
func BindRaw(c *gin.Context, v any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err := json.Unmarshal(body, v)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}
func RawWithBind(c *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var user User
	err := BindRaw(c, &user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
}
func _raw(c *gin.Context) {

	//获得原始数据
	data, err := c.GetRawData()
	fmt.Println(data, err)
	fmt.Println(string(data))

	//从头部字段获得传递过来内容的类型
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		type User struct {
			Name string `json:"name1"`
			Age  int    `json:"age"`
		}
		var user User
		//解码
		err := json.Unmarshal(data, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
	}

}
func main() {
	router := gin.Default()

	//访问格式  /query?name=kds&name=jode
	router.GET("/query", _query)

	//访问格式  /param/kds/jode   这里的kds与jode是客户传入浏览器的，这里是举个例子
	router.GET("/param/:name/:serv", _param)

	//表单提交的是post，post在提交方面比get有优势,
	//需要填写body中的form-data
	//访问格式 /form   在下边填写其他的
	//form能接受两种form-data，x-www-form-urlencoded
	router.POST("/form", _form)

	//原始参数。以原始数据传送，表现形式为 传 ‘a’，拿到手是97
	//有几种接受格式,
	/*	(1)form-data
		Content-Disposition: form-data; name="name"

		a
		----------------------------736524640675176274211346--
	*/
	//(2)x-www-form-urlencoded
	//[110 97 109 101 61 97] <nil>

	//(3)raw
	//传什么就是什么，这样就可以把原始数据解析成其他格式的数据
	router.GET("/raw", _raw)

	router.GET("/RawBind", RawWithBind)
	router.Run(":80")
}
