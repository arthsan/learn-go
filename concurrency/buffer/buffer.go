package main

import (
	"fmt"
	"time"
)

func rotin(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Start")
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3)
	go rotin(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
