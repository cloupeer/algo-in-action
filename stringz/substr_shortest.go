package stringz

import "math"

// 最小覆盖子串
// 在字符串 s 中找到一个长度最短的子串，这个子串必须包含字符串 t 中所有的字符（包括重复字符的数量）。
//
//	包含所有：如果 t 里有两个 'A'，你的子串里也至少得有两个 'A'。
//	顺序不限：子串里的字符顺序无所谓，只要数量够就行。
//	最短：满足上述条件的子串可能有很多个，我们要找最短的那个。
//
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC" (因为它是所有满足条件的子串里最短的)
func MinimumWindowSubstring(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	for i := range len(t) {
		need[t[i]]++
	}

	window := make(map[byte]int)
	matchCount := 0
	left := 0

	start, minLen := 0, math.MaxInt32

	for right := range len(s) {
		c := s[right]

		// 只有需要的字符才更新窗口
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				matchCount++
			}
		}

		for matchCount == len(need) {
			if right-left+1 < minLen {
				start = left
				minLen = right - left + 1
			}

			d := s[left]
			left++

			if _, ok := need[d]; ok {
				window[d]--
				if window[d] < need[d] {
					matchCount--
				}
			}
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[start : start+minLen]
}
