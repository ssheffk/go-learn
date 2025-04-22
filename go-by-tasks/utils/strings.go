package utils

import (
	"fmt"
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RotateLeftString(s string, n int) string {
	runes := []rune(s)
	n = n % len(runes)

	return string(runes[n:]) + (string(runes[:n]))
}

func IsPalindrome(s string) bool {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func CountVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, ch := range s {
		if strings.ContainsRune(vowels, ch) {
			count++
		}
	}
	return count
}

func DemoSplitJoin() {
	parts := strings.Split("a,b,c", ",")
	fmt.Println(parts)

	joined := strings.Join(parts, "-")
	fmt.Println(joined)
}

func DemoStrings() {
	fmt.Println("### Strings")
	fmt.Println("Reverse:", Reverse("Learn GoLang"))
	fmt.Println("RotateLeft:", RotateLeftString("GoLang", 3))
	fmt.Println("IsPalindrome:", IsPalindrome("GoLang"))
	fmt.Println("CountVowels:", CountVowels("GoLang"))

	fmt.Println("DemoSplitJoin:")
	DemoSplitJoin()
}
