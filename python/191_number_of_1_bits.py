class Solution:
    def hammingWeight(self, n: int) -> int:
        count = 0
        while n:
            count += (n % 2)
            n = n >> 1
        return count

sol = Solution()
inp = int(input())
print("inp:: ", inp)
print(sol.hammingWeight(inp))