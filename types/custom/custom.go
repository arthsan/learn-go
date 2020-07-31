package main

import "fmt"

type grade float64

func (g grade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func gradeConcept(g grade) string {
	if g.between(9.0, 10.0) {
		return "A"
	} else if g.between(7.0, 8.99) {
		return "B"
	} else if g.between(5.0, 7.99) {
		return "C"
	} else if g.between(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(gradeConcept(3.7))
	fmt.Println(gradeConcept(7.7))
	fmt.Println(gradeConcept(9.7))
}
