package item

import (
	"fmt"
	"mud/common"
)

type NormalItem struct {
	Owner    string
	Name     string
	Desc     string
	Quantity int
	Actions  []common.Action
}

func (i *NormalItem) GetItemName() string {
	return i.Name
}

func (i *NormalItem) GetItemDesc() string {
	return i.Desc
}

func (i *NormalItem) GetItemQuantity() int {
	return i.Quantity
}

// 查看此物品是否实现了某个动作
func (i *NormalItem) ContainAction(action string) []common.Action {
	fmt.Println("check for action:", action)
	return nil
}

func LoadItem() *NormalItem {
	return nil
}
