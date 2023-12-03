func longestPalindrome(s string) string {
	start := 0
    end := 0

	for i := 0; i < len(s)-1; i++ {
        len := max(expand(s, i, i), expand(s, i, i+1))
		if len > end - start {
            start = i - (len - 1) / 2
            end = i + len/2
		}
	}
    
	return s[start:end+1]
}

func expand(s string, left, right int) (int) {
	for left > -1 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left -1
}


func max(a,b int) int {
    if a > b {
        return a
    }
    
    return b
}