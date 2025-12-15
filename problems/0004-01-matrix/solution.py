# 01 Matrix
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0004-01-matrix

class Solution:    
    def updateMatrix(self, mat: list[list[int]]) -> list[list[int]]:
        m = len(mat)
        if m < 1:
            return []
        n = len(mat[0])
        if n < 1:
            return []
        nr_mat = []
        max_distance = m + n
        # 1st pass: top->bottom, left->right
        for row in range(m):
            nr_mat.append([])
            for col in range(n):
                if mat[row][col] == 1:
                    top = nr_mat[row-1][col] if row > 0 else max_distance
                    left = nr_mat[row][col-1] if col > 0 else max_distance
                    nr_mat[row].append(min(top, left) + 1)
                else:
                    nr_mat[row].append(0)
        # 2nd pass: bottom->top, right->left
        for row in range(m-1, -1, -1):
            for col in range(n-1, -1, -1):
                if nr_mat[row][col] > 0:
                    bottom = nr_mat[row+1][col] if row < m-1 else max_distance
                    right = nr_mat[row][col+1] if col < n-1 else max_distance
                    nr_mat[row][col] = min(nr_mat[row][col], bottom+1, right+1)
        return nr_mat
