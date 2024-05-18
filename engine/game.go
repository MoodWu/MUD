package engine

import (
	"encoding/json"
	"fmt"
	"mud/item"
	"mud/maps"
	"mud/player"
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

var MapList map[string]*maps.Map

func init() {
	m := maps.InitMap()
	MapList = make(map[string]*maps.Map, 0)
	MapList[m.Code] = m
}

// 加载一个已有的用户
func LoadPlayer(name string) *player.Player {
	player := player.Player{
		Name:     name,
		NickName: name,
		Passwd:   "123",
		Age:      1,
		Scene:    nil,
		Invtory:  make([]*item.Thing, 0),
	}

	player.Scene = MapList["m00001"].Scenes[0]

	//todo:给玩家默认物品

	return &player
}
