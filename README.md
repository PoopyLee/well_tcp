# Well_TCP - Golang 轻量级TCP嵌入式服务器
<p align="center">
    <a href="https://github.com/lvwei25/well_tcp/blob/main/logo.png" target="_blank" style="text-align: center">
        <img src="https://github.com/lvwei25/well_tcp/blob/main/logo.png"  alt="Well_TCP" />
    </a>
</p>

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lvwei25/well_tcp)
![GitHub AppVersion](https://img.shields.io/badge/Version-V1.0-blue)

## 介绍

Well_TCP是一款基于Golang的轻量级TCP嵌入式服务器，其为用户内置了路由模块、消息模块、链接模块、分组模块、日志模块，能够满足日常的开发。
<br>

###### tips:比如目前本人使用该框架开发的“xxx智慧监控系统”，能够有效的处理各种事件！

<a href="https://github.com/lvwei25/well_tcp_demo">详情点击“Well_Tcp开发项目示例”此处</a>

## 使用go get下载依赖

````go
 go get -v -u github.com/lvwei25/well_tcp
````

只需要简单的两行代码即可启动服务:

````go
	s:=net.NewServerHandle("This is Well Server!","127.0.0.1","8888")
	s.AddServerRouter(&ServerRouter{})  //服务启动期间的路由配置
	s.AddConnRouter(&ConnRouter{})      //连接的路由配置
	s.Run() //开始监听服务器
````

<p align="center">
    <a href="https://github.com/lvwei25/well_tcp/blob/main/test_img/test_img.png" target="_blank" style="text-align: center">
        <img src="https://github.com/lvwei25/well_tcp/blob/main/test_img/test_img.png" alt="Well_TCP" />
    </a>
</p>

> Well 第一个版本发布于 2022 年 3 月 19 日


### Server（服务）


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
        
        
### Connection（链接）


```text
属性
    ID（连接唯一ID）
    IpAddr（连接ip地址）
    Port（连接端口）
方法
    Start（开始运行）
    WriteString（向客户端 写入字符串）
    WriteByte（向客户端写入字节）
    Close（关闭连接）
```


### Router（路由）


```text
ServerRouter（服务路由）
    方法
        OnStart()（服务启动后，监听前执行）
        OnClose()（服务关闭前执行）
    属性（无属性，用户继承该结构体便有其拥有的方法）
ConnRouter（连接路由）
    方法
        OnConnect(conn *WellConnection)（连接后第一次执行）
        OnMessage(request *WellRequest)（接收到数据时执行）
        OnClose(conn *WellConnection)（连接关闭时执行）
        OnError(err error)（连接出错时执行）
    属性（无属性，用户继承该结构体便有其拥有的方法）
```



### Links（链接管理）


```text
属性
    m map[int64]*WellConnection
方法
    GetLink(Id int64) *WellConnection 获取连接
    GetAllLink()map[int64]*WellConnection 获取全部连接
    SendToAll(Data string) 向所有的连接发送数据
    SendToId(Id int64,Data string) 向指定的连接ID发送数据
```



### Groups（分组管理）


```text
属性
    无
方法
    AddGroup(GroupName string,Con WellConnection) 将某个连接添加到分组中
    SendToGroup(GroupName string, Data string) 向某个组的所有成员发送信息
    DelGroupConn(GroupName string, ConnId int64) 移除某个组中的某个连接(并不会关掉连接)
    DelGroup(GroupName string) 删除分组
    GetAllGroupCon(GroupName string) []WellConnection   获取指定分组的所有连接
    GetAllGroup() []string 获取所有分组的名字
```



### Log（日志打印）


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


### Message（信息）


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


## 注意

关于连接管理和分组管理：当某个连接断开时，会触发相应连接的OnClose()事件，用户不需要手动对该连接所在分组或连接器进行移除。
Well_TCP已经帮你实现了。

## 开发初衷
本人是即将大四毕业老狗，专业是嵌入式方面的，由于毕设涉及到了服务器，并且是TCP，所以就开发了Well_TCP服务器。
目前Well_TCP不是很完善，但是基本功能已经齐全，能够满足基本开发，性能上由于Golang的特性，所以性能差不到哪里去。


