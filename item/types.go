package item

type Item interface {
	GetItemName() string
	GetItemQuantity() int
	GetItemDesc() string
}

type Thing struct {
	Thing    Item
	Quantity int
}
