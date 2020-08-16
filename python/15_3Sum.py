from typing import List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        result = []
        n = len(nums)
        for i in range(0, n):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            st = i + 1
            end = n - 1
            while st < end:
                if nums[i] + nums[st] + nums[end] == 0:
                    result.append([nums[i], nums[st], nums[end]])
                    st += 1
                    while st < end and nums[st] == nums[st-1]:
                        st += 1
                    end -= 1
                    while st < end and nums[end] == nums[end+1]:
                        end -= 1
                elif nums[i] + nums[st] + nums[end] < 0:
                    st += 1
                    while st < end and nums[st] == nums[st-1]:
                        st += 1
                else:
                    end -= 1
                    while st < end and nums[end] == nums[end+1]:
                        end -= 1
        return result

sol = Solution()
inp = [int(x) for x in input().split(", ")]
print(sol.threeSum(inp))