package stringz

import (
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
