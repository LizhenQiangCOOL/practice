
# 普通解法
class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        if needle == "":
            return 0
        
        haystack_len = len(haystack)
        needle_len = len(needle)
        # 确定好 i 最后一个的位置
        for i in range(haystack_len - needle_len + 1):
            for j in range(needle_len):
                if haystack[i+j] != needle[j]:
                    break
            else:
                return i
        return -1

# KMP
# haystack: "a b x a b c a b c a b y"
# needle: "a b c a b y"