package prefixsum

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	s := NumArray{
		prefixSum: make([]int, len(nums)+1),
	}
	for i := 1; i <= len(nums); i++ {
		s.prefixSum[i] = s.prefixSum[i-1] + nums[i-1]
	}
	return s
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSum[right+1] - this.prefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
