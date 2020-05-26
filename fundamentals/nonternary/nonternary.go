package main

import "fmt"

func result(grade float64) string {
	if grade >= 6 {
		return "Success"
	}
	return "Failed"
}

func main() {
	fmt.Println(result(6.2))
}
