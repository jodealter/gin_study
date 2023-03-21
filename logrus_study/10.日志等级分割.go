package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	alllog  = "all"
	errlog  = "errlog"
	warnlog = "warn"
	infolog = "info"
)

type FilelevelHook struct {
	file     *os.File
	errfile  *os.File
	warnfile *os.File
	infofile *os.File
	logPath  string
}

func (hook FilelevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook FilelevelHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	switch entry.Level {
	case logrus.ErrorLevel:
		hook.errfile.Write([]byte(line))
	case logrus.WarnLevel:
		hook.warnfile.Write([]byte(line))
	case logrus.InfoLevel:
		hook.infofile.Write([]byte(line))
	}
	hook.file.Write([]byte(line))
	return nil
}
func InitLevel(logpath string) {
	err := os.MkdirAll(fmt.Sprintf("%s", logpath), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}
	allfile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logpath, alllog), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	errfile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logpath, errlog), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	warnfile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logpath, warnlog), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	infofile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logpath, infolog), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	files := FilelevelHook{
		file:     allfile,
		errfile:  errfile,
		warnfile: warnfile,
		infofile: infofile,
		logPath:  logpath,
	}
	logrus.AddHook(&files)
}
func main() {
	InitLevel("logrus_study/log_")
	logrus.Error("err")
	logrus.Warnln("warn")
	logrus.Info("info")
}
