package Valid_Palindrome

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	r := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}
	s = strings.Map(r, s)

	lenght := len(s)
	for i := 0; i < lenght/2; i++ {
		if s[i] != s[lenght-1-i] {
			return false
		}
	}
	return true
}
