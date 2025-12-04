# Course Schedule
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0003-course-schedule

import queue

class Solution:
    def _getAdjListAndInDegrees(
        self, numCourses: int, prerequisites: list[list[int]]
    ) -> tuple[list[list[int]], list[int]]:
        adj_list = [[] for _ in range(numCourses)]
        in_degrees = [0 for _ in range(numCourses)]
        for course, prereq in prerequisites:
            adj_list[prereq].append(course)
            in_degrees[course] += 1
        return adj_list, in_degrees

    def _isDAG(self, adj_list: list[list[int]], in_degrees: list[int]) -> bool:
        q = queue.Queue()
        for node, in_degree in enumerate(in_degrees):
            if in_degree == 0:
                q.put(node)
        remaining_nodes = len(adj_list)
        while not q.empty():
            node = q.get()
            remaining_nodes -= 1
            for neighbor in adj_list[node]:
                in_degrees[neighbor] -= 1
                if in_degrees[neighbor] == 0:
                    q.put(neighbor)
        return remaining_nodes == 0
    
    def canFinish(
        self, numCourses: int, prerequisites: list[list[int]]
    ) -> bool:
        if numCourses < 1:
            return False
        if len(prerequisites) == 0:
            return True
        adj_list, in_degrees = self._getAdjListAndInDegrees(
            numCourses, prerequisites
        )
        return self._isDAG(adj_list, in_degrees)
