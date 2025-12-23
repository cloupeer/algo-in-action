package dp

// 有一个机器人位于一个 m x n 网格的 左上角 （起始点 [0, 0]）。
// 机器人每次只能 向下 或者 向右 移动一步。 机器人试图达到网格的 右下角 （终点 [m-1, n-1]）。
// 问：总共有多少条不同的路径？
// 示例：m = 3, n = 7 输出：28
func UniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := range m {
		dp[i][0] = 1
	}
	for j := range n {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func UniquePathsOptimized(m, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// 1. 只要起点是石头，或者终点是石头，直接 0 种方法，不用算了
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// 申请一维数组
	dp := make([]int, n)

	// 2. 初始化第一行 (Base Case)
	// 逻辑：起点肯定是 1。
	// 只要碰到一个石头，后面全都是 0（因为只能向右走，被挡住了就过不去了）
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[j] = 0 // 其实默认就是0，这就相当于 break 的效果
			// 在这一行，石头后面的格子因为 dp 默认是 0，所以不需要额外操作
			// 但为了逻辑严谨，一旦碰到石头，这一行后面的都不用设为 1 了
			break
		} else {
			dp[j] = 1
		}
	}

	// 3. 从第二行开始遍历
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			// 情况 A：当前位置是石头
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0 // 此路不通，设为 0
				continue  // 继续看下一个格子，千万不能 break！
			}

			// 情况 B：当前位置是空地
			// dp[j] 目前存的是“上一行”的值（从上面下来的路）
			// 我们要加上“左边”的值（从左边过来的路）
			if j > 0 {
				dp[j] = dp[j] + dp[j-1]
			}
			// 如果 j == 0 (第一列)，dp[j] 就保持原样（只接受从上面下来的路径）
			// 但如果第一列上面那个位置本来就是 0 (被堵住了)，那它自然也是 0，逻辑自洽。
		}
	}

	return dp[n-1]
}

// LeetCode 64. 最小路径和 (Minimum Path Sum)
// 输入：
// [
//
//	[1, 3, 1],
//	[1, 5, 1],
//	[4, 2, 1]
//
// ]
// 输出：7
// 路径：1 -> 3 -> 1 -> 1 -> 1
func MinimumPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// dp[i] 存储当前位置的最小路径和
	// dp[i] = min(dp[i], dp[i-1])
	dp := make([]int, n)

	// 1. 初始化第一行
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// 2. 遍历剩余行
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				// 第一列：只能从上面下来
				dp[0] = dp[0] + grid[i][0]
			} else {
				// 其他列：min(上面, 左边) + 当前花费
				dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
			}
		}
	}

	return dp[n-1]
}
