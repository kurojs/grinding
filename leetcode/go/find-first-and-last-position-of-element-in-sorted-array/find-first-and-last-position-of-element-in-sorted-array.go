func searchRange(nums []int, target int) []int {
    foundIdx := binarySearch(nums, 0, len(nums) - 1, target)
    if foundIdx < 0 {
        return []int{-1, -1}
    }
    
    l, r := foundIdx, foundIdx
    res := []int{l, r}
    
    for l >= 0 && nums[l] == nums[foundIdx] {
        res[0] = l
        l--
    }
    
    for r < len(nums) && nums[r] == nums[foundIdx]{
        res[1] = r
        r++
    }
    
    return res
}

func binarySearch(nums []int, start, end, target int) int {
    if start > end {
        return -1
    }
    
    
    mid := start + (end - start) / 2
    
    if nums[mid] == target {
        return mid
    }
    
    if nums[mid] > target {
        return binarySearch(nums, start, mid-1, target)
    }
    
    return binarySearch(nums, mid+1, end, target)
}