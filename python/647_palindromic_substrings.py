class Solution:
    def countSubstrings(self, s: str) -> int:
        cnt = 0
        size = len(s)
        for i in range(size):
            for j in range(i+1, size+1):
                st = s[i:j]
                if st == st[::-1]:
                    cnt += 1
        return cnt

sol = Solution()
inp = input()
print(sol.countSubstrings(inp))