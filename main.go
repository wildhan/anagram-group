package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	strings := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	result := GroupAnagrams(strings)
	fmt.Println(result)
}

func GroupAnagrams(input []string) [][]string {
	anagramMap := make(map[string][]string)
	var result [][]string

	for _, str := range input {
		key := CreateKey(str)
		anagramMap[key] = append(anagramMap[key], str)
	}

	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func CreateKey(s string) string {
	type CountMap map[rune]int
	countChar := make(CountMap)
	for _, char := range s {
		countChar[char]++
	}

	var groupChar []string
	for char, count := range countChar {
		groupChar = append(groupChar, fmt.Sprintf("%c%d", char, count))
	}
	sort.Strings(groupChar)
	return strings.Join(groupChar, "")
}
