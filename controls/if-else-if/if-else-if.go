package main

import "fmt"

func grades(n float64) string {
	if n >= 9 && n <= 10 {
		return "A"
	} else if n < 9 && n >= 7 {
		return "B"
	} else if n < 7 && n >= 5 {
		return "C"
	} else if n < 5 && n >= 3 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(grades(3))
}
