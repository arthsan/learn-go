package math

import "testing"

const err = "Expected %v, Result %v"

func TestAverage(t *testing.T) {
	t.Parallel()
	expected := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != expected {
		t.Errorf(err, expected, value)
	}
}
