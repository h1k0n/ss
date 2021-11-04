/*
 * @lc app=leetcode id=2057 lang=golang
 *
 * [2057] Smallest Index With Equal Value
 */

// @lc code=start
func smallestEqual(nums []int) int {
	for i, v := range nums {
		if i%10 == v {
			return i
		}
	}
	return -1
}

// @lc code=end

