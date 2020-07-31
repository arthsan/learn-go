package main

import "fmt"

func chance(p1, p2 int) (second, first int) {
	second = p2
	first = p1
	return
}

func main() {
	x, y := chance(2, 4)
	fmt.Println(x, y)

	second, first := chance(6, 15)
	fmt.Println(second, first)
}
