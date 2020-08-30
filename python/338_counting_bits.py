from typing import List

class Solution:
    def countBits(self, num: int) -> List[int]:
        ans = []
        for i in range(0, num + 1):
            temp = i
            cnt = 0
            while temp:
                if temp&1:
                    cnt += 1
                temp >>= 1
            ans.append(cnt)
        return ans

sol = Solution()
inp = int(input())
print(sol.countBits(inp))