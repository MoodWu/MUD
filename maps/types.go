package maps

import "mud/item"

type World struct {
	Maps []*Unit
}
type Unit struct {
	X   int
	Y   int
	Map *Map
}
type Map struct {
	Name        string        // 地图名称
	Desc        string        // 地图描述
	Width       int           // 地图长度
	Length      int           //地图宽度
	RefreshTime int           //物品刷新时间
	Scenes      []*Scene      //地图上的场景
	Connections []*Connection //地图与其他地图的连接
}

type Connection struct {
	Direction string
	Dest      *Map
}

type Path struct {
	Direction string
	Dest      *Scene
}
type Scene struct {
	Map   *Map // 所属地图
	X     int
	Y     int
	Name  string
	Desc  string
	Items []*item.Item
	Paths []*Path
}
