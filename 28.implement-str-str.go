/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	nLen := len(needle)
	hLen := len(haystack)
	if nLen == 0 {
		return 0
	}

	for i := 0; i < hLen; i++ {
		if i+nLen > hLen {
			return -1
		}
		j := 0
		for j = 0; j < nLen; j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == nLen {
			return i
		}
	}
	return -1
}

// @lc code=end

