package problems

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	word1Len := len(word1)
	word2Len := len(word2)
	if word1Len == 0 {
		return word2
	}

	if word2Len == 0 {
		return word1
	}
	builder := strings.Builder{}
	for i := 0; i < word1Len && i < word2Len; i++ {
		builder.WriteString(string(word1[i]))
		builder.WriteString(string(word2[i]))
	}
	if word1Len > word2Len {
		builder.WriteString(word1[word2Len:])
	} else {
		builder.WriteString(word2[word1Len:])
	}
	return builder.String()
}
