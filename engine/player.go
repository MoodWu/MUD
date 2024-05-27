package engine

import (
	"fmt"
	"log"
	"strings"
	"time"

	lua "github.com/yuin/gopher-lua"
)

var systemCommand map[string]string

func init() {
	systemCommand = make(map[string]string, 0)
	systemCommand["get"] = ""
	systemCommand["drop"] = ""
	systemCommand["inventory"] = ""

}

// 创建一个用户
func NewPlayer(name, nickname, passwd string) *Player {
	player := Player{
		Name:      name,
		NickName:  nickname,
		Passwd:    passwd,
		Age:       1,
		Scene:     nil,
		Inventory: make([]*Goods, 0),
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
	ret.WriteString(p.Scene.GetDesc())

	return ret.String()
}

/*
// 根据动作名称找到命令的说明
func (p *Player) FindActionByName(actionName string) []string {
	result := make([]string, 0)
	//遍历所有Item，查看Item是否实现了对应的action
	for _, i := range p.Inventory {
		result = append(result, getItemActionDesc(i.Item, p, actionName)...)
	}

	// //遍历所属场景的路径，看看是否是移动命令
	// for _, p := range p.Scene.Paths {
	// 	if p.Direction == actionName {
	// 		result = append(result, "运行 "+p.Direction+" 进行移动")
	// 	}
	// }

	//遍历所属场景的物品，看看是否有合适的动作
	for _, i := range p.Scene.Items {
		result = append(result, getItemActionDesc(i.Item, p, actionName)...)
	}
	return result
}

// 判断物品上动作的前置条件
func evalute(condition string, player *Player, item *NormalItem, args ...string) bool {
	ret := false
	switch condition {
	case "hasOwner":
		ret = item.Owner != ""
	case "hasNoOwner":
		ret = item.Owner == ""
	case "containsItem":
		//先遍历人的物品
		for _, v := range player.Inventory {
			if v.Item.GetItemName() == args[0] {
				ret = true
				break
			}
		}

		//检查场景的物品
		for _, v := range player.Scene.Items {
			if v.Item.GetItemName() == args[0] {
				ret = true
				break
			}
		}
	default:
	}

	return ret
}

// 找到Item上所有可以执行的动作说明
func getItemActionDesc(i Item, p *Player, actionName string) []string {
	var result []string
	normalItem, ok := i.(*NormalItem)
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

// 获取一个物品上的某个动作,如果有动作但不满足前置条件，返回错误消息,找不到命令则返回nil
func getItemAction(i Item, p *Player, actionName string) (Action, string) {
	msg := "无法识别的Item"
	normalItem, ok := i.(*NormalItem)
	if !ok {
		return nil, msg
	}

	for _, action := range normalItem.Actions {
		bFlag := false
		if action.GetActionName() != actionName {
			continue
		}

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
		if bFlag {
			return action, action.GetFailMessage()
		}
		return action, ""
	}
	return nil, ""
}
*/

func getItemAction(item Item, p *Player, data string) Action {
	ni, ok := item.(*NormalItem)
	if !ok {
		return nil
	}

	for _, action := range ni.Actions {
		if action.GetActionName() == data {
			return action
		}
	}

	return nil
}

// 执行发来的命令，目前认为单字命令都是系统命令
func (p *Player) Process(data string) string {
	ret := "什么？"
	commands := strings.Split(data, " ")
	cmd := commands[0]

	if len(commands) == 1 || isSystemCommand(cmd) {
		return p.ProcessSystemCommand(data)
	}

	//object := commands[1]

	//先找物品，
	for _, i := range p.Inventory {
		if i.Quantity > 0 {
			action := getItemAction(i.Item, p, data)
			if action != nil {
				ret = p.PerformAction(action, i)
				return ret
			}
		}
	}

	//再找场景物品，
	for _, i := range p.Scene.Items {
		if i.Quantity > 0 {
			action := getItemAction(i.Item, p, data)
			if action != nil {
				ret = p.PerformAction(action, i)
				return ret
			}
		}
	}

	return ret
}

// 获得玩家身上的物品
func (p *Player) getInventoryByName(itemName string) Item {
	for _, i := range p.Inventory {
		if i.Item.GetItemName() == itemName && i.Quantity > 0 {
			return i.Item
		}
	}

	return nil
}

func (p *Player) getSceneItemByName(itemName string) *Goods {
	for _, i := range p.Scene.Items {
		if i.Item.GetItemName() == itemName && i.Quantity > 0 {
			return i
		}
	}

	return nil
}

// 系统命令除了单字命令外，都是动宾结构
func isSystemCommand(data string) bool {
	_, ok := systemCommand[data]

	return ok
}

// 处理系统命令
func (p *Player) ProcessSystemCommand(data string) string {
	commands := strings.Split(data, " ")
	cmd := commands[0]
	ret := "什么？"
	switch cmd {
	case "look":
		//不带参数的look，就是直接看当前环境
		if len(commands) == 1 {
			ret = p.ShowMap() + "\n" + p.Scene.GetDesc()
		} else {
			//查看某个物品
		}
	case "get":
		//捡起物品
		object := commands[1]
		p.Scene.Lock()
		defer p.Scene.Unlock()

		goods := p.getSceneItemByName(object)

		if goods == nil {
			ret = "没有找到" + object
			break
		}
		ni, ok := goods.Item.(*NormalItem)
		if ok {
			if goods.Quantity > 0 {
				//先看是否有此物品
				bFound := false
				for _, item := range p.Inventory {
					if item.Item.GetItemName() == object {
						item.Quantity++
						bFound = true
						break
					}
				}
				//如果没有此物品，添加一个
				if !bFound {
					p.Inventory = append(p.Inventory, &Goods{Item: ni, Quantity: 1})
				}
				goods.Quantity--
				action := getItemAction(goods.Item, p, "get "+ni.Name)
				if action != nil {
					ret = p.PerformAction(action, goods)
				} else {
					ret = "你拾起了" + ni.Name
				}
			} else {
				ret = "手慢了，已经被人捡走了。"
			}
		}
	case "drop":
		//丢弃物品
	case "east":
		ret = "你向东走\n"
		ret += p.moveToScene(cmd)
	case "west":
		ret = "你向西走\n"
		ret += p.moveToScene(cmd)
	case "north":
		ret = "你向北走\n"
		ret += p.moveToScene(cmd)
	case "south":
		ret = "你向南走\n"
		ret += p.moveToScene(cmd)
	case "hp":
		ret = p.getHP()
	case "inventory":
		ret = p.getInventory()
	}

	return ret
}

func (p *Player) getHP() string {
	return fmt.Sprintf("生命值 %d/%d\n", p.HP, p.MaxHP)
}

func (p *Player) getInventory() string {
	ret := ""
	for _, goods := range p.Inventory {
		ret += fmt.Sprintf("物品：%s,数量：%d %s\n", goods.Item.GetItemName(), goods.Quantity, goods.Item.GetItemDetail())
	}
	return ret
}

// 查看物品
func (p *Player) look(item string) string {

	return ""
}

func (p *Player) moveToScene(direction string) string {
	s, ok := p.Scene.Path[direction]
	if ok {
		newScene := p.Scene.Map.GetSceneByCode(s)
		if newScene != nil {
			p.Scene = newScene
			return p.Scene.GetDesc()
		}
	}
	return "前方没有路了。试试其他方向吧。"
}

// 执行动作
func (p *Player) PerformAction(action Action, goods *Goods) string {
	L := luaEngine.Get()
	defer luaEngine.Put(L)

	playerUD := L.NewUserData()
	playerUD.Value = p
	L.SetMetatable(playerUD, L.GetTypeMetatable("Player"))
	L.SetGlobal("player", playerUD)

	itemUD := L.NewUserData()
	itemUD.Value = goods
	L.SetMetatable(itemUD, L.GetTypeMetatable("Goods"))
	L.SetGlobal("goods", itemUD)

	if err := L.DoString(action.GetScript()); err != nil {
		log.Println("err execute lua:", err)
		return ""
	}

	// 获取 Lua 脚本的返回值（字符串）
	ret := L.ToString(-1)
	log.Println("Returned from Lua:", ret)

	return ret
}

func (p *Player) ShowMap() string {
	return p.Scene.Map.ShowMap()
}

/*
以下是为了Lua中都能调用go函数
*/

// 注册player类型
func registPlayer(L *lua.LState) {
	mt := L.NewTypeMetatable("Player")
	L.SetGlobal("Player", mt)
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), playerMethods))
}

var playerMethods = map[string]lua.LGFunction{
	"getItemByCode": getPlayerItemByCode,
}

// checkPlayer 从 Lua stack 中获取 *Player 实例
func checkPlayer(L *lua.LState) *Player {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Player); ok {
		return v
	}
	L.ArgError(1, "Player expected")
	return nil
}

// playerGetItemByCode 是适配器函数，用于调用 Player 的 getItemByCode 方法
func getPlayerItemByCode(L *lua.LState) int {
	p := checkPlayer(L)
	return p.getItemByCode(L)
}

// 根据物品code返回玩家的装备中物品
func (p *Player) getItemByCode(L *lua.LState) int {
	code := L.ToString(2)
	bFound := false
	for _, i := range p.Inventory {
		if i.Item.GetItemName() == code {
			itemUD := L.NewUserData()
			itemUD.Value = i
			L.SetMetatable(itemUD, L.GetTypeMetatable("Item"))
			L.Push(itemUD)
			bFound = true
			break
		}
	}
	if !bFound {
		L.Push(lua.LNil)
	}

	return 1
}