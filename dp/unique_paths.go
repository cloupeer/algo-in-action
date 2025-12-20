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
