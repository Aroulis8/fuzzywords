package fuzzywords

import (
	"testing"
)

func TestReverse(t *testing.T) {
	got := Reverse("Hello")
	want := "olleH"

	if got != want {
		t.Errorf("Wanted %q, got %q",want,got)
	}
}