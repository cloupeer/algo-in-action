package dp

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。 每次你可以爬 1 或 2 个台阶。 你有多少种不同的方法可以爬到楼顶？
// 走到 10 阶的方法数 = 走到 9 阶的方法数 + 走到 8 阶的方法数。
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	d := make([]int, n+1)

	d[1] = 1
	d[2] = 2

	for i := 3; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}

	return d[n]
}

func ClimbStairsOptimized(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1
	prev1 := 2

	for i := 3; i <= n; i++ {
		current := prev1 + prev2

		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// 题目：使用最小花费爬楼梯 (LeetCode 746)
// 给你一个整数数组 cost，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。
// 一旦你支付此费用，你可以选择向上爬一个或者两个台阶。
// 你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
// 请你计算并返回达到楼梯顶部的最低花费。
// 输入：cost = [10, 15, 20] 输出：15
// 解释：从下标 1 (花费 15) 开始。往上走两步，直接登顶。总花费 15。
// dp[i] 代表什么？
// 公式是什么？ (min 应该出现在哪里？)
// 初始值是什么？
func ClimbStairsCost(cost []int) int {
	n := len(cost)

	// dp[i] 表示到达第 i 阶台阶的最小花费
	// 我们需要计算到第 n 阶（楼顶），所以长度是 n+1
	dp := make([]int, n+1)

	// 初始值：你可以自由选择从 0 或 1 开始，所以到达这两个位置的初始花费都是 0
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		// 方案1：从 i-1 阶上来，花费 = 到达 i-1 的钱 + i-1 本身的过路费
		onestep := dp[i-1] + cost[i-1]
		// 方案2：从 i-2 阶上来，花费 = 到达 i-2 的钱 + i-2 本身的过路费
		twostep := dp[i-2] + cost[i-2]

		dp[i] = min(onestep, twostep)
	}

	return dp[n]
}
