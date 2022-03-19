package net

//信息
type WellMessage struct {
	Data    []byte
	DataLen uint32
}

type WellMessageInterface interface {
	GetDataLen() uint32 //获取消息数据段长度
	GetData() []byte    //获取消息内容

	SetData([]byte)    //设置消息内容
	SetDataLen(uint32) //设置消息数据段长度
}

//获取消息数据段长度
func (this *WellMessage) GetDataLen() uint32 {
	return this.DataLen
}

//获取消息内容
func (this *WellMessage) GetData() []byte {
	return this.Data
}

//设置消息数据段长度
func (this *WellMessage) SetDataLen(len uint32) {
	this.DataLen = len
}

//设计消息内容
func (this *WellMessage) SetData(data []byte) {
	this.Data = data
}
