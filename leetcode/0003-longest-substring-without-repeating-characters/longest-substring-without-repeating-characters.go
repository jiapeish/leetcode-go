package _003_longest_substring_without_repeating_characters


func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var times [256]int
	output, left, right := 0, 0, -1

	length := len(s)
	for left < length {
		if right+1 < length && times[s[right+1] - 'a'] == 0 {
			times[s[right+1] - 'a']++
			right++
		} else {
			times[s[left] - 'a']--
			left++
		}
		output = max(output, right-left+1)
	}
	return output
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstringBitmap(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitmap [256]bool
	output, left, right := 0, 0, 0

	length := len(s)
	for left < length {
		if bitmap[s[right]] == false {
			bitmap[s[right]] = true
			right++
		} else {
			bitmap[s[left]] = false
			left++
		}

		if output < right - left {
			output = right - left
		}

		if left+output >= length || right >= length {
			break
		}
	}
	return output
}