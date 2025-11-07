package arrayhashing

// ContainsDuplicate 给定一个整数数组 nums，如果任一值在数组中出现 至少两次，返回 true；如果数组中每个元素互不相同，返回 false。
//
// 算法原理：
// 使用哈希集合（map）来记录已经访问过的元素。遍历数组时，如果当前元素已经在集合中，
// 说明出现了重复，立即返回 true；否则将元素加入集合。如果遍历完整个数组都没有重复，返回 false。
//
// 时间复杂度：O(n)，需要遍历一次数组
// 空间复杂度：O(n)，最坏情况下需要存储所有 n 个元素
//
// 优化说明：
// 使用 map[int]struct{} 而不是 map[int]bool，因为 struct{} 不占用内存空间（零大小），
// 这样可以节省内存。这是 Go 语言中实现集合的惯用方式。
//
// 示例：
// 输入： nums = [1, 2, 3, 1]
// 输出： true
// 输入： nums = [1, 2, 3, 4]
// 输出： false
func ContainsDuplicate(nums []int) bool {
	// 边界情况：数组长度小于等于1，不可能有重复
	if len(nums) <= 1 {
		return false
	}

	// 使用 map 实现集合，key 为元素值，value 使用 struct{} 节省内存
	seen := make(map[int]struct{})

	for _, num := range nums {
		// 检查当前元素是否已经在集合中
		if _, ok := seen[num]; ok {
			return true
		}

		// 将当前元素加入集合
		seen[num] = struct{}{}
	}

	return false
}
