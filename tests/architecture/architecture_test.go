package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Doesn't work on amd64")
	}
	t.Fail()
}
