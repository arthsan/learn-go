package main

import "fmt"

func main() {
	letter := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15000.0,
			"Gustav":         8099.0,
		},
		"J": {
			"Josh": 12000.0,
		},
	}

	fmt.Println(letter)

	delete(letter, "J")

	for lett, emp := range letter {
		fmt.Println(lett, emp)
	}
}
