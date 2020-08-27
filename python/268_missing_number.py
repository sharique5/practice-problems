from typing import List

class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        ans = 0
        for i in range(0, len(nums)+1):
            ans = ans ^ i
        for i in range(0, len(nums)):
            ans = ans ^ nums[i]
        return ans

sol = Solution()
inp = [int(x) for x in input().split(",")]
print(sol.missingNumber(inp))