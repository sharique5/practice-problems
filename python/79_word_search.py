from typing import List
from sys import stdin

class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        rows = len(board)
        cols = len(board[0])
        for i in range(rows):
            for j in range(cols):
                if word[0] == board[i][j] and self.dfs(board, i, j, word[1:]):
                    return True
        return False

    def dfs(self, board, x, y, word) -> bool:
        if not word:
            return True
        
        temp = board[x][y]
        board[x][y] = '*'
        rows = len(board)
        cols = len(board[0])

        if 0 <= x - 1 < rows and 0 <= y < cols:
            if word[0] == board[x - 1][y] and self.dfs(board, x - 1, y, word[1:]):
                return True
        
        if 0 <= x < rows and 0 <= y + 1 < cols:
            if word[0] == board[x][y + 1] and self.dfs(board, x, y + 1, word[1:]):
                return True
        
        if 0 <= x + 1 < rows and 0 <= y < cols:
            if word[0] == board[x + 1][y] and self.dfs(board, x + 1, y, word[1:]):
                return True

        if 0 <= x < rows and 0 <= y - 1 < cols:
            if word[0] == board[x][y - 1] and self.dfs(board, x, y - 1, word[1:]):
                return True

        board[x][y] = temp
        return False

sol = Solution()
word = input()
inp = []
for line in stdin:
    temp = [x.rstrip() for x in line.split(',')]
    inp.append(temp)
print(sol.exist(inp, word))