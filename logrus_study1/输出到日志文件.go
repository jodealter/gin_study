package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	/*file, err := os.OpenFile("./logrus_study1/log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}*/
	file_c, err := os.Create("./logrus_study1/log.txt")
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	//logrus.SetOutput(io.MultiWriter(os.Stdout, file))

	logrus.SetOutput(io.MultiWriter(os.Stdout, file_c))
	logrus.Warnln("这是 警告 ")

}
