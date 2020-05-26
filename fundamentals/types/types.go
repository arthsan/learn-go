package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer = ", reflect.TypeOf(32000))

	var b byte = 255
	fmt.Println("Byte =", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("Type =", i1)

	var i2 rune = 'a'
	fmt.Println(i2)

	var x float32 = 49.99
	fmt.Println(x)
	fmt.Println("X = ", reflect.TypeOf(x))

	bo := true
	fmt.Println("Bo = ", reflect.TypeOf(bo))

	s1 := "Hello World"
	fmt.Println("String" + s1)
	fmt.Println(len(s1))

	s2 := `Hello 
	World`
	fmt.Println(s2)
}
