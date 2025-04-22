package utils

import (
	"fmt"
	"strings"
)

// Zero value of a map is nil  muse use make() to create before use

func WordFrequency(s string) map[string]int {
	freq := make(map[string]int)

	words := strings.Fields(s)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

func InvertMap(m map[string]string) map[string]string {
	inverted := make(map[string]string)

	for k, v := range m {
		inverted[v] = k
	}

	return inverted
}

func GroupByLength(words []string) map[int][]string {
	groups := make(map[int][]string)

	for _, word := range words {
		l := len(word)
		groups[l] = append(groups[l], word)
	}
	return groups
}

func MostFrequentChar(s string) rune {
	freq := make(map[rune]int)
	maxChar := rune(0)
	maxCount := 0

	for _, r := range s {
		freq[r]++
		if freq[r] > maxCount {
			maxCount = freq[r]
			maxChar = r
		}
	}
	return maxChar
}

func AreAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freq := make(map[rune]int)
	for _, r := range s1 {
		freq[r]++
	}

	for _, r := range s2 {
		freq[r]--
		if freq[r] < 0 {
			return false
		}
	}
	return true
}

func DemoMaps() {
	fmt.Println("WordFrequency: ", WordFrequency("go is fun and go is fast"))
	fmt.Println("InvertMap: ", InvertMap(map[string]string{"a": "1", "b": "2"}))
	fmt.Println("GroupByLength: ", GroupByLength([]string{"go", "map", "slice", "if", "else"}))
	fmt.Printf("%c\n", MostFrequentChar("Learning by doing GoLang"))
	fmt.Println("AreAnagrams: ", AreAnagrams("listen", "silent"))
}
