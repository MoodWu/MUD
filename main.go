package main

import (
	"fmt"
	"mud/engine"
)

func main() {
	fmt.Println("hello")

	//初始化世界
	//InitWorld()

	//初始化WebSocketServer
	engine.InitWebSocket()
}

// 初始化世界
func InitWorld() *engine.World {
	m0 := engine.InitMap()
	u0 := &engine.Unit{X: 0, Y: 0, Map: m0}
	world := &engine.World{}
	world.Maps = make([]*engine.Unit, 0)
	world.Maps = append(world.Maps, u0)

	return world
}
