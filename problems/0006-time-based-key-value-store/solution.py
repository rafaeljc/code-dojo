# Time Based Key-Value Store
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0006-time-based-key-value-store

import array as arr
import bisect as bs

class TimeMap:
    def __init__(self):
        self._data = {}

    def set(self, key: str, value: str, timestamp: int) -> None:
        if key == "" or value == "":
            return
        if timestamp < 1:
            return
        d = self._data.get(key, None)
        if d is None:
            d = {
                "values:": [],
                # using array is more memory efficient than list for integers
                # since array.array stores the data itself in a contiguous 
                # block of memory, avoiding unnecessary jumps to access each 
                # element and it is more cache friendly.
                "timestamps": arr.array("I", []),
            }
            self._data[key] = d
        d["values"].append(value)
        d["timestamps"].append(timestamp)

    def get(self, key: str, timestamp: int) -> str:
        d = self._data.get(key, None)
        if d is None:
            return ""
        # binary search to find the position of the value that have a timestamp 
        # less than or equal to the input's timestamp
        i = bs.bisect_right(d["timestamps"], timestamp)
        if i > 0:
            return d["values"][i - 1]
        return ""
 