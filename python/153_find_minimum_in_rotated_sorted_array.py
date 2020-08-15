from typing import List

class Solution:
    def findMin(self, nums: List[int]) -> int:
        l = 0
        h = len(nums) - 1
        if h < 0:
            return ''
        while l + 1 < h:
            m = l + (h - l) // 2
            if nums[m] > nums[l] and nums[m] > nums[h]:
                l = m
            else:
                h = m
        return min(nums[l], nums[h])

sol = Solution()
inp = [int(x) for x in input().split(",")]
print(sol.findMin(inp))