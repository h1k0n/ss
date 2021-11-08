/*
 * @lc app=leetcode id=2042 lang=golang
 *
 * [2042] Check if Numbers Are Ascending in a Sentence
 */

// @lc code=start
func isBigger(newN, oldN string) bool {
	if len(newN) > len(oldN) {
		return true
	} else if len(newN) < len(oldN) {
		return false
	}
	for i := 0; i < len(newN); i++ {
		if newN[i] > oldN[i] {
			return true
		} else if newN[i] < oldN[i] {
			return false
		}
	}
	return false
}

func areNumbersAscending(s string) bool {
	p, q := 0, 0
	length := len(s)
	var oldstr string
	for i, v := range s {
		switch {
		case v == ' ':
			q = i
		case i == length-1:
			q = i + 1
		default:
			continue
		}
		if !(s[q-1] >= '0' && s[q-1] <= '9') {
			p = i + 1
			continue
		}
		if !isBigger(s[p:q], oldstr) {
			return false
		}
		oldstr = s[p:q]
		p = i + 1
	}
	return true
}

// @lc code=end

