package engine

type Item interface {
	GetItemName() string
	GetItemQuantity() int
	GetItemDesc() string
	GetItemDetail() string
}

type Goods struct {
	Item     Item
	Quantity int
}
