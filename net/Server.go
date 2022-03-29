package net

import (
	"fmt"
	"github.com/lvwei25/well_tcp/log"
	"net"
	"runtime"
	"strings"
)

var Logo = `
██╗    ██╗███████╗██╗     ██╗      ██████╗  ██████╗ 
██║    ██║██╔════╝██║     ██║     ██╔════╝ ██╔═══██╗
██║ █╗ ██║█████╗  ██║     ██║     ██║  ███╗██║   ██║
██║███╗██║██╔══╝  ██║     ██║     ██║   ██║██║   ██║
╚███╔███╔╝███████╗███████╗███████╗╚██████╔╝╚██████╔╝
 ╚══╝╚══╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝  ╚═════╝
`
var topLine = `=========================================WellGo=========================================`
var bottomLine = `========================================================================================`

var github = `https://github.com/lvwei25/well_tcp`
var gitee = `https://gitee.com/leelvwei/well_tcp`

type WellServer struct {
	Name         string
	IpAddr       string
	Port         string
	ServerRouter WellServerRouter
	ConnRouter   WellConnRouter
	IpVersion    string
	Version      string
	lin          *net.TCPListener
}

type wellServerInterface interface {
	Run()
	AddServerRouter(s WellServerRouter)
	AddConnRouter(c WellConnRouter)
	Close()
}

func (this *WellServer) Run() {
	addr, err := net.ResolveTCPAddr(this.IpVersion, this.IpAddr+":"+this.Port)
	if err != nil {
		log.NewLoger().Error(err)
	}
	lin, err := net.ListenTCP(this.IpVersion, addr)
	if err != nil {
		log.NewLoger().Error(err)
	}
	this.lin = lin
	defer lin.Close()

	log.NewLoger().Info("The Server Name:", this.Name, "The Server Version:", this.Version, "And Now Server is Running!")

	this.ServerRouter.OnStart() //服务启动前执行函数

	for {
		con, err := lin.AcceptTCP()
		if err != nil {
			log.NewLoger().Error("Connection error execution", err)
			this.ConnRouter.OnError(err) //连接出错时执行
		}
		c := NewConnHandle(con)
		c.setConnRouter(this.ConnRouter)
		go c.Start()
	}

}

func (this *WellServer) AddServerRouter(s WellServerRouter) {
	log.NewLoger().Info("Add ServerRouter Success!")
	this.ServerRouter = s
}

func (this *WellServer) AddConnRouter(c WellConnRouter) {
	log.NewLoger().Info("Add ConnRouter Success!")
	this.ConnRouter = c
}

func (this *WellServer) Close() {
	this.lin.Close()
}

//创建句柄
func NewServerHandle(Name, IpAddr, Port string) wellServerInterface {
	s := WellServer{
		Name:         Name,
		IpAddr:       IpAddr,
		Port:         Port,
		ServerRouter: &ServerRouter{},
		ConnRouter:   &ConnRouter{},
		Version:      "1.1.1",
		IpVersion:    "tcp",
		lin:          nil,
	}
	runlogo(s)
	return &s
}

func runlogo(s WellServer) {
	fmt.Println(topLine)
	fmt.Println(fmt.Sprintf("\tWellGo Version:%s\tNeed Go Version:1.16\tGo Version:%s", s.Version, runtime.Version()))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("\t%s:%s:%s\tGoroutine Count:%d", strings.ToUpper(s.IpVersion), s.IpAddr, s.Port, runtime.NumGoroutine()))
	fmt.Println("")
	fmt.Println(bottomLine)
	fmt.Println(fmt.Sprintf("[Github] %s", github))
	fmt.Println(fmt.Sprintf("[Gitee]  %s", gitee))
}
