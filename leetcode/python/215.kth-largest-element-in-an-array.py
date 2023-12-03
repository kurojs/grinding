# @before-stub-for-debug-begin
from python3problem215 import *
from typing import List
# @before-stub-for-debug-end

#
# @lc app=leetcode id=215 lang=python3
#
# [215] Kth Largest Element in an Array
#

# @lc code=start
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:

        def mergeSort(nums: List[int]):
            if len(nums) <= 1:
                return nums
            left, right = partition(nums)
            return _merge(mergeSort(left), mergeSort(right))

        def partition(nums: List[int]) -> (List[int], List[int]):
            pivot = int(len(nums) / 2)
            return (nums[:pivot], nums[pivot:])

        def _merge(a: List[int], b: List[int]) -> List[int]:
            i, j = 0, 0
            res = []
            while i < len(a) and j < len(b):
                if a[i] >= b[j]:
                    res.append(a[i])
                    i += 1
                else:
                    res.append(b[j])
                    j += 1
            
            while i < len(a):
                res.append(a[i])
                i += 1
            while j < len(b):
                res.append(b[j])
                j += 1
            
            return res

        sorted_arr = mergeSort(nums)
        return sorted_arr[k-1]

# @lc code=end

