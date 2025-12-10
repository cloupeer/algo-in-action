package stringz

// 题目：找出字符串第一个匹配项的下标
func SubStrIndex(s, substr string) int {
	if s == substr {
		return 0
	}

	m, n := len(s), len(substr)

	if n == 0 {
		return 0
	}

	if m == 0 || m < n {
		return -1
	}

	// 计算 sub 长度 max ，是双指针最大区间
	// left 指向 str 第一个字符，不断向右移动，直到找到与 sub 第一个字符相同的索引
	// right = left + 1 ； right <= max; 对比 str[right] 和 sub[i] 是否相等，直到遍历至 max
	// 相等，则意味着找到了，返回 left
	// 不相等，则当前区间不满足，继续向右移动 left，重复这个过程
	for l := 0; l <= m-n; l++ {
		if s[l] != substr[0] {
			continue
		}

		match := true
		for i := 1; i < n; i++ {
			if s[l+i] != substr[i] {
				match = false
				break
			}
		}

		if match {
			return l
		}
	}

	return -1
}

func SubStrIndexBySliceCompare(s, substr string) int {
	if s == substr {
		return 0
	}

	m, n := len(s), len(substr)

	if n == 0 {
		return 0
	}

	if m == 0 || m < n {
		return -1
	}

	for i := 0; i <= m-n; i++ {
		if s[i:i+n] == substr {
			return i
		}
	}

	return -1
}
