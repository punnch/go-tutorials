package sentence

import (
	"fmt"
	"testing"
)

// errorString needed to get code consice and readable
func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = %q, want %q", list, got, want)
}

func TestNoElements(t *testing.T) {
	list := []string{}
	want := ""
	got := JoinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestOneElement(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "mandarine"}
	want := "apple, orange, and mandarine"
	got := JoinWithCommas(list)

	if got != want {
		t.Error(errorString(list, got, want))
	}
}
