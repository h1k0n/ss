/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	length := len(digits)
	res := make([]int, length+1)

	plusN := 1
	for i := length - 1; i >= 0; i-- {
		num := digits[i] + plusN
		if num >= 10 {
			res[i+1] = num - 10
			plusN = 1
		} else {
			res[i+1] = num
			plusN = 0
		}
	}
	if plusN == 1 {
		res[0] = 1
		return res
	}
	return res[1:]
}

// @lc code=end

