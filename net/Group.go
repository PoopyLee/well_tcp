package net

import (
	"Well/log"
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
func (this *Group) SendToGId(GroupName string, Data string) {
	go func() {
		for k, v := range groups.m {
			if strings.Compare(k, GroupName) == 0 {
				for _, m := range v {
					m.WriteString(Data)
				}
				return
			}
		}
		log.NewLoger().Error(log.RunFuncName(), "--->GroupId Not Exits,Send To ", GroupName, " Failed!")
	}()
}

//删除某个组中的某个连接
func (this *Group) DelGroupConn(GroupName string, Conn *WellConnection) {
	go func() {
		for k, v := range groups.m {
			if strings.Compare(k, GroupName) == 0 {
				for n, m := range v {
					if m.ConnId == Conn.ConnId {
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
		log.NewLoger().Error(log.RunFuncName(), "--->GroupId Not Exits,Delete Failed!")
	}()
}

//删除指定分组ID的分组
func (this *Group) DelGroup(GroupName string) {
	delete(groups.m, GroupName)
}

//删除连接ID的连接
func (this *Group) delGroup(Id int64) {
	go func() {
		for k, v := range groups.m {
			for n, m := range v {
				if m.ConnId == Id {
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
		log.NewLoger().Error(log.RunFuncName(), "--->Id Not Exits,Delete Failed!")
	}()
}

//获取所有分组的ID和Name
func (this *Group) GetAllGroup() []string {
	c := make(chan []string, 1)
	defer close(c)
	go func() {
		g := make([]string, 1)
		for k, _ := range groups.m {
			g = append(g, k)
		}
		c <- g
	}()
	return <-c
}

//获取所有分组的ID和Name
func (this *Group) GetAllGroupCon(GroupName string) []WellConnection {
	c := make(chan []WellConnection, 1)
	defer close(c)
	go func() {
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
