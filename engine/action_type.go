package engine

// 探查接口，暴露出可以执行的动作
type Detectable interface {
	GetActions() []Action
}

// 动作接口
type Action interface {
	GetActionName() string
	GetActionDesc() string
	GetScript() string
}
