package log

import (
	"bytes"
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

func colorprint(color int, msg string, IsBg bool) {
	if !IsBg {
		fmt.Printf("\033[3%dm%s\033[0m", color, msg)
		return
	}
	fmt.Printf("\033[4%dm%s\033[0m", color, msg)

}

type Myformater struct {
}

func (f Myformater) Format(entry *logrus.Entry) ([]byte, error) {

	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	//fmt.Fprintf(b, "%s", entry.Message)
	return b.Bytes(), nil
}

func InitFile(logPath, appName string) {
	//logrus.SetFormatter(&Myformater{})
	fileDate := time.Now().Format("2006-01-02_12-04")
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
