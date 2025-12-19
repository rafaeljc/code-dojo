# Number of Islands
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0005-number-of-islands

from collections import deque

class Solution:
    _DIFFS = ((1, 0), (-1, 0), (0, 1), (0, -1))

    def _isAValidCell(self, row: int, col: int, grid: list[list[str]]) -> bool:
        m = len(grid)
        n = len(grid[0])
        if row < 0 or row >= m:
            return False
        if col < 0 or col >= n:
            return False
        return True
    
    def _runBFS(
        self,
        start_row: int,
        start_col: int,
        grid: list[list[str]],
        visited: list[list[bool]]
    ) -> None:
        q = deque([])
        q.append((start_row, start_col))
        visited[start_row][start_col] = True
        while q:
            r, c = q.popleft()
            for dr, dc in Solution._DIFFS:
                nb_r = r + dr
                nb_c = c + dc
                valid_cell = self._isAValidCell(nb_r, nb_c, grid)
                if valid_cell and not visited[nb_r][nb_c]:
                    q.append((nb_r, nb_c))
                    visited[nb_r][nb_c] = True
        return
    
    def numIslands(self, grid: list[list[str]]) -> int:
        m = len(grid)
        if m < 1:
            return 0
        n = len(grid[0])
        if n < 1:
            return 0
        visited = [
            [
                False if cell == "1" else True
                for cell in row
            ]
            for row in grid
        ]
        num_islands = 0
        for row in range(m):
            for col in range(n):
                if not visited[row][col]:
                    self._runBFS(row, col, grid, visited)
                    num_islands += 1
        return num_islands
