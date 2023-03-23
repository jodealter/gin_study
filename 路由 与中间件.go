package main

import (
	"github.com/gin-gonic/gin"
)

type Serv struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func m1(c *gin.Context) {
	c.String(200, "m1...in")
	c.Set("user1", Serv{
		Name: "jode",
		Age:  22,
	})
	c.Next()
	c.String(200, "m1...out")
}

func m2(c *gin.Context) {
	c.String(200, "m2...in")
	c.Set("user2", Serv{
		Name: "jode",
		Age:  22,
	})
	c.Abort()
	c.String(200, "m2...out")
}

type UsersInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ArticleInfo struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type ResponseInfo struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func UserListView(context *gin.Context) {
	var users []UsersInfo = []UsersInfo{
		{"kds", 22},
		{"jode", 22},
	}
	context.JSON(200, ResponseInfo{
		Code: 200,
		Data: users,
		Msg:  "success",
	})
}

func ArticleListView(context *gin.Context) {
	var Article []ArticleInfo = []ArticleInfo{
		{"fate/stady-night", "2002"},
		{"fate/zero", "2022"}}
	context.JSON(200, ResponseInfo{
		Code: 200,
		Data: Article,
		Msg:  "success",
	})
}

func UserInit(router *gin.RouterGroup) {
	userMagager := router.Group("usManager")
	{
		userMagager.GET("/list", UserListView)
	}
}

func ArticleInit(router *gin.RouterGroup) {
	articleManager := router.Group("arManager")
	{
		articleManager.GET("/list", ArticleListView)
	}
}
func MiddleWare(msg string) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "1234" {
			context.Next()
			return
		}
		context.JSON(200, ResponseInfo{
			Code: 200,
			Data: nil,
			Msg:  msg,
		})
	}
}
func main() {
	router := gin.Default()

	router.GET("/", m1, func(context *gin.Context) {
		context.JSON(200, "")
		context.JSON(200, gin.H{"data": "nihao"})
		_user, _ := context.Get("user1")
		serv := _user.(Serv)
		context.JSON(200, gin.H{"data": serv})
	})
	//router.Use(m2)

	api := router.Group("api")
	UserInit(api)
	ArticleInit(api)

	api2 := router.Group("api2")
	ai := api2.Group("/login", MiddleWare("信息错误"), func(context *gin.Context) {
		context.JSON(200, gin.H{"data": "wule"})
	})
	ai.GET("/username", func(context *gin.Context) {

	})
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
