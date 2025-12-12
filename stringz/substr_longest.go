package stringz

// 给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
// s = "abcabcbb"  => 3 (abc)
// s = "bbbbb"  => 1 (b)
func SubStrLongest(s string) int {
	if len(s) == 0 {
		return 0
	}

	seen := make(map[rune]int)
	length := 0
	left := 0

	for i, r := range s {
		if lastIdx, ok := seen[r]; ok && lastIdx >= left {
			left = lastIdx + 1
		}

		seen[r] = i
		length = max(length, i-left+1)
	}

	return length
}

func lengthOfLongestSubstringBest(s string) int {
	// ASCII 只有 128 个字符
	// lastOccurred[x] 存储字符 x 上次出现的索引
	// 初始化为 -1 表示未出现过（或者用 0，但需要小心 index=0 的情况处理）
	// 为了方便，这里我们存索引+1，或者在逻辑里小心处理
	lastOccurred := [128]int{}
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	left := 0
	maxLength := 0

	for i, r := range s {
		// 强转 rune 为 int 作为数组下标，前提是确认 s 只有 ASCII
		// 如果题目没说只有 ASCII，这样做会 panic，还是得用 Map
		if r < 128 {
			if lastPos := lastOccurred[r]; lastPos != -1 && lastPos >= left {
				left = lastPos + 1
			}
			lastOccurred[r] = i
		} else {
			// 遇到非 ASCII 字符的处理逻辑（或者直接报错/退化回 Map）
		}

		maxLength = max(maxLength, i-left+1)
	}
	return maxLength
}
