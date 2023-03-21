package main

import (
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type ArticleInfo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Response struct {
	Code int    `json:"code" :"code"`
	Data any    `json:"data" :"data"`
	Msg  string `json:"msg"`
}

func UserListView(c *gin.Context) {
	var userList []UserInfo = []UserInfo{{"kds", 22}, {"jode", 22}}
	c.JSON(200, Response{
		Code: 200,
		Data: userList,
		Msg:  "消息成功",
	})
}
func ArticleListView(c *gin.Context) {
	var articlelistview []ArticleInfo = []ArticleInfo{{"御座", "guardianship"}, {"fgo", "fgo"}}
	c.JSON(200, Response{
		Code: 200,
		Data: articlelistview,
		Msg:  "长传成功",
	})
}
func UserInit(router *gin.RouterGroup) {
	userManager := router.Group("user_manager")
	{
		//访问这个路径 是/api/user_manager/users
		userManager.GET("/users", UserListView)
		userManager.GET("user", UserListView)
	}
}
func ArticleInit(router *gin.RouterGroup) {
	ArticleManager := router.Group("article_manager")
	{
		ArticleManager.GET("/articles", ArticleListView)
		ArticleManager.GET("/article", ArticleListView)
	}
}
func main() {
	router := gin.Default()
	//设置一个路由组,命名的时候不需要加 “/”，但是加/的时候需要加“/”
	api := router.Group("api")
	UserInit(api)
	ArticleInit(api)
	router.Run(":8080")
}
