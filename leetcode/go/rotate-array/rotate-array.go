func rotate(nums []int, k int)  {
    step := k % len(nums)
    head := nums[:len(nums) - step]
    tail := nums[len(nums) - step:]
    
    for i, n := range append(tail, head...) {
        nums[i] = n
    }
}