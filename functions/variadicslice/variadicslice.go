package main

import "fmt"

func printSuccess(success ...string) {
	fmt.Println("Success List")
	for i, suc := range success {
		fmt.Printf("%d) %s\n", i+1, suc)
	}
}

func main() {
	suc := []string{"Jorge", "Pedro", "Javier"}
	printSuccess(suc...)
}
