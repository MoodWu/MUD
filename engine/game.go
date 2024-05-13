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

	return processor.Process(command.Data)
}

// 注册系统级别的命令，login，loginpwd，passwd
func initSystemCommand() {

	RegisterCMD("login", &LoginCMD{})
}

type LoginCMD struct{}

func (c *LoginCMD) Process(data string) ([]byte, error) {
	cmd := Command{CMD: "loginpwd", Data: "enter the password"}
	d, _ := json.Marshal(cmd)
	return d, nil
}

type LoginPwdCMD struct{}

func (c *LoginPwdCMD) Process(data string) ([]byte, error) {
	cmd := Command{CMD: "loginpwd", Data: "enter the password"}
	d, _ := json.Marshal(cmd)
	return d, nil
}
