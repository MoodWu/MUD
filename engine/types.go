package engine

type Command struct {
	CMD  string `json:"cmd"`
	Data string `json:"data"`
}

// 所有命令都应该实现此接口Pro
type ICommand interface {
	Process(data string) (string, error)
}
