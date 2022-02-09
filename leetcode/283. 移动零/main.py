
from typing import List

class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        fast = slow =0
        while fast < len(nums):
            if nums[fast] != 0 and nums[slow] == 0:
                nums[slow], nums[fast] = nums[fast], nums[slow]
                fast += 1
                slow += 1
                # 防止数组越界
                continue

            if fast == slow:
                fast += 1
                # 防止数组越界
                continue 

            if nums[fast] == 0:
                fast += 1
            if nums[slow] != 0:
                slow += 1
            