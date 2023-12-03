func moveZeroes(nums []int)  {
    end := len(nums)
    for i := range nums {
        if nums[i] == 0 {
            for j := i + 1; j < end; j++ {
                if nums[j] != 0 {
                    nums[i] = nums[j]
                    nums[j] = 0
                    break
                }
            }
        }
    }
}