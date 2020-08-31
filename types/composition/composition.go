package main

import "fmt"

type Sport interface {
	turnOn()
}

type Luxury interface {
	park()
}

type SportLux interface {
	Sport
	Luxury
}

type bmw7 struct{}

func (b bmw7) turnOn() {
	fmt.Println("Turbo")
}

func (b bmw7) park() {
	fmt.Println("Parking")
}

func main() {
	var b SportLux = &bmw7{}
	b.park()
	b.turnOn()
}
