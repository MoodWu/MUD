package player

import (
	"mud/item"
	"mud/maps"
)

type Player struct {
	Name     string
	NickName string
	Passwd   string
	Age      int
	Scene    *maps.Scene
	Invtory  []*item.Thing
}
