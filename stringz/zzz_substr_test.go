package stringz

import (
	"fmt"
	"testing"
)

func Test_SubStrIndex(t *testing.T) {
	fmt.Println("==> ", SubStrIndex("sadbutsad", "sad"))
	fmt.Println("==> ", SubStrIndex("sabbbutsad", "sad"))
	fmt.Println("==> ", SubStrIndex("abcsadbbutsad", "sad"))
	fmt.Println("==> ", SubStrIndex("butsad", "sad"))

	fmt.Println()

	fmt.Println("==> ", SubStrIndexBySliceCompare("sadbutsad", "sad"))
	fmt.Println("==> ", SubStrIndexBySliceCompare("sabbbutsad", "sad"))
	fmt.Println("==> ", SubStrIndexBySliceCompare("abcsadbbutsad", "sad"))
	fmt.Println("==> ", SubStrIndexBySliceCompare("butsad", "sad"))
}

func Test_SubStrLongest(t *testing.T) {
	fmt.Println("==> ", LongestSubString("abcabcbb"))
	fmt.Println("==> ", LongestSubString("bbbbb"))
	fmt.Println("==> ", LongestSubString("pwwkew"))

	fmt.Println()

	fmt.Println("==> ", LongestSubStringGeneral("abcabcbb"))
	fmt.Println("==> ", LongestSubStringGeneral("bbbbb"))
	fmt.Println("==> ", LongestSubStringGeneral("pwwkew"))

	fmt.Println()

	fmt.Println("==> ", LongestSubStringBest("abcabcbb"))
	fmt.Println("==> ", LongestSubStringBest("bbbbb"))
	fmt.Println("==> ", LongestSubStringBest("pwwkew"))
}

func Test_SubStringLongestTwoChar(t *testing.T) {
	fmt.Println("==> ", LongestSubStringTwoChar("abcbdc"))
}

func Test_MinimumWindowSubstring(t *testing.T) {
	fmt.Println("==> ", MinimumWindowSubstring("ADOBECODEBANC", "ABC"))
}

func Test_FindAllAnagrams(t *testing.T) {
	fmt.Println("==> ", FindAnagrams("cbaebabacd", "abc"))
	fmt.Println("==> ", FindAnagramsOptimized("cbaebabacd", "abc"))
}
