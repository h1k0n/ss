/*
 * @lc app=leetcode id=1800 lang=golang
 *
 * [1800] Maximum Ascending Subarray Sum
 */

// @lc code=start
func maxAscendingSum(nums []int) int {
	queueLast := -1
	sum := 0
	max := 0
	for i, v := range nums {
		if v > queueLast { //只管累加
			sum += v
			queueLast = v
		} else { //比较和，清空队列，重新进队
			if sum > max {
				max = sum
			}
			sum = 0
			queueLast = v
			sum += v
		}
		if i == len(nums)-1 {
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// @lc code=end

