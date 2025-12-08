package stringz

func ReverseWordsTwoPointers(s string) string {

	b := []byte(s)

	start := 0
	for i := range len(b) {
		if b[i] == ' ' {
			reverseword2(b, start, i-1)
			start = i + 1
		}
	}

	reverseword2(b, start, len(b)-1)
	return string(b)
}

func reverseword2(b []byte, left int, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
