class Solution:
    def getSum(self, a: int, b: int) -> int:
        mask = 0xffffffff
        s = (a ^ b) & mask
        c = a & b
        while c!=0:
            a = s
            b = (c << 1) & mask
            s = (a ^ b) & mask
            c = a & b
        if s & 0x80000000:
            s -= 0x100000000
        return s

sol = Solution()
a = int(input())
b = int(input())
print(sol.getSum(a, b))