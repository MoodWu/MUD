package engine

import (
	"fmt"
)

var commands map[string]ICommand

func init() {
	commands = make(map[string]ICommand, 0)
}

func RegisterCMD(cmd string, processor ICommand) {
	commands[cmd] = processor
}

func ProcessCMD(command Command) (string, error) {
	processor, ok := commands[command.CMD]
	if !ok {
		return "", fmt.Errorf("command %s can not be recongnized.", command.CMD)
	}

	return processor.Process(command.Data)
}

// 注册系统级别的命令，login，loginpwd，passwd
func initSystemCommand() {

}
