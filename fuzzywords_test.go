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

	answer,err := Fuzz("Hello World And World World")

	if err != nil {
		t.Errorf("error is not nill")
	}
	t.Logf("Answer: %q",answer)
}

func TestRandomCharacters(t *testing.T) {
	answer,err := RandomCharacters(10)

	if len(answer) != 10 {
		t.Errorf("the length of answer is not 10, it is %q", len(answer))
	} else if err != nil {
		t.Errorf("'err' is not nil")
	}

	t.Logf("Answer: %q",answer)
}

