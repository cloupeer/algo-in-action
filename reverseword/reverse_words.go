package reverseword

import "strings"

func ReverseWordsFields(s string) string {
	words := strings.Fields(s)

	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}

func ReverseWordsOptimized(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))

	right := len(s) - 1

	for right >= 0 {
		// 跳过右边所有空格
		if s[right] == ' ' {
			right--
			continue
		}

		// 寻找右边第一个单词的范围
		left := right
		for left >= 0 && s[left] != ' ' {
			left--
		}

		// 将单词写入 sb（写入之前判断是否曾经写入过，在单词前加空格）
		if sb.Len() > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(s[left+1 : right+1])

		// 此时，将 right 移动到 left 位置，继续下一轮
		right = left
	}

	return sb.String()
}
