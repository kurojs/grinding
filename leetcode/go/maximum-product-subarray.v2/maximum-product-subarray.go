func maxProduct(nums []int) int {
	max := nums[0]
	prevMax := make([]int, len(nums))
	prevMin := make([]int, len(nums))

	if nums[0] > 0 {
		prevMax[0] = nums[0]
	} else {
		prevMin[0] = nums[0]
	}

	i := 1
	for i < len(nums) {
		if nums[i] == 0 {
			prevMax[i-1] = 0
			prevMin[i-1] = 0
		}

		if nums[i] > 0 {
			prevMax[i] = removeZero(prevMax[i-1]*nums[i], nums[i])
			prevMin[i] = prevMin[i-1] * nums[i]
		}

		if nums[i] < 0 {
			prevMax[i] = prevMin[i-1] * nums[i]
			prevMin[i] = removeZero(prevMax[i-1]*nums[i], nums[i])
		}

		if prevMax[i] > max {
			max = prevMax[i]
		}
		i++
	}

	return max
}

func removeZero(a, b int) int {
	if a == 0 {
		return b
	}

	return a
}
