/*
 * @lc app=leetcode id=1700 lang=golang
 *
 * [1700] Number of Students Unable to Eat Lunch
 */

// @lc code=start
func countStudents(students []int, sandwiches []int) int {
	length := len(students)
	p, q := 0, length-1
	x := 0
	var allDontWant bool
	for x < length {
		if students[p] == sandwiches[x] {
			x++
			p++
		} else {
			allDontWant = true
			for i := p + 1; i <= q; i++ {
				if students[i] == sandwiches[x] {
					allDontWant = false
				}
			}
			if allDontWant {
				return q - p + 1
			}
			students = append(students, students[p])
			p++
			q++
		}
	}
	return 0
}

// @lc code=end

