/*
 * @lc app=leetcode id=1837 lang=golang
 *
 * [1837] Sum of Digits in Base K
 */

// @lc code=start
func sumBase(n int, k int) int {
	remain := 0
	sum := 0

	for n != 0 {
		remain = n % k
		sum += remain
		n = n / k
	}
	return sum

}

// @lc code=end

