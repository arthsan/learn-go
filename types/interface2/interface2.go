package main

import "fmt"

type Sport interface {
	turnTurboON()
}

type Ferrari struct {
	model   string
	turboOn bool
	speed   int
}

func (f *Ferrari) turnTurboOn() {
	f.turboOn = true
	f.speed += 50
}

func main() {
	car1 := Ferrari{"F40", false, 0}
	car1.turnTurboOn()

	var car2 = &Ferrari{"F50", true, 100}
	car2.turnTurboOn()

	fmt.Println(car1, car2)
}
