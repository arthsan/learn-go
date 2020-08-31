package main

import (
	"fmt"
	"time"
)

func rotin(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("Next")
}

func main() {
	c := make(chan int)

	go rotin(c)

	fmt.Println(<-c)
	fmt.Println("Wait")
	fmt.Println(<-c)
	fmt.Println("End")
}
