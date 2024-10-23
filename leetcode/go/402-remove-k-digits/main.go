package remove_k_digits

func removeKdigits(num string, k int) string {
	return trim(helper(num, k))
}

func helper(num string, k int) string {
	if k == 0 {
		return num
	}
	if k >= len(num) {
		return ""
	}

	removed := helper(num[1:], k-1)
	keeped := string(num[0]) + helper(num[1:], k)
	return min(removed, keeped)
}

func min(a, b string) string {
	for i := range a {
		if a[i] > b[i] {
			return b
		} else if b[i] > a[i] {
			return a
		}
	}

	return a
}

func trim(num string) string {
	if len(num) == 0 {
		return "0"
	}

	cut := 0
	for i := 0; i < len(num)-1; i++ {
		if string(num[i]) == "0" {
			cut++
		} else {
			break
		}
	}

	return num[cut:]
}
