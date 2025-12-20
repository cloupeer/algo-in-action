package stringz

import (
	"fmt"
	"testing"
)

func Test_ClimbStairs(t *testing.T) {
	fmt.Println("==> ", ClimbStairs(3))
	fmt.Println("==> ", ClimbStairsOptimize(3))
	fmt.Println("==> ", ClimbStairs(10))
	fmt.Println("==> ", ClimbStairsOptimize(10))

	fmt.Println()
	cost := []int{10, 15, 20}
	fmt.Println("==> ", ClimbStairsCost(cost))
}
