package log

import (
	"github.com/fatih/color"
	"log"
	"runtime"
	"time"
)

type Log struct {
	prefix string
	log.Logger
}

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

func (this *Log) Debug(a ...interface{}) {
	color.New(color.FgGreen).Println(this.prefix, "DEBUG", a)
	this.Println(this.prefix, "DEBUG", a)
}

func (this *Log) Info(a ...interface{}) {
	color.New(color.FgBlue).Println(this.prefix, "INFO", a)
	this.Println(this.prefix, "INFO", a)
}
func (this *Log) Warn(a ...interface{}) {
	color.New(color.FgYellow).Println(this.prefix, "WARN", a)
	this.Println(this.prefix, "WARN", a)
}
func (this *Log) Error(a ...interface{}) {

	color.New(color.FgRed).Println(this.prefix, "ERROR", a)
	this.Println(this.prefix, "ERROR", a)
}
func (this *Log) Panic(a ...interface{}) {
	color.New(color.FgCyan).Println(this.prefix, "PANIC", a)
	this.Println(this.prefix, "PANIC", a)
}

func (this *Log) Println(a ...interface{}) {
	color.New().Println(a)
	this.Println(a)
}

func (this *Log) SetPrefix(Prefix string) {
	this.prefix = Prefix
}

func (this *Log) SetLogFile(File string) {
	this.SetLogFile(File)
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
