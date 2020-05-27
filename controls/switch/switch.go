package main

import (
	"fmt"
	"time"
)

func grade(n float64) string {
	var rank = int(n)
	switch rank {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5, 4:
		return "C"
	default:
		return "D"
	}
}

func main() {
	fmt.Println(grade(3))
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 18:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Night")
	}
	fmt.Println(theType(2.1))
}
