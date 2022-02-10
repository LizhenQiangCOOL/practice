import unittest
from typing import List


class Solution:

    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        fast = slow = 0
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

    def moveZeroes_leetcode(self, nums: List[int]) -> None:
        n = len(nums)
        fast = slow = 0
        while fast < n:
            if nums[fast] != 0:
                nums[slow], nums[fast] = nums[fast], nums[slow]
                slow += 1
            fast += 1


class TestSolution(unittest.TestCase):

    def test_moveZeroes(self):
        func = Solution().moveZeroes

        testnums = [
            # 正常例子-示例
            {
                "nums": [0, 1, 0, 3, 12],
                "want": [1, 3, 12, 0, 0]
            },
            # 特殊例子
            {
                "nums": [0, 0, 0],
                "want": [0, 0, 0]
            },
            {
                "nums": [1, 2, 3],
                "want": [1, 2, 3]
            },
            {
                "nums": [],
                "want": []
            },
            # 边界测试
            {
                "nums": [0, 1, 2, 3, 0],
                "want": [1, 2, 3, 0, 0]
            },
            {
                "nums": [1, 0, 0, 0, 2],
                "want": [1, 2, 0, 0, 0]
            }
        ]

        for item in testnums:
            func(item["nums"])
            self.assertListEqual(item["nums"], item["want"])


if __name__ == "__main__":
    unittest.main()
