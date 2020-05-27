package main

import "fmt"

func print(grade float64) {
	if grade >= 7 {
		fmt.Println("Success ", grade)
	} else {
		fmt.Println("Failed ", grade)
	}
}

func main() {
	print(3)
}
