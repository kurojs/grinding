function findPeakElement(nums: number[]): number {
    const diffs = nums.map((n, i) => {
        let leftDiff = i > 0 ? n - nums[i-1] : 0;
        let rightDiff = i < nums.length - 1 ? n - nums[i+1] : 0;
        if (leftDiff >= 0 && rightDiff >= 0) {
            return leftDiff + rightDiff;
        } else {
            return -Infinity;
        }
    });
    
    // Find max in diffs
    let max = diffs[0];
    let maxIdx = 0;
    diffs.forEach((diff, i) => {
        if (max < diff) {
            max = diff;
            maxIdx = i;
        }
    });
    
    return maxIdx;
};