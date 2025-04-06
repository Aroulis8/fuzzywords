package fuzzywords

import (
	"strings"
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

	t.Logf("Got: %q",got)
}

func TestFuzz(t *testing.T) {

	got,err := Fuzz("Hello World And World World")

	if err != nil {
		t.Errorf("error is not nill")
	}
	t.Logf("Got: %q",got)
}

func TestRandomCharacters(t *testing.T) {
	doesContainA := true
	doesContainB := true
	doesContainC := true
	randomCharactersSlice = []string {"a","b","c"}

	got,err := RandomCharacters(10)

	if len(got) != 10 {
		t.Errorf("the length of answer is not 10, it is %q", len(got))
	} else if err != nil {
		t.Errorf("'err' is not nil")
	}

	if (!strings.Contains(got,"a")) {
		doesContainA = false
	}
	if (!strings.Contains(got, "b")) {
		doesContainB = false
	}
	if (!strings.Contains(got, "c")) {
		doesContainC = false
	}

	if (doesContainA == false && doesContainB == false && doesContainC == false) {
		t.Error("answer does not contain any valid characters")
	}

	t.Logf("Got: %q",got)
}

func TestAddPrefix(t *testing.T) {
	got := AddPrefix("Hello there","a")
	want := "aHello athere"
	if got != want {
		t.Errorf("Wanted %q, got %q",want,got)
	}

	t.Logf("Got: %q", got)
}

func TestAddSuffix(t *testing.T) {
	got := AddSuffix("Hello there", "I")
	want := "HelloI thereI"

	if got != want {
		t.Errorf("Got: %q, wanted: %q",got,want)
	}

	t.Logf("Got: %q", got)
}

func TestGetStringData(t *testing.T) {
	got := GetStringData("Hello World")
	want := map[string]int {
		"words": 2,
		"spaces": 1,
		"characters":10,
	}

	for k,v := range want {
		if (got[k] == v) {
			continue
		} else {
			t.Errorf("Maps don't match \nGot:%v\nWant: %v",got,want)
		}
	}

	t.Logf("Got: %v\n Want: %v",got,want)
}

