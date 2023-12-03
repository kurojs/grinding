# @before-stub-for-debug-begin
from python3problem239 import *
from typing import *
# @before-stub-for-debug-end

#
# @lc app=leetcode id=239 lang=python3
#
# [239] Sliding Window Maximum
#

# @lc code=start
from typing import List


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        def findMax(nums: List[int]) -> int:
            m = nums[0]
            for n in nums:
                if n > m:
                    m = n

            return m

        left, right = 0, k
        window = nums[left:right]
        current_max = findMax(window)
        max_nums = [current_max]

        # Find max in window O(k)
        # Move window next -> exclude prev -> add new
        # What if we remove max ?
        while right < len(nums):
            next_r = right
            if next_r == len(nums):
                return max_nums

            next_right = nums[next_r]
            window = nums[left+1:right+1]
            if next_right >= current_max:
                current_max = next_right
            elif nums[left] == current_max:
                current_max = findMax(window)

            max_nums.append(current_max)
            right += 1
            left += 1

        return max_nums


        
# @lc code=end

