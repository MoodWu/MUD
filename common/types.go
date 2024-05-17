package common

// 探查接口，暴露出可以执行的动作
type Detectable interface {
	GetActions() []Action
}

// 动作接口
type Action interface {
	Excute(args ...interface{}) []ActionResult
	GetActionName() string
	GetActionDesc() string
	GetCondition() string //执行动作的前置条件
	GetActionEffect() string
	GetActionEffectDesc() string
}

// 动作执行结果
type ActionResult struct {
	Object string
	Effect map[string]interface{}
}

/*定义地图格式*/
type YAML_Map struct {
	Name      string       `yaml:"map"`
	Code      string       `yaml:"code"`
	Desc      string       `yaml:"desc"`
	Scale     string       `yaml:"scale"`
	InitPoint string       `yaml:"initpoint"`
	Scenes    []YAML_Scene `yaml:"scenes"`
}

type YAML_Scene struct {
	Name      string            `yaml:"name"`
	Code      string            `yaml:"code"`
	Desc      string            `yaml:"desc"`
	Position  string            `yaml:"position"`
	Items     []YAML_Thing      `yaml:"items"`
	Direction map[string]string `yaml:"path"`
}

type YAML_Thing struct {
	Item     []YAML_Item `yaml:"item"`
	Quantity int         `yaml:"qty"`
}

type YAML_Item struct {
	Name     string        `yaml:"name"`
	Desc     string        `yaml:"desc"`
	Quantity int           `yaml:"qty"`
	Actions  []YAML_Action `yaml:"actions"`
}

type YAML_Action struct {
	Name       string `yaml:"cmd"`
	Desc       string `yaml:"desc"`
	Effect     string `yaml:"effect"`
	EffectDesc string `yaml:"effectDesc"`
	Condition  string `yaml:"condition"`
}

type MoveAction struct {
	Direction string
}

func (m *MoveAction) Excute(args ...interface{}) []ActionResult {

	return nil
}
