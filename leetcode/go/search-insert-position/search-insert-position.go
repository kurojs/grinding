func searchInsert(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right {
        pivot := (left + right) / 2
        
        if nums[pivot] == target {
            return pivot
        }
        
        if nums[pivot] > target {
            right = pivot - 1
        } else {
            left = pivot + 1
        }
    }
    
    return left
}