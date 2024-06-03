package main

type Command struct {
	CMD  string `json:"cmd"`
	Data string `json:"data"`
}

type SceneInfo struct {
	X    int
	Y    int
	Name string
	Code string
	Path map[string]string
}
type MapInfo struct {
	Name   string
	Code   string
	Long   int
	Width  int
	Scenes []SceneInfo
	X      int
	Y      int
}
type MapDetail struct {
	MapInfo   *MapInfo
	SceneDesc string
}
type World struct {
	Maps []*Unit
}
type Unit struct {
	X   int
	Y   int
	Map *Map
}

type Player struct {
	Name      string
	NickName  string
	Passwd    string
	Age       int
	Scene     *Scene
	Map       *Map
	Inventory []Item
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
	X     int
	Y     int
	Name  string
	Code  string // 唯一编码
	Desc  string
	Items []*Goods
	Path  map[string]string
}

type Item interface {
	GetItemName() string
	GetItemQuantity() int
	GetItemDesc() string
	GetItemDetail() string
	SetItemQuantity(v int) int
}

type Goods struct {
	Item     Item
	Quantity int
}
