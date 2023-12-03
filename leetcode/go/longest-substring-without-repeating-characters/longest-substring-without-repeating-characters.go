func lengthOfLongestSubstring(s string) int {
	start := 0
	end := 0
	max := 0
	windows := make(map[byte]int)

	for end < len(s) {
		char := s[end]
		windows[char]++

		// Remove one by one characters from sub string until it's unique again
		for windows[char] > 1 {
			windows[s[start]]--
			start++
		}

		end++

		if max < (end - start) {
			max = end - start
		}
	}

	return max
}