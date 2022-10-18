package main

// res_1 := math.Max(0, -0)
// func longestPalindrome(s string) string {
// 	start, maxLen, l := 0, 0, len(s)
// 	for i := 0; i < l; i++ {
// 		expandAroundCenter(s, i, i, &start, &maxLen)
// 		expandAroundCenter(s, i, i+1, &start, &maxLen)
// 	}
// 	return s[start : start+maxLen]
// }

// func expandAroundCenter(s string, l, r int, start, maxLen *int) {
// 	for l >= 0 && r < len(s) && s[l] == s[r] {
// 		l, r = l-1, r+1
// 	}
// 	if r-l-1 > *maxLen {
// 		*maxLen = r - l - 1
// 		*start = l + 1
// 	}
// }

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
