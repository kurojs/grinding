func findMedianSortedArrays(a []int, b []int) float64 {
	merged := mergeArrays(a, b)
	median := (len(a) + len(b)) / 2

	// Need avg of two medians
	if (len(a)+len(b))%2 == 0 {
		return float64(merged[median]+merged[median-1]) / 2
	}

	return float64(merged[median])
}

func mergeArrays(a []int, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		for i < len(a) && j < len(b) && a[i] < b[j] {
			res = append(res, a[i])
			i++
		}

		for i < len(a) && j < len(b) && a[i] >= b[j] {
			res = append(res, b[j])
			j++
			continue
		}
	}

	for i < len(a) {
		res = append(res, a[i])
		i++
	}
	for j < len(b) {
		res = append(res, b[j])
		j++
	}

	return res
}