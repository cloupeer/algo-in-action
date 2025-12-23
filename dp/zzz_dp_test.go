package dp

import (
	"fmt"
	"testing"
)

func Test_ClimbStairs(t *testing.T) {
	fmt.Println("==> ", ClimbStairs(3))
	fmt.Println("==> ", ClimbStairsOptimized(3))
	fmt.Println("==> ", ClimbStairs(10))
	fmt.Println("==> ", ClimbStairsOptimized(10))

	fmt.Println()
	cost := []int{10, 15, 20}
	fmt.Println("==> ", ClimbStairsCost(cost))
}

func Test_Robot(t *testing.T) {
	fmt.Println("==> ", UniquePaths(3, 7))
	fmt.Println("==> ", UniquePathsOptimized(3, 7))
}

func Test_MinimumPathSum(t *testing.T) {
	weights := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println("==> ", MinimumPathSum(weights))
}
