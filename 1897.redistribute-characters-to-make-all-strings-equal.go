/*
 * @lc app=leetcode id=1897 lang=golang
 *
 * [1897] Redistribute Characters to Make All Strings Equal
 */

// @lc code=start
func makeEqual(words []string) bool {
	m := make([]int, 26)
	length := len(words)
	for _, word := range words {
		for _, ch := range []byte(word) {
			m[ch-'a']++
		}
	}
	for _, v := range m {
		if v%length != 0 {
			return false
		}
	}
	return true
}

// @lc code=end
