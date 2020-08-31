package main

import (
	"fmt"

	"github.com/coders/html"
)

func send(orign <-chan string, destiny chan string) {
	for {
		destiny <- <-orign
	}
}

func merge(ch1, ch2 <-chan string) <-chan string {
	c := make(chan string)
	go send(ch1, c)
	go send(ch2, c)
	return c
}

func main() {
	c := merge(
		html.Tittle("https://www.tradersclub.com.br", "https://www.amazon.com.br"),
		html.Tittle("https://www.google.com.br", "https://www.youtube.com.br"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
