class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        res = 0
        substr = ''
        for ch in s:
            if ch not in substr:
                substr += ch
            else:
                ind = substr.find(ch)
                substr = substr[ind + 1:]
                substr += ch
            res = max(res, len(substr))
        return res

sol = Solution()
inp = input()
print(sol.lengthOfLongestSubstring(inp))