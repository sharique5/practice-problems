class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        res = 0
        freq = [0 for i in range(26)]
        l, frequent_ch = 0, 0
        for i in range(len(s)):
            ind = ord(s[i]) - ord('A')
            freq[ind] += 1
            frequent_ch = max(frequent_ch, freq[ind])

            while i-l-frequent_ch+1 > k:
                freq[ord(s[l]) - ord('A')] -= 1
                l += 1
            
            res = max(res, i-l+1)
        return res

sol = Solution()
inp = input()
k = int(input())
print(sol.characterReplacement(inp, k))