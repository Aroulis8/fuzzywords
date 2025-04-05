package fuzzywords
import (
	"strings"
	"math/rand"
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