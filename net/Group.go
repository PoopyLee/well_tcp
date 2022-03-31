package net

import (
	"github.com/lvwei25/well_tcp/log"
	"strings"
	"sync"
)

//全局分组管理
var groups = struct {
	sync.RWMutex
	m map[string][]WellConnection
}{m: make(map[string][]WellConnection, 1)}

type Group struct {
}

//将连接对象添加到某个组中
func (this *Group) AddGroup(GroupName string, Con WellConnection) {
	groups.RLock()
	value, ok := groups.m[GroupName]
	groups.RUnlock()
	if !ok {
		value = make([]WellConnection, 1)
	}
	value = append(value, Con)
	groups.Lock()
	groups.m[GroupName] = value
	groups.Unlock()
}

//向某个组的所有成员发送信息
func (this *Group) SendToGroup(GroupName string, Data string) {
	name := make(chan string, 1)
	data := make(chan string, 1)
	name <- GroupName
	data <- Data
	defer close(name)
	defer close(data)
	go func() {
		groups.RLock()
		defer groups.RUnlock()
		for k, v := range groups.m {
			if strings.Compare(k, <-name) == 0 {
				for _, m := range v {
					m.WriteString(<-data)
				}
				return
			}
		}
		log.NewLoger().Error("GroupName Not Exits,Send To ", GroupName, " Failed!")
	}()
}

//移除某个组中的某个连接(并不会关掉连接)
func (this *Group) DelGroupConn(GroupName string, ConnId int64) {
	name := make(chan string, 1)
	id := make(chan int64, 1)
	id <- ConnId
	name <- GroupName
	defer close(name)
	defer close(id)
	go func() {
		groups.RLock()
		defer groups.RUnlock()
		for k, v := range groups.m {
			if strings.Compare(k, <-name) == 0 {
				for n, m := range v {
					if m.ConnId == <-id {
						groups.RLock()
						value, ok := groups.m[k]
						groups.RUnlock()
						if !ok {
							value = make([]WellConnection, 1)
						}
						value = append(value[:n], value[n+1:]...)
						groups.Lock()
						groups.m[k] = value
						groups.Unlock()
					}
				}
				return
			}
		}
		log.NewLoger().Error("GroupName Or ConnId Not Exits,Delete Failed!")
	}()
}

//删除分组
func (this *Group) DelGroup(GroupName string) {
	groups.Lock()
	defer groups.Unlock()
	delete(groups.m, GroupName)
}

//删除连接ID的连接
func (this *Group) delGroup(Id int64) {
	id := make(chan int64, 1)
	id <- Id
	defer close(id)
	go func() {
		groups.RLock()
		defer groups.RUnlock()
		len := len(groups.m)
		if len <= 0 {
			return
		}
		groups.RLock()
		defer groups.RUnlock()
		for k, v := range groups.m {
			for n, m := range v {
				if m.ConnId == <-id {
					groups.RLock()
					value, ok := groups.m[k]
					groups.RUnlock()
					if !ok {
						value = make([]WellConnection, 1)
					}
					value = append(value[:n], value[n+1:]...)
					groups.Lock()
					groups.m[k] = value
					groups.Unlock()
					return
				}
			}
		}
		log.NewLoger().Error("Id Not Exits,Delete Failed!")
	}()
}

//获取所有分组的名字
func (this *Group) GetAllGroup() []string {
	c := make(chan []string, 1)
	defer close(c)
	go func() {
		groups.RLock()
		defer groups.RUnlock()
		g := make([]string, 1)
		for k, _ := range groups.m {
			g = append(g, k)
		}
		c <- g
	}()
	return <-c
}

//获取指定分组的所有连接
func (this *Group) GetAllGroupCon(GroupName string) []WellConnection {
	c := make(chan []WellConnection, 1)
	defer close(c)
	go func() {
		groups.RLock()
		defer groups.RUnlock()
		g := make([]WellConnection, 1)
		for k, v := range groups.m {
			if strings.Compare(k, GroupName) == 0 {
				for _, m := range v {
					g = append(g, m)
				}
			}
		}
		c <- g
	}()
	return <-c
}
