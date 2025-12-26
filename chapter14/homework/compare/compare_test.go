package compare

import (
	"fmt"
	"testing"
)

func TestFirstLarger(t *testing.T) {
	want := 3
	got := Larger(3, 1)

	if got != want {
		t.Error(errorString(3, 1, got, want))
	}
}

func TestSecondLarger(t *testing.T) {
	want := 7
	got := Larger(5, 7)

	if got != want {
		t.Error(errorString(5, 7, got, want))
	}
}

func errorString(a int, b int, got int, want int) string {
	return fmt.Sprintf("Larger(%d, %d) = %d, want %d", a, b, got, want)
}
