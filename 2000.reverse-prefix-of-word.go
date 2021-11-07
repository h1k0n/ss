/*
 * @lc app=leetcode id=2000 lang=golang
 *
 * [2000] Reverse Prefix of Word
 */

// @lc code=start
func reversePrefix(word string, ch byte) string {
	end := -1
	for i, v := range []byte(word) {
		if v == ch {
			end = i
			break
		}
	}
	if end == -1 {
		return word
	}
	x := []byte(word)
	for i, j := 0, end; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	return string(x)
}

// @lc code=end

