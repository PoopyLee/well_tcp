package net

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

}

func (this *ServerRouter) OnClose() {

}

func (this *ConnRouter) OnConnect(conn *WellConnection) {

}

func (this *ConnRouter) OnError(err error) {

}

func (this *ConnRouter) OnMessage(request *WellRequest) {

}

func (this *ConnRouter) OnClose(conn *WellConnection) {

}
