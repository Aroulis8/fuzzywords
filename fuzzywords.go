package fuzzywords

import (
	"math/rand"
	"strings"
)
var RandomCharactersSlice = []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

var FuzzySentenceVerbs = []string {"played with","killed","had a bath with"}
var FuzzySentenceObjects = []string {"a parrot","a spaceship","a board game", "Mark"}
var FuzzySentencePlaces = []string {"at home","in the jungle"}

//Returns a random sentence made with the words from str
//Deprecated: This function is not useful
func Fuzz(str string) string{
	if strings.TrimSpace(str) == "" || str == "" {
		return "nil"
	}
	strSlice := strings.Split(str, " ")
	strSliceLength := len(strSlice)
	result := ""
	for i := 0; i < strSliceLength; i++ {
		result += strSlice[rand.Intn(strSliceLength)] + " "
	}

	return result
}

//Returns str, but desrever(reversed)               
func Reverse(str string) string {
	if strings.TrimSpace(str) == "" || str == "" {
		return "nil"
	}

	strLength := len(str)
	result := ""
	for i := strLength - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

//Returns a string of "n" random characters.
//If you want the string to be made with specific characters, change the "randomCharactersSlice"
func RandomCharacters(n int) string{
	if n <= 0 {
		return "<EMPTY>"
	}
	result := ""
	for i := 0;i < n; i++ {
		result += RandomCharactersSlice[rand.Intn(len(RandomCharactersSlice))]
	}
	return result
}

//Adds a prefix to every word in the string
func AddPrefix(str string, prefix string) string{
	strSlice := strings.Split(str," ")
	result := ""
	for i := 0; i < len(strSlice); i++ {
		result += prefix + strSlice[i] + " "
	}

	return strings.TrimSpace(result)
}

//Adds a suffix to every word in the string
func AddSuffix(str string, suffix string) string{
	strSlice := strings.Split(str," ")
	result := ""
	for i := 0; i < len(strSlice); i++ {
		result +=  strSlice[i] + suffix +" "
	}

	return strings.TrimSpace(result)
}

//Returns a map structured like this:
//"words": 4
//"spaces": 3
//"characters": 21
//Note that the spaces are subtracted from the characters
func GetStringData(str string)	map[string]int {
	words := len(strings.Split(str," "))
	spaces := words - 1
	characters := len(str) - spaces

	result := map[string]int {
		"words": words,
		"spaces": spaces,
		"characters": characters,
	}

	return result
}

func CreateFuzzySentence(name string) string {
	fuzzyVerb := FuzzySentenceVerbs[rand.Intn(len(FuzzySentenceVerbs))]
	fuzzyObject := FuzzySentenceObjects[rand.Intn(len(FuzzySentenceObjects))]
	fuzzyPlace := FuzzySentencePlaces[rand.Intn(len(FuzzySentencePlaces))]
	return name + " " + fuzzyVerb + " " + fuzzyObject + " " + fuzzyPlace
}