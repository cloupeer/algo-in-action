package stringz

import (
	"slices"
	"strings"
)

// 给定两个字符串 s 和 t ，判断它们是否是同构的。
// 如果 s 中的字符可以按某种映射关系替换得到 t，那么这两个字符串是同构的。
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
//
// 输入：s = "egg", t = "add"
// 输出：true
//
// 输入：s = "foo", t = "bar"
// 输出：false
//
// 输入：s = "paper", t = "title"
// 输出：true
func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	for i := range len(s) - 1 {
		if strings.Index(s[i+1:], s[i:i+1]) != strings.Index(t[i+1:], t[i:i+1]) {
			return false
		}
	}

	return true
}

// 0:  e->a  a->e
// 1:  g->d  d->g
// 2:  g->d?  d->g?
func IsIsomorphicMapping(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var s2t [256]byte
	var t2s [256]byte

	for i := range len(s) {
		c1 := s[i]
		c2 := t[i]

		if s2t[c1] == 0 {
			if t2s[c2] != 0 {
				return false
			}

			s2t[c1] = c2
			t2s[c2] = c1
		} else {
			if s2t[c1] != c2 || t2s[c2] != c1 {
				return false
			}
		}
	}

	return true
}

// egg  add
// 0: e->0  a->0
// 1: g->1  d->1
// 2: g=1->2  d=1->2
func IsIsomorphicIndexing(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var lastSeenS [256]int
	var lastSeenT [256]int

	for i := range len(s) {
		c1 := s[i]
		c2 := t[i]

		if lastSeenS[c1] != lastSeenT[c2] {
			return false
		}

		lastSeenS[c1] = i + 1
		lastSeenT[c2] = i + 1
	}

	return true
}

func IsIsomorphicGeneric(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return slices.Equal(GetStructure(s), GetStructure(t))
}

func GetStructure(s string) []int {
	ids := make([]int, 0, len(s))

	mapping := make(map[rune]int)
	nextID := 0

	for _, v := range s {
		if id, exist := mapping[v]; exist {
			ids = append(ids, id)
			continue
		}

		nextID++
		mapping[v] = nextID
		ids = append(ids, nextID)
	}

	return ids
}

// GetStructureGeneric 泛型核心函数
// T comparable: 约束 T 必须是支持 == 比较的类型 (如 int, string, rune, byte 等)
// 输入: 任意类型的切片 []T
// 输出: 结构指纹 []int
func GetStructureGeneric[T comparable](input []T) []int {
	// 预分配内存，避免扩容
	ids := make([]int, 0, len(input))

	// mapping 的 key 变成了泛型 T
	mapping := make(map[T]int)
	nextID := 0

	for _, v := range input {
		// 逻辑与之前完全一致：
		// 见过 -> 复用 ID
		// 没见过 -> 分配新 ID
		if id, exists := mapping[v]; exists {
			ids = append(ids, id)
		} else {
			mapping[v] = nextID
			ids = append(ids, nextID)
			nextID++
		}
	}
	return ids
}
