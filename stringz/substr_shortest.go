package stringz

import (
	"math"
)

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

// 找到字符串中所有字母异位词 (Find All Anagrams in a String)
// 输入：s = "cbaebabacd", p = "abc"
// 输出：[0, 6] (起始索引)
// 逻辑：找 s 中所有是 p 的异位词（排列组合）的子串。这意味着子串长度必须等于 len(p)。
func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	need := make(map[byte]int)
	for i := range len(p) {
		need[p[i]]++
	}

	window := make(map[byte]int)
	left := 0
	valid := 0

	anagrams := make([]int, 0)

	for right := range len(s) {
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		if right-left+1 >= len(p) {
			if valid == len(need) {
				anagrams = append(anagrams, left)
			}

			d := s[left]
			if _, ok := need[d]; ok {
				window[d]--
				if window[d] < need[d] {
					valid--
				}
			}

			left++
		}
	}

	return anagrams
}

func FindAnagramsOptimized(s, p string) []int {
	if len(s) < len(p) {
		return nil
	}

	// 用数组代替 map，下标 0-25 对应 a-z
	var sCount, pCount [26]int

	// 先把 p 的统计出来，同时也先把 s 的前 len(p) 个统计出来（初始窗口）
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	var res []int
	// 检查初始窗口 (index 0)
	if sCount == pCount { // Go 数组可以直接比较！这是 Map 做不到的
		res = append(res, 0)
	}

	// 开始滑动，i 是当前左指针，我们要算出新的右指针
	for i := 0; i < len(s)-len(p); i++ {
		// 移出左边的 s[i]
		sCount[s[i]-'a']--
		// 加入右边的 s[i+len(p)]
		sCount[s[i+len(p)]-'a']++

		// 此时窗口也就是 s[i+1 : i+1+len(p)]
		if sCount == pCount {
			res = append(res, i+1)
		}
	}
	return res
}

// 滑动窗口万能模板
// for right < len(s) {
//     // 1. 进窗口 (通用)
//     window[c]++
//     update(valid)

//     // 2. 判定点 (只有这里不同！)

//     // 【模式 A：找最短/最小】(如 最小覆盖子串)
//     for valid == need {
//         // 记录结果
//         left++
//     }

//     // 【模式 B：找最长】(如 无重复最长子串)
//     for invalid {
//         left++
//     }

//     // 【模式 C：固定长度】(如 找异位词)
//     if right - left + 1 == targetLen {
//         // 记录结果
//         left++
//     }
// }

// 下一步挑战：带“容错/替换”的滑动窗口
// 面试中还有一类稍难的变种，涉及“你可以修改 k 个字符”。
// 推荐题目：LeetCode 424. 替换后的最长重复字符 (Longest Repeating Character Replacement)
//
// 题目：给你一个字符串 s 和一个整数 k。你可以将字符串中的任意字符改为任意其他字符，最多改 k 次。请找出一个包含相同字母的最长子串的长度。
// 输入：s = "AABABBA", k = 1
// 输出：4
// 解释：将中间的 'B' 替换为 'A'，得到 "AAAABBA"，最长全是 A 的子串是 "AAAA"，长度 4。
// 提示：
// 这道题依然是滑动窗口（找最长），但是判断窗口是否合法的条件变了。
//
// 窗口长度：`right - left + 1`
// 窗口内出现次数最多的字符个数：`maxCount`
// 合法条件：`窗口长度 - maxCount <= k` （如果不也是主要字符的“杂质”数量 <= k，说明我们可以把杂质全替换掉，变成合法子串）。
