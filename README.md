# Well - Golang 轻量级TCP服务器
<p align="center">
    <a href="/" target="_blank" style="text-align: center">
        <img src="https://github.com/lvwei25/well_tcp/blob/main/logo.png" style="width: 100px;height: 100px" alt="Well_TCP" />
    </a>
</p>

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

##开发初衷
本人是即将大四毕业老狗，专业是嵌入式方面的，由于毕设涉及到了服务器，并且是TCP，所以就开发了Well服务器。
目前Well不是很完善，但是基本功能已经齐全，能够满足基本开发，性能上由于Golang的特性，所以性能差不到哪里去。。

