package main

import (
	"fmt"
	"time"

	"github.com/coders/html"
)

func faster(url1, url2, url3 string) string {
	c1 := html.Tittle(url1)
	c2 := html.Tittle(url2)
	c3 := html.Tittle(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(2000 * time.Millisecond):
		return "No response"
	}
}

func main() {
	champ := faster(
		"https://www.empiricus.com.br/",
		"https://www.xpi.com.br/",
		"https://www.tradersclub.com.br",
	)
	fmt.Println(champ)
}
