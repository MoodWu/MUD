package common

import "fmt"

type NoramlAction struct {
	Name       string
	Desc       string
	Effect     string
	EffectDesc string
	Condition  string
}

func (item *NoramlAction) Excute(args ...interface{}) []ActionResult {
	fmt.Println(item.Name+" executed with args ", args)
	return nil
}

func (item *NoramlAction) GetActionName() string {
	return item.Name
}

func (item *NoramlAction) GetActionDesc() string {
	return item.Desc
}

func (item *NoramlAction) GetCondition() string {
	return item.Condition
}

func (item *NoramlAction) GetActionEffect() string {
	return item.Effect
}

func (item *NoramlAction) GetActionEffectDesc() string {
	return item.EffectDesc
}
