package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (part: %s) - index: waiting (%d) <> found (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Cod3r show", "Cod3r", 0},
		{"", "", 0},
		{"Hello", "World", -1},
		{"Go", "Learning", -1},
	}
	for _, test := range tests {
		t.Logf("Mass: %v", test)
		atual := strings.Index(test.text, test.part)

		if atual != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, atual)
		}
	}
}
