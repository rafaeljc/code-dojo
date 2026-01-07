# Minimum Window Substring
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0007-minimum-window-substring

import array as arr

class Solution:
    FIRST_CHAR_UNICODE = ord('A')
    LAST_CHAR_UNICODE = ord('z')
    ARR_SIZE = LAST_CHAR_UNICODE - FIRST_CHAR_UNICODE + 1
    ARR_OFFSET = FIRST_CHAR_UNICODE

    def _indexOf(self, c: str) -> int:
        return ord(c) - Solution.ARR_OFFSET
    
    def minWindow(self, s: str, t: str) -> str:
        m = len(s)
        n = len(t)
        if m < 1 or n < 1:
            return ""
        if n > m:
            return ""
        t_char_missing = arr.array('l', [0] * Solution.ARR_SIZE)
        for letter in t:
            t_char_missing[self._indexOf(letter)] += 1
        # sliding window: [w_begin, w_end)
        t_char_total_missing = n
        w_begin = 0
        w_end = 0
        min_w_begin = 0
        min_w_size = m + 1 # impossible size
        while w_end < m:
            idx = self._indexOf(s[w_end])
            if t_char_missing[idx] > 0:
                t_char_total_missing -= 1
            t_char_missing[idx] -= 1
            w_end += 1
            # if all chars from t are in the current window
            # shrink the window from the begining until we have the minimum
            # possible window that still contains all chars from t
            while t_char_total_missing == 0:
                idx = self._indexOf(s[w_begin])
                # the minimum possible window that contains all chars from t
                # is the window where removing the next char would make it
                # invalid
                if t_char_missing[idx] == 0:
                    # update minimum window if needed
                    w_size = w_end - w_begin
                    if w_size < min_w_size:
                        min_w_begin = w_begin
                        min_w_size = w_size
                    t_char_total_missing += 1
                t_char_missing[idx] += 1
                w_begin += 1
        # min_w_size will be <= m only if we found a valid window
        if min_w_size <= m:
            return s[min_w_begin : min_w_begin + min_w_size]
        return ""
