package engine

import (
	"fmt"
)

type NormalItem struct {
	Owner    string
	Name     string
	Desc     string
	Detail   string
	Quantity int
	Actions  []Action
}

func (i *NormalItem) GetItemName() string {
	return i.Name
}

func (i *NormalItem) GetItemDesc() string {
	return i.Desc
}

func (i *NormalItem) GetItemDetail() string {
	return i.Detail
}

func (i *NormalItem) GetItemQuantity() int {
	return i.Quantity
}

func (i *NormalItem) SetItemQuantity(v int) int {
	i.Quantity = v
	return i.Quantity
}

// 查看此物品是否实现了某个动作
func (i *NormalItem) ContainAction(action string) []Action {
	fmt.Println("check for action:", action)
	return nil
}

func LoadItem() *NormalItem {
	return nil
}
