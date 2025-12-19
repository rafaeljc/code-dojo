---
id: 5
title: Number of Islands
tags: [array, depth-first-search, breadth-first-search, union-find, matrix]
source: LeetCode
---

# Number of Islands

**📚 Source**: LeetCode

**🏷️ Tags**: [Array](../../tags/array.md), [Depth First Search](../../tags/depth-first-search.md), [Breadth First Search](../../tags/breadth-first-search.md), [Union Find](../../tags/union-find.md), [Matrix](../../tags/matrix.md)

## 📋 Problem Statement

Given an `m x n` 2D binary grid `grid` which represents a map of `1`s (land) 
and `0`s (water), return the number of islands.

An **island** is surrounded by water and is formed by connecting adjacent 
lands horizontally or vertically. You may assume all four edges of the grid 
are all surrounded by water.

**Constraints:**

- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 300`
- `grid[i][j]` is `0` or `1`.

## 💡 Approach

**Option 1: Graph Traversal**

In order to count the number of islands, we have to group up all the one-cells 
that are connected and the number of islands will be the number of connected 
components. That can be done by doing graph traversals (DFS/BFS) starting from 
each non-visited one-cell. In each traversal, we visit every land that is part 
of the same island. Implementing the DFS is simple and intuitive, but it is 
recursive and brings problems like stack overflow risk and can be harder to 
debug. On the other hand, the BFS is iterative (usually faster than recursive) 
and, in this problem, the code will not get hard to understand.

**Option 2: Union Find**

Instead of running graph traversals, we can simply traverse the grid and 
group adjacent one-cells into the same set. In order to group the cells 
efficiently, we have to use the disjoint sets data structure with path 
compression and union by rank. At the end of the grid traversal, the number of 
islands will be the number of roots of the disjoint sets.

**Chosen Approach:**

Besides the two options are asymptotically equal in both runtime and space, 
the constants of BFS are smaller than those of Union Find (the find operation 
have amortized runtime complexity and, have to maintain the root and rank for 
each cell all time). In addition, if the chosen programming language does not 
offer any support for Union Find, we will also have to implent this non-simple 
data structure from scrach.

## ⚡ Complexity Analysis

- **Time Complexity**: O(m * n)

At start, we initialize in linear time a grid to track the visited cells. 
After that, we traverse the graph in linear time to count the number of 
islands. So:

`O(m * n) + O(m * n) = 2 * O(m * n) = O(m * n)`

- **Space Complexity**: O(m * n)

The auxiliary grid that helps us track visited cells and the queue that 
determines the order in which cells are visited require linear space. So:

`O(m * n) + O(m * n) = 2 * O(m * n) = O(m * n)`

## 💻 Code

See the [solution.py](solution.py) file.
