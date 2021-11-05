/*
 * @lc app=leetcode id=2053 lang=golang
 *
 * [2053] Kth Distinct String in an Array
 */

// @lc code=start
func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	for _, v := range arr {
		m[v]++
	}
	p := 0
	for _, v := range arr {
		if m[v] == 1 {
			p++
		}
		if p == k {
			return v
		}
	}
	return ""
}

// @lc code=end

