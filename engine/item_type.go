package engine

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
