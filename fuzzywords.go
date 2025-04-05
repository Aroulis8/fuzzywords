package fuzzywords

import (
	"math/rand"
	"strings"
	"errors"
)

//Returns a random sentence made with the words from str
func Fuzz(str string) string{
	strSlice := strings.Split(str, " ")
	strSliceLength := len(strSlice)
	result := ""
	for i := 0; i < strSliceLength; i++ {
		result += strSlice[rand.Intn(strSliceLength)] + " "
	}

	return result
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