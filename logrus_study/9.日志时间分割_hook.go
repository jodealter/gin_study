package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type FileDateHook struct {
	file     *os.File
	logPath  string
	Filedate string
	appname  string
}

func (hook FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook FileDateHook) Fire(entry *logrus.Entry) error {
	fmt.Println(entry)
	timer := entry.Time.Format("2006-01-02_12-04")
	line, _ := entry.String()
	if hook.Filedate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	hook.file.Close()
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)
	filename := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appname)
	hook.file, _ = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	hook.file.Write([]byte(line))
	return nil
}
func InitFile(logPath, appName string) {
	fileDate := time.Now().Format("2006-01-02-12-04")
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}
	filename := fmt.Sprintf("%s/%s/%s", logPath, fileDate, appName)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	filehook := FileDateHook{
		file:     file,
		logPath:  logPath,
		Filedate: fileDate,
		appname:  appName,
	}
	logrus.AddHook(&filehook)
}
func main() {
	InitFile("logrus_study/log", "fatego")
	for {
		logrus.Warn("ruler")

		time.Sleep(20 * time.Second)
		logrus.Error("alter")
	}
}
