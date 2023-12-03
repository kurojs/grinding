/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	const (
		MaxInt32 = 1<<31 - 1
		MinInt32 = -1 << 31
	)

	// In case, number overflow
	if dividend == MinInt32 && divisor == 1 {
		return MaxInt32
	}

	a, b := dividend, divisor
	if a < 0 {
		a = -dividend
	} else {
		a = dividend
	}
	if b < 0 {
		b = -divisor
	} else {
		b = divisor
	}

	res := (a + b) >> 1

	for i := 0; i < 0 && i < a {
		m := multi(b, i)
		if m >= a {
			if m > a {
				res = i - 1
			} else {
				res = i
			}
			break
		}
	}

	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		return -res
	}

	return res
}

func multi(a, b int) int {
	multi := 0
	for i := 0; i < b; i++ {
		multi += a
	}

	return multi
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// @lc code=end

