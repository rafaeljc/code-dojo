---
id: 8
title: Insert Interval
tags: [array, binary-search]
source: LeetCode
---

# Insert Interval

**📚 Source**: LeetCode

**🏷️ Tags**: [Array](../../tags/array.md), [Binary Search](../../tags/binary-search.md)

## 📋 Problem Statement

You are given an array of non-overlapping intervals `intervals` where 
`intervals[i] = [start_i, end_i]` represent the start and the end of the `i^th` 
interval and `intervals` is sorted in ascending order by `start_i`. You are 
also given an interval `newInterval = [start, end]` that represents the start 
and end of another interval.

Insert `newInterval` into `intervals` such that `intervals` is still sorted in 
ascending order by `start` and `intervals` still does not have any overlapping 
intervals (merge overlapping intervals if necessary).

Return `intervals` after the insertion.

**Note** that you don't need to modify `intervals` in-place. You can make a new 
array and return it.

**Constraints:**

- `0 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i <= end_i <= 10^5`
- `intervals` is sorted by `start_i` in **ascending** order.
- `newInterval.length == 2`
- `0 <= start <= end <= 10^5`

## 💡 Approach

We can solve this problem in three phases: (1) append (in order) in the return 
array every interval that comes before and do not overlaps the new one; (2) 
merge every interval with the new one that overlaps it, and, after all merges, 
append the resulting interval in the return array; (3) finally, append 
(in order) all the remaining intervals. Since the array of intervals is sorted 
and there are no overlaps, we just iterate through the intervals array once 
because each step starts where the previous one ended thus avoiding unnecessary 
work.

Leveraging the fact that the intervals array is sorted, we could improve the 
merge phase by doing a binary search to find the last interval that overlaps 
the new one in logarithmic time and do a single merge without iterating over 
all overlapping intervals. Asymptotically, the worst case scenario for both 
approachs is a linear runtime but there are scenarios were the binary search is 
faster (lot of merges). Unfortunately, binary search introduces drawbacks like 
random memory access and it is not cache and branch prediction friendly. 
Depending on the distribution of the input values, the programming language, 
the data structure to store the intervals, and the hardware being used, the 
first approach will be faster in most cases.

Since we are going to implement the solution using Python and the invervals 
input is given as a List of Lists, we will use the binary search approach 
because in that scenario theres no guarantee that the values themselves will be 
stored sequentially in memory.

## ⚡ Complexity Analysis

In this analysis, `n == intervals.length` and 
`k == number of overlapping intervals`

- **Time Complexity**: O(n)

The first and the thirth phases are done in linear time. The merge phase is 
done in logarithmic time due to the use of a binary search. So,

`O(n - k) + O(log k) + O(n - k) = O(n - k) + O(log k)`

**Best scenario:** `k == n - 1` then `O(n - (n - 1)) + O(log (n - 1)) = O(1) + O(log n) = O(log n)`

**Worst scenario:** `k == 0` then `O(n - 0) + O(log 0) = O(n)`

- **Space Complexity**: O(1)

Since we only need a few variables to iterate through the intervals and do the 
binary search, a constant extra space is needed. So,

`O(1)`

## 💻 Code

See the [solution.py](solution.py) file.
