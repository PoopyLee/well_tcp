package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"time"
)

type Log struct {
	prefix string
}

var file *os.File

type Logging interface {
	Println(a ...interface{})
	Debug(a ...interface{})
	Info(a ...interface{})
	Warn(a ...interface{})
	Error(a ...interface{})
	Panic(a ...interface{})
	SetPrefix(Prefix string)
	SetLogFile(File string)
}

func init() {
	file = nil
}
func (this *Log) Debug(a ...interface{}) {
	if file != nil {
		file.WriteString(fmt.Sprintln(this.prefix, "DEBUG", a))
	}
	color.New(color.FgGreen).Println(this.prefix, "DEBUG", a)
}

func (this *Log) Info(a ...interface{}) {
	if file != nil {
		file.WriteString(fmt.Sprintln(this.prefix, "INFO", a))
	}
	color.New(color.FgBlue).Println(this.prefix, "INFO", a)
}
func (this *Log) Warn(a ...interface{}) {
	if file != nil {
		file.WriteString(fmt.Sprintln(this.prefix, "WARN", a))
	}
	color.New(color.FgYellow).Println(this.prefix, "WARN", a)
}
func (this *Log) Error(a ...interface{}) {
	if file != nil {
		file.WriteString(fmt.Sprintln(this.prefix, "ERROR", a))
	}
	color.New(color.FgRed).Println(this.prefix, "ERROR", a)
}
func (this *Log) Panic(a ...interface{}) {
	if file != nil {
		file.WriteString(fmt.Sprintln(this.prefix, "PANIC", a))
	}
	color.New(color.FgCyan).Println(this.prefix, "PANIC", a)
}

func (this *Log) Println(a ...interface{}) {
	if file != nil {
		file.WriteString(fmt.Sprintln(a))
	}
	color.New().Println(a)
}

func (this *Log) SetPrefix(Prefix string) {
	this.prefix = Prefix
}

func (this *Log) SetLogFile(File string) {
	file2, err := os.Create(File)
	if err != nil {
		color.New(color.FgRed).Println(this.prefix, "ERROR", err)
	}
	file = file2
}

//获取函数名字
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func NewLoger() Logging {
	return &Log{prefix: time.Now().Format("2006-01-02 15:04:05")}
}
