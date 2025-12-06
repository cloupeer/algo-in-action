package reverseword

import "strings"

func ReverseWordsSplit(s string) string {
	splited := strings.Split(s, " ")
	newwords := make([]string, 0, len(splited))

	for _, str := range splited {
		newwords = append(newwords, reverseword(str))
	}

	return strings.Join(newwords, " ")
}

func reverseword(s string) string {
	newword := make([]byte, 0, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		newword = append(newword, s[i])
	}

	return string(newword)
}
