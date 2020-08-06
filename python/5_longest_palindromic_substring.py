class Solution:
    def longestPalindrome(self, s: str) -> str:
        st = '$#'
        for ch in s:
            st += ch
            st += '#'
        st += '@'
        # print('updated st: ', st)
        p = [0 for i in range(len(st))]
        # print('p::: ', p)
        c, r = 0, 0
        for i in range(1, len(st)-1):
            if i < r:
                p[i] = min(r-i, p[2*c - i])
            while st[i + (1 + p[i])] == st[i - (1 + p[i])]:
                p[i] += 1
            if i + p[i] > r:
                c = i
                r = i + p[i]
        res = ''
        r, c = max(p), p.index(max(p))
        if st[c] != '#' and st[c] != '$' and st[c] != '@':
            res = st[c]
        for i in range(c+1, c+r+1):
            if st[i] != '#':
                res = st[i] + res + st[i]
        return res

sol = Solution()
inp = input()
print(sol.longestPalindrome(inp))