package model

var (
	DefaultItems []*ShoppingItem = []*ShoppingItem{
		NewShoppingItem("Milk"),
		NewShoppingItem("Eggs"),
		NewShoppingItem("Bread"),
		NewShoppingItem("Cheese"),
	}
)

type ShoppingItem struct {
	name string
}

func NewShoppingItem(name string) *ShoppingItem {
	return &ShoppingItem{name: name}
}

func (si *ShoppingItem) Name() string {
	return si.name
}
