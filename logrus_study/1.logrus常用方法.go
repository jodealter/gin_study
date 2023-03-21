package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	//设置这个日志等级之后，比这个等级低的日志就不会打印了，毕竟你只觉得这个等级之上的日志才觉得对你有用
	//根据特定情况选定日志等级
	logrus.SetLevel(logrus.WarnLevel)
	logrus.Error("出错了")
	logrus.Warnln("警告")
	logrus.Infof("信息")
	logrus.Debugf("deBug")
	logrus.Println("打印")
	fmt.Println(logrus.GetLevel())
}
