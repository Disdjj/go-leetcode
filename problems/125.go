package problems

import "strings"

func isAlphOrNum(ch byte) bool {
	return (ch >= 'A' && ch < 'Z') || (ch >= 'a' && ch < 'z') || (ch >= '0' && ch < '9')
}
func isPalindrome(s string) bool {
	l := 0
	r := len(s) - 1
	s = strings.ToLower(s)
	for l < r {
		if !isAlphOrNum(s[l]) {
			l++
			continue
		}

		if !isAlphOrNum(s[r]) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
