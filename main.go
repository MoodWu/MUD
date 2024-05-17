package main

import (
	"fmt"
	"mud/engine"
	"mud/maps"
)

func main() {
	fmt.Println("hello")

	//初始化世界
	//InitWorld()

	//初始化WebSocketServer
	engine.InitWebSocket()
}

// 初始化世界
func InitWorld() *maps.World {
	m0 := maps.InitMap()
	u0 := &maps.Unit{X: 0, Y: 0, Map: m0}
	world := &maps.World{}
	world.Maps = make([]*maps.Unit, 0)
	world.Maps = append(world.Maps, u0)

	return world
}
