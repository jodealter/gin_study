package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

type Color int

const (
	cBlock  = 0
	cRed    = 1
	cGreed  = 2
	cYellow = 3
	cBlue   = 4
	cPurple = 5
	cCyan   = 6
	cGray   = 7
)

func colorprint(color int, msg string, IsBg bool) {
	if !IsBg {
		fmt.Printf("\033[3%dm%s\033[0m", color, msg)
		return
	}
	fmt.Printf("\033[4%dm%s\033[0m", color, msg)

}

type Myformater struct {
	TimeFormat string
	Prefix     string
}

func (f Myformater) Format(entry *logrus.Entry) ([]byte, error) {
	var color Color
	switch entry.Level {
	case logrus.ErrorLevel:
		color = cRed
	case logrus.WarnLevel:
		color = cYellow
	case logrus.InfoLevel:
		color = cGray
	case logrus.DebugLevel:
		color = cRed
	default:
		color = cGray
	}
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}
	//自定义时间输出格式
	format_time := entry.Time.Format(f.TimeFormat)

	//这个是自定义输出行号格式
	//第一个参数是指定文件，第二个参数指定行号
	//fileval := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)

	//这个与上边这个相比加上了path.base ,效果就是只输出路径不一样的地方，去除相同的前缀
	fileval := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

	fmt.Fprintf(b, "[%s] \033[3%dm[%s]\033[0m [%s] %s %s\n", f.Prefix, color, entry.Level, format_time, fileval, entry.Message)
	return b.Bytes(), nil
}

func main() {
	//开启函数名和行号
	logrus.SetReportCaller(true)

	//设置日志等级
	logrus.SetLevel(logrus.DebugLevel)

	//自定义输出格式的好处灵活，很多参数可以通过结构体的方式来定义自己的输出格式
	//只要实现这个方法签名的就可以作为格式被填进去
	logrus.SetFormatter(&Myformater{"2006-01-02 05:06:07", "GROM"})

	//这里的日志实际上是去调用上边的这个格式
	logrus.Errorln("nihao")
	logrus.Warnln("nihao")
	logrus.Debugln("nihao")
	logrus.Infoln("nihao")
}
