package engine

import "sync"

type World struct {
	Maps []*Unit
}
type Unit struct {
	X   int
	Y   int
	Map *Map
}
type Map struct {
	Name        string            // 地图名称
	Code        string            //唯一编码
	Desc        string            // 地图描述
	Long        int               // 地图长度
	Wide        int               //地图宽度
	RefreshTime int               //物品刷新时间
	Scenes      []*Scene          //地图上的场景
	Connection  map[string]string //地图与其他地图的连接
}

type Scene struct {
	Map   *Map // 所属地图
	X     int
	Y     int
	Name  string
	Code  string // 唯一编码
	Desc  string
	mu    sync.Mutex
	Items []*Goods
	Path  map[string]string
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
	Name      string               `yaml:"name"`
	Code      string               `yaml:"code"`
	Desc      string               `yaml:"desc"`
	Position  string               `yaml:"position"`
	Items     map[string]YAML_Item `yaml:"items"`
	Direction map[string]string    `yaml:"path"`
}

type YAML_Item struct {
	Name     string        `yaml:",omitempty"`
	Count    int           `yaml:"count"`
	Code     string        `yaml:"code"`
	Desc     string        `yaml:"desc"`
	Detail   string        `yaml:"detail"`
	Quantity int           `yaml:"qty"`
	Actions  []YAML_Action `yaml:"actions"`
}

type YAML_Action struct {
	Name   string `yaml:"cmd"`
	Desc   string `yaml:"desc"`
	Script string `yaml:"script"`
}
