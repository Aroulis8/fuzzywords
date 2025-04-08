package fuzzywords

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	got := Reverse("Hello")
	want := "olleH"

	if got != want {
		t.Errorf("Wanted %q, got %q",want,got)
	}

	t.Logf("Got: %q",got)
}

func TestRandomCharacters(t *testing.T) {
	doesContainA := true
	doesContainB := true
	doesContainC := true
	RandomCharactersSlice = []string {"a","b","c"}

	got := RandomCharacters(10)

	if len(got) != 10 {
		t.Errorf("the length of answer is not 10, it is %q", len(got))
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
	var want StringData
	want.spaces = 1
	want.characters = 10
	want.words = 2

	if got != want {
		t.Errorf("Got %q, wanted %q",got, want)
	}

	t.Logf("Got: %v\n Want: %v",got,want)
}

func TestFuzzySentence(t *testing.T) {
	FuzzySentenceObjects = []string {"Aris"}
	FuzzySentenceVerbs = []string {"killed"}
	FuzzySentencePlaces = []string {"at home"}
	got := CreateFuzzySentence("Viss")
	want := "Viss killed Aris at home"

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}

	t.Logf("Got %q", got)
}