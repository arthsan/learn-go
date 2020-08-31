package main

import (
	"fmt"
	"time"
)

func speak(person, txt string, amount int) {
	for i := 0; i < amount; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", person, txt, i+1)
	}
}

func main() {
	go speak("May", "What are talking about?", 10)
	speak("Pete", "You!", 3)
}
