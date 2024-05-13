package player

import (
	"mud/item"
)

// 创建一个用户
func NewPlayer(name, nickname, passwd string) *Player {
	player := Player{
		Name:     name,
		NickName: nickname,
		Passwd:   passwd,
		Age:      1,
		Scene:    nil,
		Invtory:  make([]*item.Item, 0),
	}

	//todo:给玩家默认物品

	return &player
}

// 加载一个已有的用户
func LoadPlayer(name string) *Player {

	return nil
}

// 验证密码
func (p *Player) CheckPass(pass string) bool {
	return true
}
