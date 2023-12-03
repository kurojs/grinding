func maxProduct(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    currMin, currMax := nums[0], nums[0]
    ans := max(currMin, currMax)
    
    for i := 1; i < len(nums); i++ {
        n := nums[i]
        if n < 0 {
            t := currMax
            currMax = currMin
            currMin = t
        }
        
        currMin = min(n, currMin * n)
        currMax = max(n, currMax * n)
        
        ans = max(ans, currMax)
    }
    
    return ans
}

func max(a,b int) int {
    if a > b {
        return a
    }
    
    return b
}


func min(a,b int) int {
    if a < b {
        return a
    }
    
    return b
}