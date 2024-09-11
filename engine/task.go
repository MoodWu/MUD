package engine

type Task struct{
	Name string
	Code string
	Data map[string]string
	Done bool
	Script string
	SubTask []TaskData
}

type TaskData struct{
	Name string
	Code string
	Done string
	Script string
}

func (t *Task) OnProgress(key string,player *Player) {
	//根据Key找到当前所属的
	//遍历所有没有完成的SubTask
}