package main

import (
	"fmt"
)

type Print interface {
	toString() string
}

type Product struct {
	name  string
	price float64
}

type Person struct {
	name    string
	surname string
}

func (p Person) toString() string {
	return p.name + " " + p.surname
}

func (p Product) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.name, p.price)
}

func printer(x Print) {
	fmt.Println(x.toString())
}

func main() {
	var thing Print = Person{"Robert", "Richard"}
	fmt.Println(thing.toString())
	printer(thing)

	thing = Product{"Mac", 2500}
	fmt.Println(thing.toString())
	printer(thing)
}
