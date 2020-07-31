package main

import "fmt"

type Item struct {
	productID int
	qtd       int
	price     float64
}

type Order struct {
	userID int
	items  []Item
}

func (o Order) Value() float64 {
	total := 0.0
	for _, Item := range o.items {
		total += Item.price * float64(Item.qtd)
	}
	return total
}

func main() {
	order := Order{
		userID: 1,
		items: []Item{
			Item{1, 2, 12.10},
			Item{2, 3, 30.2},
			Item{11, 100, 3.1},
		},
	}
	fmt.Printf("Total value is %.2f", order.Value())
}
