package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum = ", a+b)
	fmt.Println("Sub = ", a-b)
	fmt.Println("Div = ", a/b)
	fmt.Println("Mul = ", a*b)
	fmt.Println("Mod = ", a%b)

	fmt.Println("AND => ", a&b)
	fmt.Println("OR => ", a|b)

	c := 3.0
	d := 2.0

	fmt.Println("Bigger =>", math.Max(float64(c), float64(d)))

}
