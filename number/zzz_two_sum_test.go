package number

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	nums := []int{1, 3, 7, 4, 5}
	target := 9

	fmt.Println(TwoSum(nums, target))
}
