package arrayhashing

// GroupAnagram 给定一个字符串数组 strs，将 字母异位词 组合在一起。你可以按任意顺序返回答案列表。
// 字母异位词 (Anagram) 是指由基本相同字母组成的、只是字母顺序不同的单词。
//
// 算法原理：
// 使用字符计数数组作为哈希表的 key。对于每个字符串，统计每个字符（'a'-'z'）出现的次数，
// 形成一个长度为 26 的整数数组。具有相同字符计数的字符串就是字母异位词，会被分到同一组。
//
// 时间复杂度：O(n*k)，其中 n 是字符串数量，k 是平均字符串长度
// 空间复杂度：O(n*k)，需要存储所有字符串
//
// 优化说明：
// 1. 使用 [26]int 数组作为 key 比排序字符串更高效（排序需要 O(k*log(k))，计数只需 O(k)）
// 2. 使用 make([][]string, 0, len(groups)) 预分配容量，减少内存重新分配
// 3. 添加字符范围检查，确保只处理小写字母，避免数组越界
//
// 示例：
// 输入： strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出： [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
func GroupAnagram(strs []string) [][]string {
	// 边界情况：空数组
	if len(strs) == 0 {
		return [][]string{}
	}

	// 使用字符计数数组作为 key，将字母异位词分组
	// key: [26]int 数组，表示每个字符的出现次数
	// value: []string，属于同一组的字符串列表
	groups := make(map[[26]int][]string)

	for _, s := range strs {
		// 边界情况：空字符串单独处理
		if s == "" {
			groups[[26]int{}] = append(groups[[26]int{}], s)
			continue
		}

		// 统计每个字符的出现次数
		var key [26]int
		for _, char := range s {
			// 字符范围检查：确保是小写字母 'a'-'z'
			if char < 'a' || char > 'z' {
				// 如果包含非小写字母，可以跳过或使用其他处理方式
				// 这里假设输入都是小写字母，但添加检查更安全
				continue
			}
			key[char-'a']++
		}

		// 将当前字符串添加到对应的组中
		groups[key] = append(groups[key], s)
	}

	// 将 map 中的值转换为二维切片
	// 预分配容量，提高性能
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
