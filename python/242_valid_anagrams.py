class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        unique_char = {}
        for ch in s:
            if ch in unique_char:
                unique_char[ch] += 1
            else:
                unique_char[ch] = 1
        for ch in t:
            if ch in unique_char:
                unique_char[ch] -= 1
            else:
                unique_char[ch] = 1
        for ch in unique_char:
            if unique_char[ch] != 0:
                return False
        return True

sol = Solution()

s = input()
t = input()
print(sol.isAnagram(s, t))