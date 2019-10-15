package test1

import (
	"sort"
	"strings"
)

func Test1(word string) map[string]string{
	var vowels []string
	var consonants []string
	var data []string

	res := make(map[string]string)

	words := strings.Split(word, "")

	for _, char := range words {
		switch char {
			case "a", "i", "u", "e", "o", "A", "I", "U", "E", "O":
				vowels = append(vowels, char)
			default:
				consonants = append(consonants, char)
		}
	}

	sort.Strings(vowels)
	sort.Strings(consonants)

	data = append(vowels, consonants...)

	res["vowels"] = strings.Join(vowels, "")
	res["consonants"] = strings.Join(consonants, "")
	res["result"] = strings.Join(data, "")
	

	return res
}