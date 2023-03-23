package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		fileheader, err := c.FormFile("file")
		if err != nil {
			return
		}
		/*
			保存人间的一种方式
			fmt.Println(fileheader.Size / 1024)
			fmt.Println(fileheader.Filename)
			c.SaveUploadedFile(fileheader, "./upload/12.png")
		*/

		//注意这里的readall与io.copy 都会读取文件内容(这里就是fileheader)，但是只能读一遍，第二遍就不能读了(可能有一个指针指向了末尾)
		file, _ := fileheader.Open()

		//这个目录相对 "9.文件上传.go" 来算的
		dst := "./" + fileheader.Filename

		//create 已经m默认打开了文件，所以需要关闭
		out, _ := os.Create(dst)
		defer func(out *os.File) {
			err := out.Close()
			if err != nil {

			}
		}(out)
		num, err := io.Copy(out, file)
		if err != nil {
			return
		}
		fmt.Println(num)
		data, _ := io.ReadAll(file)
		fmt.Println(string(data))
		c.JSON(200, gin.H{"msg": "上传成功"})
	})

	router.POST("/uploads", func(c *gin.Context) {
		//multi代表多了
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files {
			err := c.SaveUploadedFile(file, "./upload/"+file.Filename)
			if err != nil {
				return
			}
		}
		c.JSON(200, gin.H{"msg": "上传成功"})
	})
	err := router.Run(":80")
	if err != nil {
		return
	}
}
