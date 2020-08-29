class Solution:
    def reverseBits(self, n: int) -> int:
        ans, bits = 0, 32
        for i in range(bits):
            ans |= (n&0x1)<<(bits-1-i)
            n >>= 1
        return ans