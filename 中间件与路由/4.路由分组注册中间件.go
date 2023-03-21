package main

import (
	"github.com/gin-gonic/gin"
)

type _UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type _Response struct {
	Code int    `json:"code" :"code"`
	Data any    `json:"data" :"data"`
	Msg  string `json:"msg"`
}

// 这种闭包的可以进行传参
func Middleware(msg string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "1234" {
			c.Next()
			return
		}
		c.JSON(200, _Response{
			Code: 1001,
			Data: nil,
			Msg:  msg,
		})
		c.Abort()
	}
}
func _UserListView(c *gin.Context) {

	var userList []_UserInfo = []_UserInfo{{"kds", 22}, {"jode", 22}}
	c.JSON(200, _Response{
		Code: 200,
		Data: userList,
		Msg:  "消息成功",
	})
}

func _UserRouterInit(router *gin.RouterGroup) {
	userManager := router.Group("user_manager").Use(Middleware("用户验证失败")) //在这里使用一个返回值是符合条件的函数的话，就可以传参了
	{
		//访问这个路径 是/api/user_manager/users
		userManager.GET("/users", _UserListView)
		userManager.GET("user", _UserListView)
	}
}

func main() {
	router := gin.Default()
	//router := gin.New() //上边的default内部调用了new，并且在此基础上加了两个其他的中间件Logger(), Recovery()
	//具体效果的话，解开下边的panic，演示就知道了，用default很方便，但是用new可以自己添加配置logger等更加灵活

	//设置一个路由组,命名的时候不需要加 “/”，但是加/的时候需要加“/”
	api := router.Group("api")
	api.GET("/login", Middleware("登录验证错误"), func(c *gin.Context) {
		//panic("nicuole")
		c.JSON(200, gin.H{"data": "1234"})
	})
	_UserRouterInit(api)

	router.Run(":8080")
}
