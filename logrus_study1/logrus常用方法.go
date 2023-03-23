package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetLevel(logrus.WarnLevel)
	logrus.Errorln("error")
	logrus.Warnln("warn")
	logrus.Infoln("info")
	logrus.Debugln("debug")
	logrus.Println("print")
}
