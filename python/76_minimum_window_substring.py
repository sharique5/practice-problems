class Solution:
    def isTinS(self, s_freq: {}, t_freq: {}) -> bool:
        print(s_freq, t_freq)
        for ch in t_freq:
            if ch not in s_freq:
                return False
            if t_freq[ch] > s_freq[ch]:
                return False    
        return True
    
    def minWindow(self, s: str, t: str) -> str:
        if len(t) > len(s):
            return ""
        if len(t) == 1:
            if t in s:
                return t
            return ""
        freq_t = {}
        for ch in t:
            if ch in freq_t:
                freq_t[ch] += 1
            else:
                freq_t[ch] = 1
        freq_s = {}
        l, r = 0, 1
        freq_s[s[l]] = 1
        if s[r] in freq_s:
            freq_s[s[r]] += 1
        else:
            freq_s[s[r]] = 1
        res = s
        window_present = False
        while l <= r and r < len(s):
            print("lookinng in:: ", s[l:r+1])
            if self.isTinS(freq_s, freq_t):
                window_present = True
                temp = s[l:r+1]
                if len(temp) < len(res):
                    res = temp
                freq_s[s[l]] -= 1
                l += 1
            else:
                r += 1
                if r < len(s):
                    if s[r] in freq_s:
                        freq_s[s[r]] += 1
                    else:
                        freq_s[s[r]] = 1
        if window_present:
            return res
        return ""

sol = Solution()
inps = input()
inpt = input()
print(sol.minWindow(inps, inpt)) 