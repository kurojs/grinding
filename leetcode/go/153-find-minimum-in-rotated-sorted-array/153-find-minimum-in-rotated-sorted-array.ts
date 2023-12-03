function findMin(nums: number[]): number {
    let l = 0
    let r = nums.length - 1
    
    while (l <= r) {
        const mid = Math.floor((l + r) / 2)
        
        if (mid < nums.length - 1) {
            if (nums[mid] > nums[mid + 1]) {
                return nums[mid + 1]
            }
        }
        
        if (mid > 0) {
            if (nums[mid] < nums[mid-1]) {
                return nums[mid]
            }
        }
        
        if (nums[l] < nums[mid]) {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    
    return nums[0]
};