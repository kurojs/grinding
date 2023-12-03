# @before-stub-for-debug-begin
from python3problem81 import *
from typing import List
# @before-stub-for-debug-end

#
# @lc app=leetcode id=81 lang=python3
#
# [81] Search in Rotated Sorted Array II
#

# @lc code=start
class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        size = len(nums)
        if size == 0:
            return False
        
        for i in range(0, size):
            if nums[i] == target:
                return True
            if i == size - 1:
                return target == nums[i]

            if nums[i] > nums[i+1]:
                if self.bs(nums=nums, target=target, left=0, right=i):
                    return True
                else:
                    return self.bs(nums=nums, target=target, left=i+1, right=size-1)
        return False

    def bs(self, nums: List[int], target: int, left: int, right: int) -> bool:
        while left <= right:
            pivot = int((left + right) / 2)
            mid = nums[pivot]
            if target == mid:
                return True
            if target < mid:
                right = pivot - 1
            else:
                left = pivot + 1
        return False
# @lc code=end
