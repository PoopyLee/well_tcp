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
	file   *os.File
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
	if this.file == nil {
		color.New(color.FgGreen).Println(this.prefix, "DEBUG", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "DEBUG", a))
	}
}

func (this *Log) Info(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgBlue).Println(this.prefix, "INFO", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "INFO", a))
	}
}
func (this *Log) Warn(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgYellow).Println(this.prefix, "WARN", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "WARN", a))
	}
}
func (this *Log) Error(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgRed).Println(this.prefix, "ERROR", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "ERROR", a))
	}
}
func (this *Log) Panic(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgCyan).Println(this.prefix, "PANIC", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "PANIC", a))
	}
}

func (this *Log) Println(a ...interface{}) {
	if this.file == nil {
		color.New().Println(a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v", a))
	}
}

func (this *Log) SetPrefix(Prefix string) {
	this.prefix = Prefix
}

func (this *Log) SetLogFile(File string) {
	if this.file != nil {
		this.file.Close()
	}
	this.file, _ = os.Create(File)
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
