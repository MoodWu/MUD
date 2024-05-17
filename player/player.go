package player

import (
	"mud/common"
	"mud/item"
	"strings"
	"time"
)

// 创建一个用户
func NewPlayer(name, nickname, passwd string) *Player {
	player := Player{
		Name:     name,
		NickName: nickname,
		Passwd:   passwd,
		Age:      1,
		Scene:    nil,
		Invtory:  make([]*item.Thing, 0),
	}

	//todo:给玩家默认物品

	return &player
}

// 验证密码
func (p *Player) CheckPass(pass string) bool {
	return true
}

// 上线时,输出时间,并描述当前场景
func (p *Player) Online() string {
	var ret strings.Builder

	ret.WriteString(time.Now().Format("2006-01-02 15:04:05") + "," + p.NickName + "上线\n")
	ret.WriteString(p.Scene.Desc)

	return ret.String()
}

// 根据动作名称找到命令的说明
func (p *Player) FindActionByName(actionName string) []string {
	result := make([]string, 0)
	//遍历所有Item，查看Item是否实现了对应的action
	for _, i := range p.Invtory {
		result = append(result, getItemActionDesc(i.Thing, p, actionName)...)
	}

	// //遍历所属场景的路径，看看是否是移动命令
	// for _, p := range p.Scene.Paths {
	// 	if p.Direction == actionName {
	// 		result = append(result, "运行 "+p.Direction+" 进行移动")
	// 	}
	// }

	//遍历所属场景的物品，看看是否有合适的动作
	for _, i := range p.Scene.Items {
		result = append(result, getItemActionDesc(i.Thing, p, actionName)...)
	}
	return result
}

// 判断物品上动作的前置条件
func evalute(condition string, player *Player, item *item.NormalItem, args ...string) bool {
	ret := false
	switch condition {
	case "hasOwner":
		ret = item.Owner != ""
	case "hasNoOwner":
		ret = item.Owner == ""
	case "containsItem":
		//先遍历人的物品
		for _, v := range player.Invtory {
			if v.Thing.GetItemName() == args[0] {
				ret = true
				break
			}
		}

		//检查场景的物品
		for _, v := range player.Scene.Items {
			if v.Thing.GetItemName() == args[0] {
				ret = true
				break
			}
		}
	default:
	}

	return ret
}

// 找到Item上所有可以执行的动作说明
func getItemActionDesc(i item.Item, p *Player, actionName string) []string {
	var result []string
	normalItem, ok := i.(*item.NormalItem)
	if !ok {
		return result
	}

	for _, action := range normalItem.Actions {
		bFlag := false
		condition := action.GetCondition()
		if condition == "" {
			bFlag = true
		} else {
			cmds := strings.Split(condition, "(")
			cmd := cmds[0]
			var args []string
			if len(cmds) == 2 {
				args = strings.Split(strings.TrimSuffix(cmds[1], ")"), ",")
			}
			bFlag = evalute(cmd, p, normalItem, args...)
		}
		if bFlag && action.GetActionName() == actionName {
			result = append(result, action.GetActionDesc())
		}
	}
	return result
}

// 获取一个物品上的某个动作
func getItemAction(i item.Item, p *Player, actionName string) common.Action {
	normalItem, ok := i.(*item.NormalItem)
	if !ok {
		return nil
	}

	for _, action := range normalItem.Actions {
		bFlag := false
		condition := action.GetCondition()
		if condition == "" {
			bFlag = true
		} else {
			cmds := strings.Split(condition, "(")
			cmd := cmds[0]
			var args []string
			if len(cmds) == 2 {
				args = strings.Split(strings.TrimSuffix(cmds[1], ")"), ",")
			}
			bFlag = evalute(cmd, p, normalItem, args...)
		}
		if bFlag && action.GetActionName() == actionName {
			return action
		}
	}
	return nil
}

// 目前简化处理，认为命令都是 动宾结构，可以加上补语，这样有，如果命令字符串只有一个部分，一定是系统命令（chat命令除外），有两个以上，第二个就是可以执行动作的物品(chat除外)
func (p *Player) Process(data string) string {
	ret := "什么？"
	commands := strings.Split(data, " ")
	cmd := commands[0]

	if len(commands) == 1 || isSystemCommand(cmd) {
		return p.ProcessSystemCommand(data)
	}

	object := commands[1]
	var action common.Action = nil
	//先找物品
	//遍历所有Item，查看Item是否实现了对应的action
	for _, i := range p.Invtory {
		if i.Thing.GetItemName() == object {
			action = getItemAction(i.Thing, p, cmd)
			if action != nil {
				ret = p.PerformAction(action)
			}
		}
	}

	return ret
}

func isSystemCommand(data string) bool {
	return false
}

// 处理系统命令
func (p *Player) ProcessSystemCommand(data string) string {
	commands := strings.Split(data, " ")
	cmd := commands[0]
	switch cmd {
	case "look":
		return p.ShowMap() + "\n" + p.Scene.Desc
	}

	return "什么？"
}

// 执行动作
func (p *Player) PerformAction(action common.Action) string {
	//todo 先判断这个动作是否有前置条件

	//执行动作
	result := action.Excute()

	//解析结果
	for _, v := range result {
		switch strings.ToLower(v.Object) {
		case "player":
			p.ApplyPlayerResult(v.Effect)
		case "map":
			p.ApplyMapResult(v.Effect)
		}
	}
	return action.GetActionEffectDesc()
}

func (p *Player) ApplyPlayerResult(data map[string]interface{}) {

}

func (p *Player) ApplyMapResult(data map[string]interface{}) {

}

func (p *Player) ShowMap() string {
	return p.Scene.Map.ShowMap()
}
