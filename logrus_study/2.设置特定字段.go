package main

import "github.com/sirupsen/logrus"

func server() {

}
func send() {

}

func main() {
	//设置输出格式，两种json和text
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		//设置颜色的
		ForceColors: true,
	})

	//设置字段，在通过这个返回值进行日志调用，就会有这两个字段值
	log := logrus.WithField("app", "study").WithField("server", "logrus")
	logs := logrus.WithFields(logrus.Fields{
		"userid": "21",
		"ip":     "192.168.200.258",
	})
	logs = log.WithFields(logrus.Fields{
		"userid": "22",
		"ip":     "192",
	})
	logs.Error("我是你爹\n")
	log.Error("你好")
}
