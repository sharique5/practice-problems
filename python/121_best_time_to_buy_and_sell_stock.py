from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        ans = 0
        l, r = 0, 0
        while r < len(prices):
            if prices[l] <= prices[r]:
                ans = max(ans, prices[r]-prices[l])
                r += 1
            else:
                l += 1
        return ans


sol = Solution()
inp = [int(x) for x in input().split(",")]
print(sol.maxProfit(inp))