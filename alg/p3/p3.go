package p3

func lengthOfLongestSubstring(s string) int {
	var asciiCount [128]int
	i, ans := 0, 0

	for j := 0; j < len(s); j++ {
		i = maxInt(asciiCount[int(s[j])], i)
		ans = maxInt(ans, j-i+1)
		asciiCount[int(s[j])] = j + 1
	}
	return ans
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
