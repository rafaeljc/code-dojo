# K Closest Points to Origin
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0011-k-closest-points-to-origin

# Requires Python 3.14+ for heapq.heappush_max and heapq.heapreplace_max
import heapq

class Solution:
    def _heapK(
        self, points: list[list[int]], k: int, origin: tuple[int, int] = (0, 0)
    ) -> list[list[int]]:
        ox, oy = origin
        max_heap = []
        for x, y in points:
            dx = x - ox
            dy = y - oy
            proximity = dx*dx + dy*dy
            if len(max_heap) < k:
                heapq.heappush_max(max_heap, (proximity, x, y))
            elif proximity < max_heap[0][0]:
                heapq.heapreplace_max(max_heap, (proximity, x, y))
        # Iterating over the array that represents the heap instead of popping 
        # elements because the order of the output does not matter.
        return [[x, y] for _, x, y in max_heap]
    
    def _quickselect(
        self, points: list[list[int]], k: int, origin: tuple[int, int] = (0, 0)
    ) -> list[list[int]]:
        ox, oy = origin
        left = 0
        right = len(points) - 1
        while True:
            # Choose pivot as the middle element
            pivot_i = (left + right) >> 1 # (left + right) // 2
            px, py = points[pivot_i]
            dx = px - ox
            dy = py - oy
            # Do not need to calculate the square root since we are only 
            # interested in knowing if a point is closer than another.
            pivot_prox = dx*dx + dy*dy
            points[pivot_i], points[right] = points[right], points[pivot_i]
            insert_i = left
            for i in range(left, right):
                x, y = points[i]
                dx = x - ox
                dy = y - oy
                if dx*dx + dy*dy < pivot_prox:
                    points[insert_i], points[i] = points[i], points[insert_i]
                    insert_i += 1
            points[insert_i], points[right] = points[right], points[insert_i]
            if insert_i == k:
                break
            elif insert_i < k:
                left = insert_i + 1
            else: # if insert_i > k:
                right = insert_i - 1
        # Straight return the first k elements since their order does not
        # matter.
        return points[:k]
    
    def kClosest(self, points: list[list[int]], k: int) -> list[list[int]]:
        if k < 1:
            return []
        size = len(points)
        if size <= k:
            return [point[:] for point in points]
        # Heuristically, when k is 8 times smaller than the size of the list,
        # the heap approach is expected to be faster.
        if k * 8 <= size:
            return self._heapK(points, k)
        return self._quickselect([point[:] for point in points], k)
