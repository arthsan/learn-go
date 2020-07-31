package main

import "fmt"

func avr(number ...float64) float64 {
	total := 0.0
	for _, num := range number {
		total += num
	}
	return total / float64(len(number))
}

func main() {
	fmt.Printf("Average: %.2f", avr(4.2, 5.6, 8.6, 9.6))
}
