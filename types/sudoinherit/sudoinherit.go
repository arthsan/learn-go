package main

import "fmt"

type Car struct {
	name  string
	speed int
}

type Ferrari struct {
	Car
	turboOn bool
}

func main() {
	f := Ferrari{}
	f.name = "F50"
	f.speed = 270
	f.turboOn = true

	fmt.Printf("Your %s speed is: %v", f.name, f.speed)
}
