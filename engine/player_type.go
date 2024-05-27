package engine

type Player struct {
	Name      string
	NickName  string
	Passwd    string
	Age       int
	Scene     *Scene
	Inventory []*Goods
	PlayerStatus
}

// 玩家数据
type PlayerStatus struct {
	HP    int
	MaxHP int
	Level int
	Score int
	Goal  int
	JobID int
}
