package main

import "fmt"

type Product struct {
	name     string
	price    float64
	discount float64
}

func (p Product) priceWithDiscount() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 Product
	product1 = Product{
		name:     "Pencil",
		price:    1.3,
		discount: 0.20,
	}
	fmt.Println(product1, product1.priceWithDiscount())
	product2 := Product{name: "Notebook", price: 2000}
	fmt.Println(product2.priceWithDiscount())
}
