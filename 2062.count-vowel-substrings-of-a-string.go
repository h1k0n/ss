/*
 * @lc app=leetcode id=2062 lang=golang
 *
 * [2062] Count Vowel Substrings of a String
 */

// @lc code=start
func checkVowel(substr string) bool {
	a, e, i, o, u := 0, 0, 0, 0, 0
	for _, v := range substr {
		switch v {
		case 'a':
			a++
		case 'e':
			e++
		case 'i':
			i++
		case 'o':
			o++
		case 'u':
			u++
		default:
			return false
		}
	}
	if a == 0 || e == 0 || i == 0 || o == 0 || u == 0 {
		return false
	}
	return true
}

func countVowelSubstrings(word string) int {
	length := len(word)
	count := 0
	for i := 0; i+4 < length; i++ {
		for j := i + 4; j < length; j++ {
			if checkVowel(word[i : j+1]) {
				count++
			}
		}
	}
	return count
}

// @lc code=end

