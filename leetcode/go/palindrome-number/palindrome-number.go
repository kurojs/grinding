func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    if x < 10 {
        return true
    }
    
    nums := make([]int, 0)
    for x > 0 {
        nums = append(nums, x % 10)
        x = x / 10
    }
    
    l, r := 0, len(nums)-1
    for l < r {
        if nums[l] != nums[r] {
            return false
        }
        l++
        r--
    }
    
    return true
}
