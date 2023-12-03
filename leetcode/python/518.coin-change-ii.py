# @before-stub-for-debug-begin
from python3problem518 import *
from typing import *
# @before-stub-for-debug-end

from typing import List
#
# @lc app=leetcode id=518 lang=python3
#
# [518] Coin Change II
#

# @lc code=start
class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        def coinsChange(amount: int, idx: int) -> int:
            if amount == 0:
                return 1
            if idx == len(coins):
                return 0
            if memo[idx][amount] != -1:
                return memo[idx][amount]
            
            if coins[idx] > amount:
                memo[idx][amount] = coinsChange(amount, idx+1)
            else:
                memo[idx][amount] = (
                    coinsChange(amount-coins[idx], idx) + coinsChange(amount, idx+1)
                )
            
            return memo[idx][amount]

        memo = [[-1] * (amount +1) for _ in range(len(coins))]
        return coinsChange(amount, 0)
# class Solution:
#     def change(self, amount: int, coins: List[int]) -> int:
#         def coinsChange(amount: int, idx: int) -> int:
#             if amount == 0:
#                 return 1
#             if idx == len(coins):
#                 return 0
#             if memo[idx][amount] != -1:
#                 return memo[idx][amount]
            
#             if coins[idx] > amount:
#                 memo[idx][amount] = coinsChange(amount, idx+1)
#             else:
#                 memo[idx][amount] = (
#                     coinsChange(amount-coins[idx], idx) + coinsChange(amount, idx+1)
#                 )
            
#             return memo[idx][amount]

#         memo = [[-1] * (amount +1) for _ in range(len(coins))]
#         return coinsChange(amount, 0)
# @lc code=end
