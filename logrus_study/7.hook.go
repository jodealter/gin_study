package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type MyHook struct {
	//有了这个字段之后，在主函数里openfile一个文件，就可以向这里边传了
	writer io.Writer
}

func (MyHook) Levels() []logrus.Level {
	//这里返回的level等级是这个hook会搭理的日志等级，这个等级之下的不会管
	return []logrus.Level{logrus.ErrorLevel}
}
func (m *MyHook) Fire(entry *logrus.Entry) error {
	logrus.SetOutput(m.writer)
	fmt.Println(entry.Level)
	return nil
}

func main() {

	//hook钩子函数并不会影响影响其他log以及别的函数的使用，只是伴随着发生

	//一开始就打开一个文件，就不用每次调用err时打开一个文件了，也可以传到下边的结构体里边去
	file, _ := os.OpenFile("err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logrus.AddHook(&MyHook{file})
	logrus.Error("你好")
	logrus.Warnf("你好")
}
