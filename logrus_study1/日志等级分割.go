package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	alllog  = "all"
	errlog  = "err"
	warnlog = "warn"
	infolog = "infor"
)

type LevelHook struct {
	allfile  *os.File
	errfile  *os.File
	warnfile *os.File
	infofile *os.File
	Path     string
}

func (l *LevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (l *LevelHook) Fire(entry *logrus.Entry) error {
	s, err := entry.String()
	if err != nil {
		logrus.Error(err)
		return err
	}
	switch entry.Level {
	case logrus.ErrorLevel:
		l.errfile.Write([]byte(s))
	case logrus.WarnLevel:
		l.warnfile.Write([]byte(s))
	case logrus.InfoLevel:
		l.infofile.Write([]byte(s))
	}
	l.allfile.Write([]byte(s))
	return nil
}

func InitLevel(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}
	filename_all := fmt.Sprintf("%s/%s.log", path, alllog)
	filename_err := fmt.Sprintf("%s/%s.log", path, errlog)
	filename_warn := fmt.Sprintf("%s/%s.log", path, warnlog)
	filename_info := fmt.Sprintf("%s/%s.log", path, infolog)

	allfile, err := os.OpenFile(filename_all, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	errfile, err := os.OpenFile(filename_err, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	warnfile, err := os.OpenFile(filename_warn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	infofile, err := os.OpenFile(filename_info, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}

	file := LevelHook{
		allfile:  allfile,
		errfile:  errfile,
		warnfile: warnfile,
		infofile: infofile,
		Path:     path,
	}
	logrus.AddHook(&file)
}
func main() {
	InitLevel("logrus_study1/log_level")
	logrus.Error("错了的了")
}
