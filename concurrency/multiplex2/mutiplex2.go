package main

import (
	"fmt"
	"time"
)

func speak(person string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s speaking> %d", person, i)
		}
	}()
	return c
}

func merge(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				c <- s
			case s := <-ch2:
				c <- s
			}
		}
	}()
	return c
}

func main() {
	c := merge(speak("John"), speak("Mary"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}
