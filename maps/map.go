package maps

import (
	"log"
	"mud/common"
	"mud/item"
	"os"
	"strings"
	"unicode/utf8"

	"gopkg.in/yaml.v3"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", 0)
	logger.SetFlags(0)
}

// 初始化地图
func InitMap() *Map {
	// m := Map{"新手村", "这是一个不大的村子,看起来安静祥和,你可以在这里学习如何探索世界.", 3, 3, 10, nil, nil}
	// m.Scenes = make([]*Scene, 0)
	// //开始实例化地图下的场景
	// s11 := Scene{&m, 1, 1, "村子", "村子不大,村子北面有一座山,西边是一边树林,东边是一片田地,南面有一条小河流过,看起来很安静,村里有小孩在嬉戏,有老人在村头的大树下下棋.", nil, nil}
	// s01 := Scene{&m, 1, 1, "树林", "一片稀疏的果林,有桃树和李树", nil, nil}
	// s10 := Scene{&m, 1, 1, "小河", "一条10米宽的小河,水流缓慢,水很浅,可以看到河底的鹅卵石和成群的小鱼", nil, nil}
	// s12 := Scene{&m, 1, 1, "后山", "一座不太高的小山包,但是周围地势平坦,在山顶应该可以看到很远.", nil, nil}
	// s21 := Scene{&m, 1, 1, "稻田", "一片绿油油的稻田,快到收获的季节了,地里有农民在除草.", nil, nil}

	// s11.Paths = make([]*Path, 0, 4)
	// s11.Paths = append(s11.Paths, &Path{"south", &s10}, &Path{"north", &s12}, &Path{"west", &s01}, &Path{"east", &s21})

	// s01.Paths = make([]*Path, 0, 1)
	// s01.Paths = append(s01.Paths, &Path{"east", &s11})

	// s21.Paths = make([]*Path, 0, 1)
	// s21.Paths = append(s21.Paths, &Path{"west", &s11})

	// s12.Paths = make([]*Path, 0, 1)
	// s12.Paths = append(s12.Paths, &Path{"south", &s11})

	// s10.Paths = make([]*Path, 0, 1)
	// s10.Paths = append(s10.Paths, &Path{"north", &s11})

	data, err := os.ReadFile("data/newbie.yaml")
	if err != nil {
		log.Println("read file error:", err)
		return nil
	}

	m := LoadMap(data)
	return m
}

// 从yaml中加载map
func LoadMap(data []byte) *Map {
	//先反序列化为 YMAL_Map
	m := &common.YAML_Map{}
	if err := yaml.Unmarshal(data, m); err != nil {
		log.Println("yaml unmarshal err.", err)
		return nil
	}

	//开始构建Map
	m0 := Map{Name: strings.Trim(m.Name, " "), Desc: m.Desc, Code: m.Code}
	m0.Long, m0.Wide = common.Convert2XY(m.Scale)
	m0.RefreshTime = 10
	m0.Scenes = make([]*Scene, 0)

	// 加载所有的场景
	for _, v := range m.Scenes {
		// todo:现在都默认一种类型，以后需要根据type动态创建
		s := &Scene{Name: v.Name, Desc: v.Desc, Code: v.Code}
		s.X, s.Y = common.Convert2XY(v.Position)
		s.Map = &m0
		s.Items = make([]*item.Thing, 0)
		s.Paths = make([]*Path, 0)

		//加载场景上的物品
		for _, i := range v.Items {
			// todo:现在都默认一种类型，以后需要根据type动态创建
			thing := item.Thing{}

			item := &item.NormalItem{Name: i.Item[0].Name, Desc: i.Item[0].Desc, Quantity: i.Item[0].Quantity}
			item.Actions = make([]common.Action, 0)
			for _, a := range i.Item[0].Actions {
				action := &common.NoramlAction{Name: a.Name, Desc: a.Desc, Effect: a.Effect, EffectDesc: a.EffectDesc, Condition: a.Condition}
				item.Actions = append(item.Actions, action)
			}
			thing.Thing = item
			thing.Quantity = i.Quantity

			s.Items = append(s.Items, &thing)
		}

		//加载路径
		for k, v := range v.Direction {
			path := Path{Direction: k, SceneCode: v}
			s.Paths = append(s.Paths, &path)
		}

		//将移动也作为一种特殊的Item附加在场景上
		m0.Scenes = append(m0.Scenes, s)
	}

	// 地图间的连接在另外的编排文件中
	m0.Connections = make([]*Connection, 0)

	// printMap(&m0)
	return &m0
}

