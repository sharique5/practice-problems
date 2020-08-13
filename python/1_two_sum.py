from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        temp_num = dict()
        for i in range(0, len(nums)):
            need = target - nums[i]
            if need in temp_num:
                return [temp_num[need], i]
            else:
                temp_num[nums[i]] = i
        return []

sol = Solution()
nums = [int(x) for x in input().split(', ')]
target = int(input())
print(sol.twoSum(nums, target))