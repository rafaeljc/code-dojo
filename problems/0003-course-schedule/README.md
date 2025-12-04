---
id: 3
title: Course Schedule
tags: [depth-first-search, breadth-first-search, graph, topological-sort]
source: LeetCode
---

# Course Schedule

**📚 Source**: LeetCode

**🏷️ Tags**: [Depth First Search](../../tags/depth-first-search.md), [Breadth First Search](../../tags/breadth-first-search.md), [Graph](../../tags/graph.md), [Topological Sort](../../tags/topological-sort.md)

## 📋 Problem Statement

There are a total of `numCourses` courses you have to take, labeled from `0` 
to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [a_i, b_i]` indicates that you **must** take course `b_i` 
first if you want to take course `a_i`.

* For example, the pair `[0, 1]`, indicates that to take course `0` you have 
to first take course `1`.

Return `true` if you can finish all courses. Otherwise, return `false`.

**Constraints:**

* `1 <= numCourses <= 2000`
* `0 <= prerequisites.length <= 5000`
* `prerequisites[i].length == 2`
* `0 <= a_i, b_i < numCourses`
* All the pairs `prerequisites[i]` are **unique**

## 💡 Approach

In order to know if we can finish all courses, we have to resolve the courses 
dependencies and if there is a cyclical dependency, we cannot finish all 
courses. The approach to do that is by doing a topological sort. That can be 
achieved by using a DFS or Kahn's algorithm. Both will give us simple code, 
but the Kahn's algorithm will be better because it is iterative (faster than 
recursive). And, before we solve the problem, we have to convert the edge list 
graph representation into a adjacency list representation because it is faster 
to access the neighbors of a node (the dependencies of a course). For the 
adjacency list, we will use an array of lists because the node identifiers set 
consists of integers and is not sparse. That enable us to do direct indexing 
for fast lockups and use memory efficiently.

## ⚡ Complexity Analysis

In this analysis, `V = number of nodes` and `E = number of edges`

- **Time Complexity**: O(V + E)

At start, we build the adjacency and the in degree lists in linear time 
(`V + E`). Then, we check if the graph have a cycle in linear time (`V + E`) 
because we setup the starting queue in `V` time and process all edges in `E` 
time. So:

`2*O(V + E) + O(V + E) = 3 * O(V + E) = O(V + E)`

- **Space Complexity**: O(V + E)

The adjacency list requires `V + E` space since it represents the graph. The in 
degree list and the queue require `V` space (each) because they only track 
node's data. So:

`O(V + E) + 2*O(V) = O(V + E)`

## 💻 Code

See the [solution.py](solution.py) file.
