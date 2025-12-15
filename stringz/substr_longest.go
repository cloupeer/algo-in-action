package stringz

// 给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
// s = "abcabcbb"  => 3 (abc)
// s = "bbbbb"  => 1 (b)
func LongestSubString(s string) int {
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

// 滑动窗口通用范式：
//  1. [核心] 维护一个动态区间 [left, right]，通过双指针操作。
//  2. [扩张] 右指针(right)主动向右探索，寻找可行解。
//  3. [收缩] 当窗口打破约束条件(如出现重复字符)时，左指针(left)被动向右收缩，
//     通过移除左侧元素(s[left])来恢复窗口的合法性。
//  4. [结算] 每次窗口状态合法时，计算当前长度并更新全局最大值。
func LongestSubStringGeneral(s string) int {
	if len(s) == 0 {
		return 0
	}

	var chars [256]bool
	longest := 0
	left := 0
	for right := range s {
		c := s[right]

		for chars[c] {
			chars[s[left]] = false
			left++
		}

		chars[c] = true
		longest = max(longest, right-left+1)
	}

	return longest
}

func LongestSubStringBest(s string) int {
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

// 至多包含两个不同字符的最长子串
// "abcbdc" => bcb
func LongestSubStringTwoChar(s string) int {
	if len(s) == 0 {
		return 0
	}

	runes := []rune(s)

	// 子串中字符的个数
	seen := make(map[rune]int)
	// 子串长度
	length := 0
	// 左指针
	left := 0

	// 另一种更简洁的写法是：先进行 seen[c]++ 再判断 len(seen) > 2 ，这样可以省略 ok 的判断
	for right, c := range runes {
		_, ok := seen[c]

		// 如果子串条件不满足，则停止探测，改变子串范围
		// 来了 新字符 且 坑位已满
		for !ok && len(seen) >= 2 {
			b := runes[left]
			seen[b]--
			if seen[b] == 0 {
				delete(seen, b)
			}
			left++
		}

		// 子串条件满足，右指针继续移动探测
		seen[c]++
		length = max(length, right-left+1)
	}

	return length
}
