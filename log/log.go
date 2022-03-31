package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"time"
)

var log_file *os.File

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

func init() {
	os.Mkdir("logs", 0777)
	file2, err := os.Open("logs/" + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		file2, _ = os.Create("logs/" + time.Now().Format("2006-01-02") + ".log")
	}
	log_file = file2
}

func (this *Log) Debug(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgGreen).Println(this.prefix, "DEBUG", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "DEBUG", a))
		this.file.Sync()
	}
}

func (this *Log) Info(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgBlue).Println(this.prefix, "INFO", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "INFO", a))
		this.file.Sync()
	}
}
func (this *Log) Warn(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgYellow).Println(this.prefix, "WARN", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "WARN", a))
		this.file.Sync()
	}
}
func (this *Log) Error(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgRed).Println(this.prefix, "ERROR", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "ERROR", a))
		this.file.Sync()
	}
}
func (this *Log) Panic(a ...interface{}) {
	if this.file == nil {
		color.New(color.FgCyan).Println(this.prefix, "PANIC", a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v %v %v", this.prefix, "PANIC", a))
		this.file.Sync()
	}
}

func (this *Log) Println(a ...interface{}) {
	if this.file == nil {
		color.New().Println(a)
	} else {
		this.file.WriteString(fmt.Sprintf("%v", a))
		this.file.Sync()
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
	return &Log{prefix: time.Now().Format("2006-01-02 15:04:05"), file: log_file}
}
