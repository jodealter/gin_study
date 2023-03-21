package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	//设置输出格式，两种json和text
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//下边的true字段是颜色的
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 05:06:07",
	})
	logrus.Errorf("hello")
	logrus.Infof("hello")
	logrus.Warnf("hello")
	logrus.Debugf("hello")
	logrus.Println("hello")
}
