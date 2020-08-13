from typing import List

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        temp_dict = {}
        for num in nums:
            if num in temp_dict:
                return True
            else:
                temp_dict[num] = num
        return False

sol = Solution()
nums = [int(x) for x in input().split(" ")]
print(sol.containsDuplicate(nums))