from typing import List

class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        row_0 = col_0 = False
        m = len(matrix)
        n = len(matrix[0])
        for i in range(0, m):
            if matrix[i][0] == 0:
                col_0 = True
        for i in range(0, n):
            if matrix[0][i] == 0:
                row_0 = True
        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][j] == 0:
                    matrix[i][0] = 0
                    matrix[0][j] = 0
        for i in range(1, m):
            for j in range(1, n):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    matrix[i][j] = 0
        if row_0:
            for i in range(0, n):
                matrix[0][i] = 0
        if col_0:
            for i in range(0, m):
                matrix[i][0] = 0
        