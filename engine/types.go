package engine

import (
	"mud/player"

	"github.com/gorilla/websocket"
)

type Command struct {
	CMD  string `json:"cmd"`
	Data string `json:"data"`
}

// 所有命令都应该实现此接口Pro
type ICommand interface {
	Process(cmd Command) ([]byte, error)
}

type WebSocket struct {
	Conn   *websocket.Conn
	Player *player.Player
}
