package net

import (
	"net"
	"strings"
)

var tempId int64

type WellConnection struct {
	ConnId     int64
	IpAddr     string
	Port       string
	connRouter WellConnRouter
	con        *net.TCPConn
	isClose    chan bool
	Links      LinkManger
	Groups     Group
}

type wellConnInterface interface {
	Start()                  //开始读取
	WriteString(Data string) //写数据
	WriteByte(Data []byte)
	Close()

	readData()
	writeData()
	setConnRouter(c WellConnRouter)
}

func (this *WellConnection) Start() {
	this.connRouter.OnConnect(this)
	go this.readData()
	//go this.writeData()
	select {
	case <-this.isClose:
		this.con.Close()
		this.Links.deletLink(this.ConnId)
		this.Groups.delGroup(this.ConnId)
		return
	}
}

func (this *WellConnection) WriteString(Data string) {
	this.con.Write([]byte(Data))
}

func (this *WellConnection) WriteByte(Data []byte) {
	this.con.Write(Data)
}

func (this *WellConnection) readData() {
	for {
		req := new(WellRequest)
		byte := make([]byte, 2048)
		cnt, err := this.con.Read(byte)
		if err != nil {
			this.isClose <- true
		} else {
			req.SetData(byte[:cnt])
			req.SetDataLen(uint32(cnt))
			req.WellConnection = *this
			this.connRouter.OnMessage(req)
		}
	}

}

//TODO	不实现
func (this *WellConnection) writeData() {

}

func (this *WellConnection) Close() {
	this.connRouter.OnClose(this)
	this.Links.deletLink(this.ConnId)
	this.Groups.delGroup(this.ConnId)
	this.con.Close()
}

//设置路由
func (this *WellConnection) setConnRouter(c WellConnRouter) {
	this.connRouter = c
}

func NewConnHandle(con *net.TCPConn) wellConnInterface {
	tempId++
	c := WellConnection{
		ConnId:  tempId,
		IpAddr:  "",
		Port:    "",
		con:     con,
		isClose: make(chan bool),
	}
	ips := strings.Split(con.RemoteAddr().String(), ":")
	c.IpAddr = ips[0]
	c.Port = ips[1]
	c.Links.addLink(c.ConnId, &c)
	return &c
}
