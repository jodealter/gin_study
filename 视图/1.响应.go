package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func _string(c *gin.Context) {
	c.String(200, "hello world") //响应string的
}
func _json(c *gin.Context) {
	/*
		这个是结构体类型
		type UserInfo struct {
			Username string `json:"username"`
			Age      int    `json:"age"`
			Password string `json:"-"` //-表示忽略，不显示这个字段，隐藏用的
		}
		user :=
			UserInfo{
				Username: "kds",
				Age:      22,
				Password: "send",
			}
		c.JSON(http.StatusOK, user)
	*/
	/*
		 json 响应map
		usermap := map[string]string{
			"Username": "kds",
			"Age":      "22",
			"Password": "kdssnd",
		}
		c.JSON(200, usermap)
	*/
	c.JSON(200, gin.H{"Username": "kds", "Age": 22, "Password": "send123."}) // 这个是直接响应json
}

func _xml(c *gin.Context) {
	c.XML(200, gin.H{"user": "kds", "message": "hello", "status": http.StatusOK, "data": gin.H{"user": "jode"}}) //这个是响应xml的
}
func _yaml(c *gin.Context) {
	c.YAML(200, gin.H{"user": "kds", "message": "hello", "status": http.StatusOK, "data": gin.H{"user": "jode"}}) //这个是响应xml的
}
func _html(c *gin.Context) {

	//结构体貌似需要修改html，但是我不会
	//但是可以使用map
	use := map[string]string{
		"username": "kds",
	}
	c.HTML(200, "index.html", use) //这里可以传一些与html中对应的
}
func _redirect(c *gin.Context) {
	c.Redirect(302, "https://www.baidu.com")
}
func main() {
	router := gin.Default()

	//加载目录下的所有模版文件例如index.html,index2.html
	//这样就可以使用HTML样式的模版
	//例如下边的c.HTML就是使用html模版
	router.LoadHTMLGlob("./templates/*")

	//选择可以下载的路径

	//golong中没有相对路径，只有项目路径
	//这里算是把这个文件夹加载进去了，可以根据路径（这个路径是给用户提供的路径（static），加上某个文件（相对程序中的给用户提供的)路径）访问文件夹里边的内容
	//第一个是给用户提供的静态目录的前缀（其实是用户输入这个，就能访问我规定的目录）
	router.StaticFS("/static", http.Dir("static"))

	router.StaticFile("/jode", "./static/jode.png") //第一个是用户访问的路径，第二个是文件存在的真是路径

	//如果是get请求，会返回去一定格式的东西
	router.GET("/string", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/baidu", _redirect)
	err := router.Run(":80")
	if err != nil {
		return
	}
}
