# Insert Interval
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0008-insert-interval

class Solution:
    def _overlaps(self, interval_a: list[int], interval_b: list[int]) -> bool:
        notOverlap = (
            interval_a[1] < interval_b[0] or
            interval_b[1] < interval_a[0]
        )
        return not notOverlap
    
    def _isBefore(self, interval_a: list[int], interval_b: list[int]) -> bool:
        """Returns True if interval_a is before interval_b."""
        return interval_a[1] < interval_b[0]
    
    def _newIntervalContainsAllIntervals(
            self, newInterval: list[int], intervals: list[list[int]]
        ) -> bool:
        return (
            newInterval[0] <= intervals[0][0] and
            newInterval[1] >= intervals[-1][1]
        )
    
    def insert(
            self, intervals: list[list[int]], newInterval: list[int]
        ) -> list[list[int]]:
        size = len(intervals)
        if size == 0:
            return [newInterval]
        if self._newIntervalContainsAllIntervals(newInterval, intervals):
            return [newInterval]
        if self._isBefore(newInterval, intervals[0]):
            return [newInterval] + intervals
        if self._isBefore(intervals[-1], newInterval):
            return intervals + [newInterval]
        newIntervals = []
        i = 0
        # Add all intervals before newInterval
        while (
            i < size and
            not self._overlaps(newInterval, intervals[i]) and
            self._isBefore(intervals[i], newInterval)
        ):
            newIntervals.append(intervals[i])
            i += 1
        # Merge all overlapping intervals with newInterval
        if i < size and self._overlaps(newInterval, intervals[i]):
            left = i
            right = size - 1
            last = i
            while left <= right:
                mid = (left + right) // 2
                if intervals[mid][0] <= newInterval[1]:
                    last = mid
                    left = mid + 1
                else:
                    right = mid - 1
            newIntervals.append([
                min(newInterval[0], intervals[i][0]),
                max(newInterval[1], intervals[last][1])
            ])
            i = last + 1
        # If no overlapping intervals were found, add newInterval
        elif i < size and not self._overlaps(newInterval, intervals[i]):
            newIntervals.append(newInterval)
        # Add all intervals after newInterval or merged intervals
        while i < size:
            newIntervals.append(intervals[i])
            i += 1
        return newIntervals
