package log

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"sync"
	"time"
)

var log_file *os.File
var log_prefix string
var time_prefix string
var Debug bool

type Log struct {
	prefix string
	file   *os.File
	sync.RWMutex
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
	log_file = nil
	Debug = true
	log_prefix = "[WELL]"
	time_prefix = time.Now().Format("2006-01-02 / 15:04:05") + " ▶"
}

func (this *Log) Debug(a ...interface{}) {
	if this.file == nil {
		if Debug {
			color.New(color.FgGreen).Println(this.prefix, "DEBUG", a)
		}
	} else {
		this.Lock()
		defer this.Unlock()
		this.file.WriteString(fmt.Sprintln(this.prefix, "DEBUG", a))
	}
}

func (this *Log) Info(a ...interface{}) {
	if this.file == nil {
		if Debug {
			color.New(color.FgBlue).Println(this.prefix, "INFO", a)
		}
	} else {
		this.Lock()
		defer this.Unlock()
		this.file.WriteString(fmt.Sprintln(this.prefix, "INFO", a))
	}
}
func (this *Log) Warn(a ...interface{}) {
	if this.file == nil {
		if Debug {
			color.New(color.FgYellow).Println(this.prefix, RunFuncName(), "WARN", a)
		}
	} else {
		this.Lock()
		defer this.Unlock()
		this.file.WriteString(fmt.Sprintln(this.prefix, RunFuncName(), "WARN", a))
	}
}
func (this *Log) Error(a ...interface{}) {
	if this.file == nil {
		if Debug {
			color.New(color.FgRed).Println(this.prefix, RunFuncName(), "ERROR", a)
		}
		os.Exit(0)
	} else {
		this.Lock()
		defer this.Unlock()
		this.file.WriteString(fmt.Sprintln(this.prefix, RunFuncName(), "ERROR", a))
		os.Exit(0)
	}

}
func (this *Log) Panic(a ...interface{}) {
	if this.file == nil {
		if Debug {
			color.New(color.FgCyan).Println(this.prefix, RunFuncName(), "PANIC", a)
		}
		panic(fmt.Sprintln(a))
	} else {
		this.Lock()
		defer this.Unlock()
		this.file.WriteString(fmt.Sprintln(this.prefix, RunFuncName(), "PANIC", a))
		panic(fmt.Sprintln(a))
	}
}

func (this *Log) Println(a ...interface{}) {
	if this.file == nil {
		if Debug {
			color.New().Println(a)
		}
	} else {
		this.Lock()
		defer this.Unlock()
		this.file.WriteString(fmt.Sprintln(a))
	}
}

func (this *Log) SetPrefix(Prefix string) {
	log_prefix = Prefix
}

func (this *Log) SetLogFile(File string) {
	if this.file != nil {
		this.file.Close()
	}
	file2, err := os.OpenFile(File, os.O_APPEND, 0777)
	if err != nil {
		NewLoger().Error("no outfile,please restart server", err)
		file2, err = os.Create(File)
	}
	log_file = file2
}

//获取函数名字
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func NewLoger() Logging {
	return &Log{prefix: log_prefix + " " + time_prefix, file: log_file}
}
