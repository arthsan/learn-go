package main

import "fmt"

func main() {
	salary := map[string]float64{
		"Josh":   10000.0,
		"Peter":  8000.0,
		"Nadrea": 9000.0,
	}

	salary["Peter"] = 13000.0
	delete(salary, "Philipe")

	for name, wage := range salary {
		fmt.Println(name, wage)
	}
}
