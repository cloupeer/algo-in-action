package arrayhashing

import (
	"reflect"
	"sort"
	"testing"
)

// ==================== TwoSum 测试用例 ====================

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "基本示例",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "负数情况",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "正负数混合",
			nums:     []int{3, 2, 4, -1, 5},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "包含零",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "大数",
			nums:     []int{1000000, 2000000, 3000000},
			target:   5000000,
			expected: []int{1, 2},
		},
		{
			name:     "找不到答案",
			nums:     []int{1, 2, 3, 4},
			target:   10,
			expected: nil,
		},
		{
			name:     "空数组",
			nums:     []int{},
			target:   0,
			expected: nil,
		},
		{
			name:     "单个元素",
			nums:     []int{1},
			target:   2,
			expected: nil,
		},
		{
			name:     "两个元素找到答案",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "重复元素但不同索引",
			nums:     []int{3, 2, 3},
			target:   6,
			expected: []int{0, 2},
		},
		{
			name:     "目标值等于两个相同元素",
			nums:     []int{5, 5, 5},
			target:   10,
			expected: []int{0, 1},
		},
		{
			name:     "长数组",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			target:   29,
			expected: []int{13, 14},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TwoSum(tt.nums, tt.target)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}

// ==================== ContainsDuplicate 测试用例 ====================

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "基本示例-有重复",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "基本示例-无重复",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "空数组",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "单个元素",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "两个相同元素",
			nums:     []int{1, 1},
			expected: true,
		},
		{
			name:     "两个不同元素",
			nums:     []int{1, 2},
			expected: false,
		},
		{
			name:     "多个重复元素",
			nums:     []int{1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "负数",
			nums:     []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "包含零",
			nums:     []int{0, 1, 2, 0},
			expected: true,
		},
		{
			name:     "大数",
			nums:     []int{1000000, 2000000, 1000000},
			expected: true,
		},
		{
			name:     "长数组无重复",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: false,
		},
		{
			name:     "长数组有重复",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			expected: true,
		},
		{
			name:     "重复在末尾",
			nums:     []int{1, 2, 3, 4, 5, 5},
			expected: true,
		},
		{
			name:     "重复在开头",
			nums:     []int{1, 1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "所有元素相同",
			nums:     []int{7, 7, 7, 7, 7},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ContainsDuplicate(tt.nums)
			if result != tt.expected {
				t.Errorf("ContainsDuplicate(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}

// ==================== GroupAnagram 测试用例 ====================

// sortGroups 对结果进行排序以便比较
// 因为 map 的遍历顺序是不确定的，所以需要对结果进行排序
func sortGroups(groups [][]string) [][]string {
	// 对每个组内的字符串排序
	for i := range groups {
		sort.Strings(groups[i])
	}
	// 对组进行排序（按组内第一个字符串排序）
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 && len(groups[j]) == 0 {
			return false
		}
		if len(groups[i]) == 0 {
			return true
		}
		if len(groups[j]) == 0 {
			return false
		}
		return groups[i][0] < groups[j][0]
	})
	return groups
}

// compareGroups 比较两个分组结果是否相同（忽略顺序）
func compareGroups(got, want [][]string) bool {
	gotSorted := sortGroups(got)
	wantSorted := sortGroups(want)
	return reflect.DeepEqual(gotSorted, wantSorted)
}

func TestGroupAnagram(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected [][]string
	}{
		{
			name: "基本示例",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name:     "空数组",
			strs:     []string{},
			expected: [][]string{},
		},
		{
			name: "单个字符串",
			strs: []string{"abc"},
			expected: [][]string{
				{"abc"},
			},
		},
		{
			name: "两个相同字符串",
			strs: []string{"abc", "abc"},
			expected: [][]string{
				{"abc", "abc"},
			},
		},
		{
			name: "两个不同异位词",
			strs: []string{"abc", "cba"},
			expected: [][]string{
				{"abc", "cba"},
			},
		},
		{
			name: "包含空字符串",
			strs: []string{"", "a", ""},
			expected: [][]string{
				{"", ""},
				{"a"},
			},
		},
		{
			name: "所有字符串都是异位词",
			strs: []string{"abc", "bca", "cab"},
			expected: [][]string{
				{"abc", "bca", "cab"},
			},
		},
		{
			name: "所有字符串都不同",
			strs: []string{"abc", "def", "ghi"},
			expected: [][]string{
				{"abc"},
				{"def"},
				{"ghi"},
			},
		},
		{
			name: "多个分组",
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat", "tab"},
			expected: [][]string{
				{"bat", "tab"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			name: "单个字符",
			strs: []string{"a", "b", "a", "b"},
			expected: [][]string{
				{"a", "a"},
				{"b", "b"},
			},
		},
		{
			name: "长字符串",
			strs: []string{"listen", "silent", "enlist"},
			expected: [][]string{
				{"enlist", "listen", "silent"},
			},
		},
		{
			name: "重复的异位词",
			strs: []string{"eat", "tea", "eat", "tea"},
			expected: [][]string{
				{"eat", "eat", "tea", "tea"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupAnagram(tt.strs)
			if !compareGroups(result, tt.expected) {
				t.Errorf("GroupAnagram(%v) = %v, want %v", tt.strs, result, tt.expected)
			}
		})
	}
}

// ==================== 性能基准测试 ====================

func BenchmarkTwoSum(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	target := 1998

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TwoSum(nums, target)
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsDuplicate(nums)
	}
}

func BenchmarkGroupAnagram(b *testing.B) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	for i := 0; i < 100; i++ {
		strs = append(strs, strs...)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupAnagram(strs)
	}
}
