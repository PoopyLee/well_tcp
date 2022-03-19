package net

import (
	"Well/log"
	"github.com/fatih/color"
	"net"
)

var Logo = `
____    __    ____  _______  __       __      
\   \  /  \  /   / |   ____||  |     |  |     
 \   \/    \/   /  |  |__   |  |     |  |     
  \            /   |   __|  |  |     |  |     
   \    /\    /    |  |____ |  |----.|  |----.
    \__/  \__/     |_______||_______||_______|

`

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

func init() {
	color.New(color.FgHiMagenta).Println(Logo)
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
			log.NewLoger().Error("connection error execution", err)
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

func NewServerHandle(Name, IpAddr, Port string) wellServerInterface {
	s := WellServer{
		Name:         Name,
		IpAddr:       IpAddr,
		Port:         Port,
		ServerRouter: &ServerRouter{},
		ConnRouter:   &ConnRouter{},
		Version:      "1.0",
		IpVersion:    "tcp",
		lin:          nil,
	}
	return &s
}
