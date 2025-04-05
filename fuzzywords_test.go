package fuzzywords

import (
	"testing"
)

func TestReverse(t *testing.T) {
	got,err := Reverse("Hello")
	want := "olleH"

	if got != want {
		t.Errorf("Wanted %q, got %q",want,got)
	} else if err != nil {
		t.Errorf("error is nill")
	}
}

func TestFuzz(t *testing.T) {

	_,err := Fuzz("Hello")

	if err != nil {
		t.Errorf("error is nill")
	}
}