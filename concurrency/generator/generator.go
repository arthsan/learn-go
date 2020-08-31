package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func tittle(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func main() {
	t1 := tittle("https://www.google.com.br", "https://www.youtube.com", "")
	t2 := tittle("https://www.amazon.com", "https://www.tradersclub.com.br")
	fmt.Println("First:", <-t1, "|", <-t2)
	fmt.Println("Second:", <-t1, "|", <-t2)
}
