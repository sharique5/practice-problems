from typing import List

class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        res = [0] * len(nums)
        l2r = []
        r2l = []
        # l2r
        temp = 1
        for num in nums:
            temp = temp * num
            l2r.append(temp)
        # r2l
        temp = 1
        for i in range(len(nums) - 1, -1, -1):
            temp = temp * nums[i]
            r2l.insert(0, temp)
        # print("l2r::: ", l2r)
        # print("r2l::: ", r2l)
        for i in range(len(nums)):
            # print(i)
            if i == 0:
                res[i] = r2l[1]
            elif i == len(nums) - 1:
                res[i] = l2r[i - 1]
            else:
                res[i] = l2r[i-1] * r2l[i+1]
        return res


sol = Solution()
inp = [int(x) for x in input().split(",")]
print(sol.productExceptSelf(inp))