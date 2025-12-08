package stringz

import "strings"

func ReverseWordsBuilder(s string) string {
	var sb strings.Builder

	length := len(s)
	sb.Grow(length)

	start := 0

	writeword := func(end int) {
		for j := end; j >= start; j-- {
			sb.WriteByte(s[j])
		}
	}

	for i := range length {

		if s[i] == ' ' {
			writeword(i - 1)
			sb.WriteByte(' ')
			start = i + 1
		}

		if i == length-1 {
			writeword(i)
		}
	}

	return sb.String()
}