/*
	func printMap(m *Map) {
		logger.Printf("name:%s\ncode:%s\ndesc:%s\nwidth:%d\nlength:%d\nrefreshtime:%d\n", m.Name, m.Code, m.Desc, m.Width, m.Length, m.RefreshTime)
		for _, v := range m.Scenes {
			logger.Printf("\tx:%d\n\ty:%d\n\tname:%s\n\tcode:%s\n\tdesc:%s\n", v.X, v.Y, v.Name, v.Code, v.Desc)
			for _, i := range v.Items {
				logger.Printf("\tItems:\n")
				logger.Printf("\t\tqty:%d\n", i.Quantity)
				logger.Printf("\t\titem:\n")
				logger.Printf("\t\t\t%+v\n", i.Thing)
			}

			for _, p := range v.Paths {
				logger.Printf("\t%s:%s\n", p.Direction, p.SceneCode)
			}
		}

		for _, c := range m.Connections {
			logger.Printf("%s:%s\n", c.Direction, c.MapCode)
		}

}
*/
func (m *Map) ShowMap() string {
	var sb strings.Builder

	sb.WriteString(m.Name + "\n\n")

	colPadding := 1

	max := make([]int, m.Long)
	s := make([][]int, m.Wide)
	n := make([][]string, m.Wide)
	for i := 0; i < m.Wide; i++ {
		s[i] = make([]int, m.Long)
		n[i] = make([]string, m.Long)
		for j := 0; j < m.Long; j++ {
			n[i][j] = m.GetSceneName(i, j)
			s[i][j] = utf8.RuneCountInString(n[i][j])

			//记录每列最大宽度
			if max[j] < s[i][j] {
				max[j] = s[i][j]
			}
		}
	}

	//开始输出，每一行被分为三行，向北，向南的连接行，中间的数据行
	nConnector := ""
	sConnector := ""
	content := ""
	for i := m.Wide - 1; i >= 0; i-- {
		nConnector = ""
		sConnector = ""
		content = ""
		for j := 0; j < m.Long; j++ {
			nConnector += m.getConnectorString("north", max[j]+2*colPadding, j, i)
			content += m.getPaddedString(n[j][i], max[j]+2*colPadding, j, i)
			sConnector += m.getConnectorString("south", max[j]+2*colPadding, j, i)
			// sb.WriteString(content)
		}
		if i != m.Wide-1 {
			sb.WriteString(nConnector + "\n")
		}
		sb.WriteString(content + "\n")
		if i != 0 {
			sb.WriteString(sConnector + "\n")
		}
	}

	return sb.String()
}

func (m *Map) getConnectorString(direction string, width, x, y int) string {
	if m.HasPath(direction, x, y) {
		return strings.Repeat("　", width/2) + "｜" + strings.Repeat("　", width/2-1)
	}

	return strings.Repeat("　", width)
}

func (m *Map) getPaddedString(s string, width int, x, y int) string {
	ret := s
	if utf8.RuneCountInString(s) >= width-1 {
		return ret
	}
	padding := (width - utf8.RuneCountInString(s)) / 2

	if m.HasPath("west", x, y) {
		ret = strings.Repeat("－", padding) + s
	} else {
		ret = strings.Repeat("　", padding) + s
	}
	padding = width - utf8.RuneCountInString(s) - padding

	if m.HasPath("east", x, y) {
		ret += strings.Repeat("－", padding)
	} else {
		ret += strings.Repeat("　", padding)
	}

	return ret
}

// 查询x,y坐标的场景是否有direction上的路径
func (m *Map) HasPath(direction string, x, y int) bool {
	s := m.GetSceneByXY(x, y)
	if s != nil {
		for _, p := range s.Paths {
			if p.Direction == direction {
				return true
			}
		}
	}
	return false
}

func (m *Map) GetSceneByXY(x, y int) *Scene {
	for _, s := range m.Scenes {
		if s.X == x && s.Y == y {
			return s
		}
	}
	return nil
}
func (m *Map) GetSceneName(x, y int) string {
	s := m.GetSceneByXY(x, y)
	if s == nil {
		return ""
	}
	return s.Name
}
