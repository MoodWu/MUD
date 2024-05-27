package engine

type NormalAction struct {
	Name   string
	Desc   string
	Detail string
	Script string
}

func (item *NormalAction) GetActionName() string {
	return item.Name
}

func (item *NormalAction) GetActionDesc() string {
	return item.Desc
}

func (item *NormalAction) GetScript() string {
	return item.Script
}
