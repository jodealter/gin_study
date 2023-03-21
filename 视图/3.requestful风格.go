package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func _getList(c *gin.Context) {
	artiacles := []ArticleModel{
		{"go语言入门", "这篇是go语言入门"},
		{"py入门", "这是py入门"},
		{"c入门", "这是c入门"}}
	//c.JSON(200, articles)
	c.JSON(200, Response{
		Code: 0,
		Data: artiacles,
		Msg:  "成功",
	})
}
func _getDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	articles := []ArticleModel{
		{"go语言入门", "这篇是go语言入门"},
		{"py入门", "这是py入门"},
		{"c入门", "这是c入门"}}
	c.JSON(200, Response{
		Code: 0,
		Data: articles[id],
		Msg:  "nihao",
	})

}

func BindCreate(c *gin.Context, v any) (err error) {
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
func _create(c *gin.Context) {
	var ar ArticleModel
	err := BindCreate(c, &ar)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, Response{
		Code: 0,
		Data: ar,
		Msg:  "这是西游记",
	})

}
func _update(c *gin.Context) {

	//param 就是路径中的用户提供的
	fmt.Println(c.Param("id"))
	var ar ArticleModel
	err := BindCreate(c, &ar)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(0, Response{
		Code: 0,
		Data: ar,
		Msg:  "",
	})
}
func _delete(c *gin.Context) {
	fmt.Println(c.Param("id"))
	c.JSON(200, Response{
		Code: 0,
		Data: map[string]string{},
		Msg:  "删除成功",
	})
}

func main() {
	router := gin.Default()
	//下边是四种请求方式，requestful风格
	router.GET("/articles", _getList)       //文章列表
	router.GET("/articles/:id", _getDetail) //文章详情
	router.POST("/articles", _create)       //新建文章
	router.PUT("/articles/:id", _update)    //更新文章
	router.DELETE("/articles/:id", _delete) //删除文章

	router.Run(":80")
}
