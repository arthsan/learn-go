package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := random(); i > 5 {
		fmt.Println("Win")
	} else {
		fmt.Println("Lose")
	}
}
