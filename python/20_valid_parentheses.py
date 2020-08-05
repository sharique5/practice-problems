class Stack:
    def __init__(self):
        self.stack = []
    
    def isEmpty(self):
        return self.size() == 0
    
    def push(self, item):
        self.stack.append(item)
    
    def pop(self):
        return self.stack.pop()
    
    def peek(self):
        return self.stack[self.size() - 1]
    
    def size(self):
        return len(self.stack)

class Solution:
    def isValid(self, s: str) -> bool:
        st = Stack()
        opening_braces = ['(', '{', '[']
        closing_braces = [')', '}', ']']
        for ch in s:
            if ch in opening_braces:
                st.push(ch)
            elif ch in closing_braces:
                if st.isEmpty():
                    return False
                top_ch = st.peek()
                if (top_ch == '(' and ch == ')') or (top_ch == '{' and ch == '}') or (top_ch == '[' and ch == ']'):
                    st.pop()
                else:
                    return False
        return st.isEmpty()

sol = Solution()
inp = input()
print(sol.isValid(inp))