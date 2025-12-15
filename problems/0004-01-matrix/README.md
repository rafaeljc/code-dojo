---
id: 4
title: 01 Matrix
tags: [matrix, dynamic-programming, breadth-first-search]
source: LeetCode
---

# 01 Matrix

**📚 Source**: LeetCode

**🏷️ Tags**: [Matrix](../../tags/matrix.md), [Dynamic Programming](../../tags/dynamic-programming.md), [Breadth First Search](../../tags/breadth-first-search.md)

## 📋 Problem Statement

Given an `m x n` binary matrix `mat`, return the distance of the nearest `0` 
for each cell.

The distance between two cells sharing a common edge is `1`.

**Constraints:**

- `m == mat.length`
- `n == mat[i].length`
- `1 <= m, n <= 10^4`
- `1 <= m * n <= 10^4`
- `mat[i][j]` is either `0` or `1`.
- There is at least one `0` in `mat`.

## 💡 Approach

**Option 1: BFS**

Even though the input is a matrix, we can think of it as a graph where the 
cells are the nodes and the ones that have an common edge are neighbors with a 
distance of `1`. In this way, we can traverse the matrix using a graph 
algorithm like BFS. With BFS is guaranted that if we find a path between a 
zero-cell and a one-cell it is the shortest one. It is true because our 
"graph" have all edges with the same weight. At the start, we put in the queue 
all zero-cells. While the queue is not empty, just pop the first element and 
check their neighbors. If a neighbor was not visited yet, set it's distance as 
the actual node's distance to a zero-cell plus one and add that neighbor to 
the queue. After processing every cell, we will have all the distances from a 
one-cell to its nearest zero-cell.

**Option 2: Dynamic Programming (DP)**

If all neighbors of a given cell already had the distance to their nearest 
zero-cell, the distance of that cell will be the smallest distance of it's 
neighbors plus one. The issue: the neighbors of each one-cell must known their 
distance to the nearest zero-cell beforehand. So, for every one-cell we have 
to run some search algorithm to find it, which will lead us to a terrible 
runtime. Instead, we traverse the matrix cell by cell 2 times. On the first 
pass, we traverse from top to bottom and from left to right. In this way, we 
get the distances from the top and the left neighbors. On the second pass, we 
traverse from bottom to top and from right to left. Then, we get the distances 
from the bottom and the right neighbors and finally chose the shortest one.

**Chosen Approach:**

Although both are simple to implement and have the same runtime, the DP 
approach is better because it uses less memory (does not require additional 
structure such as a queue in BFS).

## ⚡ Complexity Analysis

- **Time Complexity**: O(M * N)

Each matrix traversal is done in linear time and, for each cell, we check 2 of 
its neighbors and chose the shortest distance in constant time. So:

`2 * O(M * N) = O(M * N)`

- **Space Complexity**: O(1)

Instead of storing the partial results from the first pass in a temporary 
matrix, we store them directly in the output matrix. And, no auxiliary space 
is required. So:

`O(1)`

## 💻 Code

See the [solution.py](solution.py) file.
