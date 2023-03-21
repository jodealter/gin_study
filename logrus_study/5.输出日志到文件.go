package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	//create 与openfile的区别是create会覆盖源文件，但open不会，可以自定义某些模式
	//file, _ := os.Create("logrus_study/log.txt")
	file, _ := os.OpenFile("logrus_study/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	//这个setoutput可以加入多个输出路径
	//logrus.SetOutput(file)
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))
	logrus.Errorf("jnicuole")
}
