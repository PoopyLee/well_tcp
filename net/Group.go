package net

import (
	"Well/log"
	"sync"
)

//全局分组管理
var groups = struct {
	sync.RWMutex
	m map[Group][]*WellConnection
}{m: make(map[Group][]*WellConnection, 1)}

type Group struct {
	groupId   int64  //分组ID
	groupName string //分组名称
}

//将连接对象添加到某个组中
func (this *Group) AddGroup(GroupId int64, GroupName string, Con WellConnection) {
	s := Group{
		groupId:   GroupId,
		groupName: GroupName,
	}
	groups.RLock()
	value, ok := groups.m[s]
	groups.RUnlock()
	if !ok {
		value = make([]*WellConnection, 0, 2)
		log.NewLoger().Warn(log.RunFuncName(), "--->GroupId Exits,Add Group Failed!")
		return
	}
	value = append(value, &Con)
	groups.Lock()
	groups.m[s] = value
	groups.Unlock()
}

//向某个组的所有成员发送信息
func (this *Group) SendToGId(GroupId int64, Data string) {
	go func() {
		for k, v := range groups.m {
			if k.groupId == GroupId {
				for _, m := range v {
					m.WriteString(Data)
				}
				return
			}
		}
		log.NewLoger().Error(log.RunFuncName(), "--->GroupId Not Exits,Send To ", GroupId, " Failed!")
	}()
}

//删除某个组中的某个连接
func (this *Group) DelGroupConn(GroupId int64, Conn *WellConnection) {
	go func() {
		for k, v := range groups.m {
			if k.groupId == GroupId {
				for n, m := range v {
					if m.ConnId == Conn.ConnId {
						groups.RLock()
						value, ok := groups.m[k]
						groups.RUnlock()
						if !ok {
							value = make([]*WellConnection, 0, 2)
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
func (this *Group) DelGroup(GroupId int64, GroupName string) {
	s := Group{
		groupId:   GroupId,
		groupName: GroupName,
	}
	delete(groups.m, s)
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
						value = make([]*WellConnection, 0, 2)
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
func (this *Group) GetAllGroup() []Group {
	g := make([]Group, 1)
	go func() {
		for k, _ := range groups.m {
			g = append(g, k)
		}
	}()
	return g
}
