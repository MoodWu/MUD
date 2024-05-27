package engine

import (
	"encoding/json"
	"fmt"
)

var commands map[string]ICommand

func init() {
	commands = make(map[string]ICommand, 0)
}

func RegisterCMD(cmd string, processor ICommand) {
	commands[cmd] = processor
}

func ProcessCMD(command Command) ([]byte, error) {
	processor, ok := commands[command.CMD]
	if !ok {
		return nil, fmt.Errorf("command %s can not be recongnized.", command.CMD)
	}

	return processor.Process(command)
}

// 注册系统级别的命令，login，loginpwd，passwd
func initSystemCommand() {

	RegisterCMD("login", &LoginCMD{})
}

type LoginCMD struct{}

func (c *LoginCMD) Process(cmd Command) ([]byte, error) {
	ret := Command{CMD: "login", Data: "请输入用户名："}
	d, _ := json.Marshal(ret)
	return d, nil
}

type LoginPwdCMD struct{}

func (c *LoginPwdCMD) Process(cmd Command) ([]byte, error) {
	ret := Command{CMD: "loginpwd", Data: "请输入密码："}
	d, _ := json.Marshal(ret)
	return d, nil
}

var MapList map[string]*Map

func init() {
	m := InitMap()
	MapList = make(map[string]*Map, 0)
	MapList[m.Code] = m
}

// 加载一个已有的用户
func LoadPlayer(name string) *Player {
	hp := PlayerStatus{HP: 100, MaxHP: 100}
	player := Player{
		Name:         name,
		NickName:     name,
		Passwd:       "123",
		Age:          1,
		Scene:        nil,
		Inventory:    make([]*Goods, 0),
		PlayerStatus: hp,
	}

	player.Scene = MapList["m00001"].Scenes[0]

	//todo:给玩家默认物品

	return &player
}
