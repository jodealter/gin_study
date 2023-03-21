package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m10(c *gin.Context) {
	fmt.Println("m10 ...in")
	//即使这个abort在下边的json之前，这个json也会执行，估计是执行整个函数的吧0
	//c.Abort()
	c.JSON(200, gin.H{"data": "你的吃掉了"})
	c.Next()
	fmt.Println("m10 ...out")
}

type User struct {
	Name string
	Age  int
}

func m11(c *gin.Context) {
	fmt.Println("m11 ...in")
	c.Set("name", "vale")

	//不能使用&引用，会报错
	c.Set("user", User{
		Name: "dd12",
		Age:  11,
	})
	c.Next()
	fmt.Println("m11 ...out")
}
func main() {
	router := gin.Default()
	//全局中间件指的是无论用户用哪个路径进行访问，都会处罚这几个中间件
	//这个是加入全局的中间件，
	//顺序是加入的顺序
	router.Use(m10, m11)
	router.GET("/index", func(c *gin.Context) {
		fmt.Println("index ...in")
		_user, _ := c.Get("user")
		user := _user.(User)
		fmt.Println(user.Name)
		c.Next()
		fmt.Println("index ...out")
	})
	router.GET("/jode", func(c *gin.Context) {
		fmt.Println("jode ...in")
		fmt.Println("jode ...out")
	})
	router.Run(":80")
}
