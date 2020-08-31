package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func prime(n int, c chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for prim := start; ; prim++ {
			if isPrime(prim) {
				c <- prim
				start = prim + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}

func main() {
	c := make(chan int, 30)
	go prime(cap(c), c)
	for prime := range c {
		fmt.Printf("%d ", prime)
	}
}
