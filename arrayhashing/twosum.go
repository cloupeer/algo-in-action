package arrayhashing

// TwoSum 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 算法原理：
// 使用哈希表（map）存储已访问的元素及其索引，在一次遍历中完成查找。
// 对于每个元素 nums[i]，计算 complement = target - nums[i]，
// 如果 complement 已经在哈希表中，说明找到了两个数的和为 target。
//
// 时间复杂度：O(n)，只需遍历一次数组
// 空间复杂度：O(n)，最坏情况下需要存储 n-1 个元素
//
// 示例：
// 输入： nums = [2, 7, 11, 15], target = 9
// 输出： [0, 1]
// 解释： 因为 nums[0] + nums[1] == 9 ，所以返回 [0, 1]
func TwoSum(nums []int, target int) []int {
	// 边界情况：数组长度小于2，无法找到两个数
	if len(nums) < 2 {
		return nil
	}

	// 使用哈希表存储已访问的元素值及其索引
	// key: 元素值, value: 元素索引
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		// 检查 complement 是否已经在哈希表中
		if j, ok := seen[complement]; ok {
			// 找到答案，返回两个索引（j < i，因为 j 是之前访问过的）
			return []int{j, i}
		}

		// 将当前元素存入哈希表，供后续查找使用
		seen[num] = i
	}

	// 没有找到答案
	return nil
}
