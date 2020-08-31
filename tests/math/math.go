package math

import (
	"fmt"
	"strconv"
)

//average
func Average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	avg := total / float64(len(numbers))
	avgred, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
	return avgred
}
