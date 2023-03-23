package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/string", _string)

	//需要先加载模版文件,这个目录下的所有文件
	router.LoadHTMLGlob("./templates/*")
	router.GET("/html", _html)

	//json
	router.GET("/json", _json)

	//xml
	router.GET("/xml", _xml)

	//yaml
	router.GET("yaml", _yaml)

	//redirect
	router.GET("/bilibili", _bilibili)

	//static file 这个是加载静态文件，就是直接访问本地资源
	//第一个参数是用户可以输入的路径，第二个是本地的真实目录路径
	router.StaticFS("/static", http.Dir("static"))
	//这个类似于上边的，但是不是目录而是文件
	router.StaticFile("/jode", "static/jode.png")

	//没试成功
	//这个是如果路径匹配"/jodealter"前缀，则会优先在static这个路径下查找文件
	// router.Static("/jode.png", "/jode")

	router.StaticFileFS("/jodealter.png", "jode.png", http.Dir("./static"))

	router.Run(":8080")
}

func _bilibili(context *gin.Context) {
	//301是永久重定向，即在浏览器中缓存，任何对就路径的访问都会转到新的路径，以后切搜索权重什么的也都会传给新的搜索路径
	//302 是临时重定向，不会缓存，只是暂时把访问的路径重定向到别的路径
	context.Redirect(302, "https://bilibili.com")
}

// 感觉这个例子 yaml与xml没什么区别
func _yaml(context *gin.Context) {
	context.XML(http.StatusOK, gin.H{"name": "kds", "power": "A", "data": gin.H{"power": "5A"}})
}

func _xml(context *gin.Context) {
	context.XML(http.StatusOK, gin.H{"name": "kds", "power": "A", "data": gin.H{"power": "5A"}})
}

func _json(context *gin.Context) {

	//可以是map
	/*
		user := map[string]string{
			"name": "kds",
		}
	*/

	//也可以是struct
	type serv struct {
		//但是这里边的字段需要大写
		Name string `json:"name"`
	}
	user := serv{Name: "jode"}
	context.JSON(200, &user)

}

func _html(context *gin.Context) {
	user := map[string]string{
		"username": "kds",
	}
	context.HTML(200, "index.html", &user)
}

func _string(context *gin.Context) {
	context.String(200, "你好string 马庆友")
}
