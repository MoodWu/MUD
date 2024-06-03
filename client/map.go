package main

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

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
	if s == nil {
		return false
	}
	_, ok := s.Path[direction]
	return ok
}

func (m *Map) GetSceneByXY(x, y int) *Scene {
	for _, s := range m.Scenes {
		if s.X == x && s.Y == y {
			return s
		}
	}
	return nil
}

func (m *Map) GetSceneByCode(code string) *Scene {
	for _, s := range m.Scenes {
		if s.Code == code {
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

func (s *Scene) GetDesc() string {
	ret := s.Desc + "\n"
	ret += "这里明显可见的方向有:"
	for k, _ := range s.Path {
		ret += k + " "
	}

	goodList := ""
	bFlag := false

	for _, goods := range s.Items {
		if goods.Quantity > 0 {
			bFlag = true
			goodList += goods.Item.GetItemName() + ",数量：" + strconv.Itoa(goods.Quantity) + "\n"
		}
	}
	if bFlag {
		ret += "\n这里的物品有:\n" + goodList
	}

	return ret
}
