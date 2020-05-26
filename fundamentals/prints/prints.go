package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print("Line. ")

	fmt.Println("New")
	fmt.Println("Line")

	x := 3.141516

	xs := fmt.Sprint(x)
	fmt.Println("X = " + xs)
	fmt.Println("X = ", x)

	fmt.Printf("X = %.2f.", x)

	a := 1
	b := false
	c := "opa"

	fmt.Printf("\n%d %t %s", a, b, c)
}
