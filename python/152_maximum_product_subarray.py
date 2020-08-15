from typing import List

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        t_max = nums[0]
        t_min = nums[0]
        ans = nums[0]
        for i in range(1, len(nums)):
            if nums[i] < 0:
                temp = t_max
                t_max = t_min
                t_min = temp
            t_max = max( t_max * nums[i], nums[i])
            t_min = min( t_min * nums[i], nums[i])
            ans = max( ans, t_max)
        return ans

sol = Solution()
inp = [int(x) for x in input().split(",")]
print(sol.maxProduct(inp))