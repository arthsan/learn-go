package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("String: ", "Apple" == "Apple")
	fmt.Println("!=", 3 != 2)
	fmt.Println("<", 3 < 2)
	fmt.Println(">", 3 > 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Date: ", d1 == d2)
}
