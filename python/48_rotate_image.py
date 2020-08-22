from typing import List
from sys import stdin

class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        rows = len(matrix)
        col = len(matrix[0])
        for i in range(rows):
            for j in range(i, col):
                temp = matrix[i][j]
                matrix[i][j] = matrix[j][i]
                matrix[j][i] = temp
        # print("transpose:: ", matrix)
        for i in range(rows):
            for j in range(0, col//2):
                temp = matrix[i][j]
                matrix[i][j] = matrix[i][col - j - 1]
                matrix[i][col - j - 1] = temp
        # print("rotated: ", matrix)        

sol = Solution()
inp = []
for line in stdin:
    if line == '':
        break
    temp = [int(x) for x in line.split(",")]
    inp.append(temp)
print("innitial::: ", inp)
sol.rotate(inp)
print("after func call::: ", inp)