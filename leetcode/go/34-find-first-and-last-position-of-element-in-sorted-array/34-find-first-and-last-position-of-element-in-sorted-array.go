func searchRange(nums []int, target int) []int {
    
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := (l + r) / 2
        if nums[mid] < target {
            l = mid + 1
            continue
        }
        
        if nums[mid] > target {
            r = mid - 1
            continue
        }
        
        res := []int{
            mid, mid,
        }
        
        // Find the left most and right most
        t := mid
        for l <= t && nums[t] == nums[mid] {
            res[0] = t
            t--
        }
        
        t = mid
        for t <= r && nums[t] == nums[mid] {
            res[1] = t
            t++
        }
        
        return res
    }
    
    return []int{-1, -1}
}