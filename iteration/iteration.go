package iteration

import "strings"

func Repeat(character string, amount int) string {
	var repeated string
	for i := 0; i < amount; i++ {
		repeated += character
	}

	return repeated
}

func checkIndex(word, character string) int {
	index := strings.Index(word, character)

	return index
}
