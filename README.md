# Well - Golang 轻量级TCP服务器
<p align="center">
    <a href="/" target="_blank" style="text-align: center">
        <img src="https://github.com/lvwei25/well_tcp/blob/main/logo.png" style="width: 100px;height: 100px" alt="Well_TCP" />
    </a>
</p>



![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lvwei25/well_tcp)
![GitHub AppVersion](https://img.shields.io/badge/Version-V1.0-blue)




##介绍

Well是一款基于Golang的轻量级TCP服务器，其为用户内置了路由模块、消息模块、链接模块、分组模块、日志模块。
<br>
只需要简单的两行代码即可启动服务:

````go
	s:=net.NewServerHandle("This is Well Server!","127.0.0.1","8888")
	s.AddServerRouter(&ServerRouter{})  //服务启动期间的路由配置
	s.AddConnRouter(&ConnRouter{})      //连接的路由配置
	s.Run() //开始监听服务器
````

<p align="center">
    <a href="/" target="_blank" style="text-align: center">
        <img src="https://github.com/lvwei25/well_tcp/blob/main/test_img/run.jpg" alt="Well_TCP" />
    </a>
</p>

> Well 第一个版本发布于 2022 年 3 月 19 日


###Server（服务）


```text
属性
    Name（服务名称）
    IpAddr（服务监听地址）
    Port（监听端口）
    ServerRouter（服务路由）
    ConnRouter（连接路由）
    Version(服务版本)
方法
    Run（服务启动）
    Close（关闭服务）
``` 
        
        
###Connection（链接）


```text
属性
    ID（连接唯一ID）
    IpAddr（连接ip地址）
    Port（连接端口）
    connRouter（连接路由）
    con （net.Conn）
    isClose（是否关闭）
方法
    Start（开始运行）
    WriteString（向客户端 写入字符串）
    WriteByte（向客户端写入字节）
    Close（关闭连接）
```


###Router（路由）


```text
ServerRouter（服务路由）
    方法
        OnStart（服务启动后，监听前执行）
        OnClose（服务关闭前执行）
    属性（无属性，用户继承该结构体便有其拥有的方法）
ConnRouter（连接路由）
    方法
        OnConnect（连接后第一次执行）
        OnMessage（接收到数据时执行）
        OnClose（连接关闭时执行）
        OnError（连接出错时执行）
    属性（无属性，用户继承该结构体便有其拥有的方法）
```



###Links（链接管理）


```text
属性
    m map[int64]*WellConnection
方法
    addLink(Id int64,Con *WellConnection)添加连接
    deletLink(Id int64)删除连接
    GetLink(Id int64) *WellConnection 获取连接
    GetAllLink()map[int64]*WellConnection 获取全部连接
    SendToAll(Data string) 向所有的连接发送数据
    SendToId(Id int64,Data string) 向指定的连接ID发送数据
```



###Groups（分组管理）


```text
属性
    groupId int64	//分组ID
    groupName string	//分组名称
方法
    AddGroup(GroupId int64,GroupName string,Con WellConnection) 将某个连接添加到分组中
    SendToGId(GroupId int64,Data string) 向某个分组发送数据
    DelGroupConn(GroupId int64,Conn *WellConnection) 删除某个组中的某个连接
    DelGroup(GroupId int64,GroupName string) 删除分组
    delGroup(Id int64)删除指定连接ID的分组
```



###Log（日志打印）


```text
属性
    prefix 打印前缀
方法
    Println(a...interface{}) 换行打印
    Debug(a ...interface{}) 调试打印
    Info(a ...interface{})信息打印
    Warn(a ...interface{}) 警告打印
    Error(a ...interface{}) 错误打印
    Panic(a ...interface{}) Panic打印
    SetPrefix(Prefix string) 设置打印前缀
    SetLogFile(File string) 设置日志输出的文件路径
```


###Message（信息）


```text
属性
    Data 数据
    DataLen 数据长度
方法
    GetDataLen() uint32    //获取消息数据段长度
    GetData() []byte    //获取消息内容
    SetData([]byte)        //设置消息内容
    SetDataLen(uint32)    //设置消息数据段长度
```



##开发初衷
本人是即将大四毕业老狗，专业是嵌入式方面的，由于毕设涉及到了服务器，并且是TCP，所以就开发了Well服务器。
目前Well不是很完善，但是基本功能已经齐全，能够满足基本开发，性能上由于Golang的特性，所以性能差不到哪里去。


