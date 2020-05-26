package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	grade := 6.9
	final := int(grade)
	fmt.Println("Final grade: ", final)

	fmt.Println("Convert " + string(100))
	fmt.Println("Convert" + strconv.Itoa(123))
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("true")
	}
}
