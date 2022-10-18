package main

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - ((len - 1) / 2)
			end = i + (len / 2)
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	if left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1

}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
