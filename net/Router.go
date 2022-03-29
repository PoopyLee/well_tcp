package net

import "github.com/lvwei25/well_tcp/log"

//TODO	服务路由类
type ServerRouter struct {
}

type WellServerRouter interface {
	OnStart()
	OnClose()
}

//TODO	连接路由类
type ConnRouter struct {
}

//方法
type WellConnRouter interface {
	OnConnect(conn *WellConnection)
	OnError(err error)
	OnMessage(request *WellRequest)
	OnClose(conn *WellConnection)
}

func (this *ServerRouter) OnStart() {
	log.NewLoger().Info("OnStart")
}

func (this *ServerRouter) OnClose() {
	log.NewLoger().Info("OnClose")
}

func (this *ConnRouter) OnConnect(conn *WellConnection) {
	log.NewLoger().Info("OnConnect")
}

func (this *ConnRouter) OnError(err error) {
	log.NewLoger().Info("OnError")
}

func (this *ConnRouter) OnMessage(request *WellRequest) {
	log.NewLoger().Info("OnMessage")
}

func (this *ConnRouter) OnClose(conn *WellConnection) {
	log.NewLoger().Info("OnClose")
}
