package fuzzywords

import (
	"math/rand"
	"strings"
	"errors"
)


var randomCharactersSlice = []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

//Returns a random sentence made with the words from str
func Fuzz(str string) (string,error){
	if strings.TrimSpace(str) == "" || str == "" {
		return "nil",errors.New("empty string")
	}
	strSlice := strings.Split(str, " ")
	strSliceLength := len(strSlice)
	result := ""
	for i := 0; i < strSliceLength; i++ {
		result += strSlice[rand.Intn(strSliceLength)] + " "
	}

	return result,nil
}

//Returns str, but desrever(reversed)               
func Reverse(str string) (string, error) {
	if strings.TrimSpace(str) == "" || str == "" {
		return "nil",errors.New("empty string")
	}

	strLength := len(str)
	result := ""
	for i := strLength - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result,nil
}

//Returns a string of "n" random characters.
//If you want the string to be made with specific characters, change the "randomCharactersSlice"
func RandomCharacters(n int) (string, error){
	if n <= 0 {
		return "nil",errors.New("'n' is 0 or less")
	}
	result := ""
	for i := 0;i < n; i++ {
		result += randomCharactersSlice[rand.Intn(len(randomCharactersSlice))]
	}
	return result,nil
}

func AddPrefix(str string, prefix string) string{
	strSlice := strings.Split(str," ")
	result := ""
	for i := 0; i < len(strSlice); i++ {
		result += prefix + strSlice[i] + " "
	}

	return strings.TrimSpace(result)
}