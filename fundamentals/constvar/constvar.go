package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 5.5

	area := PI * math.Pow(raio, 2)

	fmt.Println("Area =", area)

	const (
		a = 2
		b = 3
	)

	var e, f bool = true, false

	fmt.Println(e, f)

	c, d := false, 3

	fmt.Println("Number: ", d, c)

	fmt.Println("Variables: ", a, b)
}
