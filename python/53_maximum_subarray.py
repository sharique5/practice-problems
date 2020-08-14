from typing import List

class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        dp = [0] * len(nums)
        dp[0] = nums[0]
        ans = -9999999
        for i in range(1, len(nums)):
            dp[i] = max(dp[i-1], 0) + nums[i]
            if dp[i] > ans:
                ans = dp[i]
        return ans

sol = Solution()
nums = [int(x) for x in input().split(" ")]
print(sol.maxSubArray(nums))