package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name    string
	surname string
}

func (p Person) getFullName() string {
	return p.name + " " + p.surname
}

func (p *Person) setFullName(FullName string) {
	parts := strings.Split(FullName, " ")
	p.name = parts[0]
	p.surname = parts[1]
}

func main() {
	p1 := Person{"Peter", "Parker"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Bruce Banner")
	fmt.Println(p1.getFullName())
}
