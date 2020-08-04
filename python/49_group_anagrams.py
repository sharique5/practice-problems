from typing import List

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        group_anagrams = {}
        for strings in strs:
            key_val = ''.join(sorted(strings))
            # print(key_val)
            if key_val in group_anagrams:
                group_anagrams[key_val].insert(len(group_anagrams[key_val]), strings)
            else:
                group_anagrams[key_val] = [strings]
        res = []
        for keys in group_anagrams:
            res.insert(len(res), group_anagrams[keys])
        return res

sol = Solution()
inp = input().split(", ")
print(sol.groupAnagrams(inp))