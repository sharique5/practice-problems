from typing import List
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        if len(nums) == 0:
            return -1
        if len(nums) < 3:
            for i in range(0, len(nums)):
                if nums[i] == target:
                    return i
            return -1
        left, right = 0, len(nums) - 1

        if nums[left] == target:
            return left
        elif nums[right] == target:
            return right

        while True:      
            if right - left <= 1:
                return -1

            mid = (left + right) // 2
            if nums[mid] == target:
                return mid
            if nums[left] < target < nums[mid]:
                right = mid
            elif nums[mid] < target < nums[right]:
                left = mid
            elif nums[mid] > nums[left]:
                left = mid
            elif nums[right] > nums[mid]:
                right = mid
            else:
                return -1

sol = Solution()
inp = [int(x) for x in input().split(",")]
target = int(input())
print(sol.search(inp, target))