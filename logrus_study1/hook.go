package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type MyHook struct {
	writer io.Writer
}

func (m *MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (m *MyHook) Fire(entry *logrus.Entry) error {
	logrus.SetOutput(m.writer)
	fmt.Println(entry.Level)
	return nil
}

func main() {
	file, err := os.OpenFile("./logrus_study1/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Error(err)
		return
	}
	//log := logrus.WithField("AppName", "fate")
	logrus.AddHook(&MyHook{file})
	logrus.Warnln("warn了")
	logrus.Errorln("错了")

}
