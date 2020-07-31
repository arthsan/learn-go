package main

import "fmt"

func value(number int) int {
	defer fmt.Println("end")
	if number > 5000 {
		fmt.Println("High Value")
		return number
	} else {
		fmt.Println("Low Value")
		return number
	}
}

func main() {
	fmt.Println(value(6000))
	fmt.Println(value(2000))
}
