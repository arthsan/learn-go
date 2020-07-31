package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	s3 := a2[:2]
	s4 := s3[:1]

	fmt.Println(a2, s2, s3, s4)

}

