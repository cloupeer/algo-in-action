package number

// 两数之和
// 给定一个整数数组 nums  和一个整数目标值 target, 请你在该数组中找出和为目标值 target  的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
// 输入： nums = [2, 7, 11, 15]    target = 9
// 输出： [0, 1]
func TwoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))

	for i, num := range nums {
		// 检查 target - current 是否已经在 map 中
		if idx, ok := numMap[target-num]; ok {
			return []int{idx, i}
		}
		// 将当前数字存入 map
		numMap[num] = i
	}

	return nil
}

func TwoSum2(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	var indexes []int

	left := 0

	for i := 1; i < len(nums); i++ {
		if nums[left]+nums[i] == target {
			indexes = append(indexes, left, i)
			break
		}

		left = i
	}

	return indexes
}
