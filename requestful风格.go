package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Article struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Response struct {
	Code int
	Data any
	Msg  string
}

func main() {
	router := gin.Default()
	router.GET("/article", _getlist)
	router.GET("/article/:id", _getDetil)
	router.POST("/article", _create)
	router.PUT("/article/:id", _update)
	router.DELETE("/article/:id", _delete)

	router.Run(":8080")
}

func _delete(context *gin.Context) {
	fmt.Println(context.Param("id"))
	context.JSON(200, Response{
		Code: 200,
		Data: nil,
		Msg:  "删除成功",
	})
}

func _update(ctx *gin.Context) {
	var ar Article
	fmt.Println(ctx.Param("id"))
	err := BindCreate(ctx, &ar)
	if err != nil {
		logrus.Info(err)
		return
	}
	ctx.JSON(200, Response{
		Code: 200,
		Data: ar,
		Msg:  "success",
	})
}

func BindCreate(context *gin.Context, v any) error {
	t := context.GetHeader("Content-Type")
	body, _ := context.GetRawData()
	switch t {
	case "application/json":
		err := json.Unmarshal(body, v)
		if err != nil {
			return err
		}
	}

	return nil
}
func _create(context *gin.Context) {
	var article Article
	err := BindCreate(context, &article)
	if err != nil {
		logrus.Info("解析错误")
		return
	}

	//这里的是any，但是不用指针我不知道是为什么
	context.JSON(200, Response{
		Code: 200,
		Data: article,
		Msg:  "success",
	})
}

func _getDetil(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		logrus.Info(err)
		return
	}
	articles := []Article{
		{Name: "kds", Title: "master"},
		{Name: "jode", Title: "server"},
	}
	context.JSON(200, Response{
		Code: 200,
		Data: articles[id],
		Msg:  "success",
	})
}

func _getlist(context *gin.Context) {
	var articles []Article
	articles = []Article{
		{Name: "kds", Title: "master"},
		{Name: "jode", Title: "server"},
	}
	context.JSON(200, Response{
		Code: 200,
		Data: articles,
		Msg:  "success",
	})
}
