package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	router := gin.Default()

	router.POST("/files", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			return
		}
		files := form.File["files"]
		count := 0
		for _, file := range files {
			err := context.SaveUploadedFile(file, "./upload/"+file.Filename)
			if err != nil {
				logrus.Error(err)
			}
			count++
		}
		context.JSON(200, gin.H{"count": count})
	})
	router.POST("/file", func(context *gin.Context) {
		fileheader, _ := context.FormFile("jode")
		file, err := fileheader.Open()
		if err != nil {
			return
		}
		context.SaveUploadedFile(fileheader, "./uploda/jodealter.png")
		image, err := os.Create("./upload/jode.png")

		// io.ReadAll(file) 读取这个文件有多少啊个字节
		defer func(image *os.File) {
			err := image.Close()
			if err != nil {
				logrus.Info("文件关闭错误")
			}
		}(image)
		count, err := io.Copy(image, file)
		if err != nil {
			return
		}
		context.JSON(200, gin.H{"count": count})
	})
	err := router.Run(":8080")
	if err != nil {
		logrus.Error(err)
		return
	}
}
