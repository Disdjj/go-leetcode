package problems

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	l1, l2 := len(str1), len(str2)
	maxL := gcd(l1, l2)
	return str1[:maxL]
}
