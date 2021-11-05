/*
 * @lc app=leetcode id=2047 lang=golang
 *
 * [2047] Number of Valid Words in a Sentence
 */

// @lc code=start
func checkPass(word string) bool {
	if len(word) == 0 {
		return false
	}
	hCount, pCount := 0, 0
	for i, v := range word {
		if v >= '0' && v <= '9' {
			return false
		}
		if i == 0 && v == '-' {
			return false
		}
		if i == len(word)-1 && v == '-' {
			return false
		}
		if v == '-' {
			hCount++
		}
		if hCount >= 2 {
			return false
		}

		if v == ',' || v == '.' || v == '!' {
			pCount++
			if i != len(word)-1 {
				return false
			}
		}
		if pCount >= 2 {
			return false
		}
		if (v == ',' || v == '.' || v == '!') && i > 1 && word[i-1] == '-' {
			return false
		}
	}
	return true
}

func countValidWords(sentence string) int {
	p := 0
	count := 0
	for i, v := range sentence {
		if v == ' ' {
			if checkPass(sentence[p:i]) {
				count++
			}
			p = i + 1
		} else if i == len(sentence)-1 {
			if checkPass(sentence[p : i+1]) {
				count++
			}
		}
	}
	return count
}

// @lc code=end

