package stringz

import (
	"fmt"
	"testing"
)

func Test_SubIndex(t *testing.T) {
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
