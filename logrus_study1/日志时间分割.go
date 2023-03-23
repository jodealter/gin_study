package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Myhook struct {
	File     *os.File
	AppName  string
	FileDate string
	LogPath  string
}

func (m *Myhook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (m *Myhook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02-15-04")
	line, _ := entry.String()
	if timer == m.FileDate {
		m.File.Write([]byte(line))
		return nil
	}
	m.File.Close()
	err := os.MkdirAll(fmt.Sprint("./%s/%s", m.LogPath, timer), os.ModePerm)
	if err != nil {
		logrus.Errorln("目录创建错误")
		return err
	}
	file, err := os.OpenFile(fmt.Sprintf("./%s/%s/%s.log", m.LogPath, timer, m.AppName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Errorln("Init文件创建失败")
		return err
	}
	file.Write([]byte(line))
	return nil
}
func InitLogrus(logPath string, appName string) {
	fileDate := time.Now().Format("2006-01-02-15-04")

	err := os.MkdirAll(fmt.Sprintf("./%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Errorln("目录创建错误")
		return
	}
	file, err := os.OpenFile(fmt.Sprintf("./%s/%s/%s.log", logPath, fileDate, appName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Errorln("Init文件创建失败")
		return
	}
	myhook := Myhook{
		File:     file,
		AppName:  "fate",
		FileDate: fileDate,
		LogPath:  logPath,
	}
	logrus.AddHook(&myhook)
}
func main() {
	InitLogrus("logrus_study1", "fate")
	logrus.Warnln("警告")
}
