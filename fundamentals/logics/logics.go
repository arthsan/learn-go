package main

import "fmt"

func buy(work1, work2 bool) (bool, bool, bool) {
	buy50 := work1 && work2
	buy32 := work1 != work2
	buyIce := work1 || work2

	return buy50, buy32, buyIce
}

func main() {
	tv50, tv32, ice := buy(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Ice: %t, Health: %t", tv50, tv32, ice, !ice)
}
