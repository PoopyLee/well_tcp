package net

import (
	"github.com/lvwei25/well_tcp/log"
	"sync"
)

//全局
var links = struct {
	sync.RWMutex
	m map[int64]*WellConnection
}{m: make(map[int64]*WellConnection, 10)}

type LinkManger struct {
}

func (this *LinkManger) addLink(Id int64, Con *WellConnection) {
	links.Lock()
	defer links.Unlock()
	links.m[Id] = Con

}

func (this *LinkManger) deletLink(Id int64) {
	links.Lock()
	defer links.Unlock()
	delete(links.m, Id)
}

func (this *LinkManger) GetLink(Id int64) *WellConnection {
	links.RLock()
	defer links.RUnlock()
	return links.m[Id]
}

func (this *LinkManger) GetAllLink() map[int64]*WellConnection {
	links.RLock()
	defer links.RUnlock()
	return links.m
}

func (this *LinkManger) SendToAll(Data string) {
	go func() {
		links.RLock()
		defer links.RUnlock()
		for _, v := range this.GetAllLink() {
			v.WriteString(Data)
		}
	}()
}

func (this *LinkManger) SendToId(Id int64, Data string) {
	go func() {
		links.RLock()
		defer links.RUnlock()
		for _, v := range this.GetAllLink() {
			if v.ConnId == Id {
				v.WriteString(Data)
				return
			}
		}
		log.NewLoger().Error("--->Id Not Exits,Send To ", Id, "Failed!")
	}()
}
