package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

type MyFormat struct {
	Prefix     string
	TimeFormat string
}

const (
	//cBlock  = 0
	cRed    = 1
	cGreed  = 2
	cYellow = 3
	cBlue   = 4
	//cPurple = 5
	//cCyan   = 6
	cGray = 7
)

// Format byte是要输出的东西
func (m MyFormat) Format(e *logrus.Entry) ([]byte, error) {
	var color int
	switch e.Level {
	case logrus.ErrorLevel:
		color = cRed
	case logrus.WarnLevel:
		color = cYellow
	case logrus.InfoLevel:
		color = cGreed
	case logrus.DebugLevel:
		color = cBlue
	default:
		color = cGray
	}
	var buf *bytes.Buffer

	if e.Buffer != nil {
		buf = e.Buffer
	} else {
		buf = &bytes.Buffer{}
	}
	format := e.Time.Format(m.TimeFormat)

	fileFormat := fmt.Sprintf("%s:%d", path.Base(e.Caller.File), e.Caller.Line)

	_, err := fmt.Fprintf(buf, "[%s] \033[3%dm[%s]\033[0m [%s] %s %s", m.Prefix, color, e.Level, format, fileFormat, e.Message)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {

	//将调用方法也加入到输出格式
	logrus.SetReportCaller(true)

	logrus.SetFormatter(MyFormat{
		Prefix:     "[GROM]",
		TimeFormat: "2006-01-02 01:02:03",
	})
	logrus.Warn("你好")
}
