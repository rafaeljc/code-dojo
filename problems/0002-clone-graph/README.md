---
id: 2
title: Clone Graph
tags: [Graph, Hash Table, Depth-First Search, Breadth-First Search]
source: LeetCode
---

# Clone Graph

**📚 Source**: LeetCode

**🏷️ Tags**: [Graph](../../tags/graph.md), [Hash Table](../../tags/hash-table.md), [Depth-First Search](../../tags/depth-first-search.md), [Breadth-First Search](../../tags/breadth-first-search.md)

## 📋 Problem Statement

Given a reference of a node in a connected undirected graph, return a deep 
copy (clone) of the graph.

Each node in the graph contains a value (`int`) and a list (`List[Node]`) of 
its neighbors.

```java
class Node {
    public int val;
    public List<Node> neighbors;
}
```

**Constraints:**

- The number of nodes in the graph is in the range `[0, 100]`.
- `1 <= Node.val <= 100`
- `Node.val` is unique for each node.
- There are no repeated edges and no self-loops in the graph.
- The graph is connected and all nodes can be visited starting from the given 
node.

## 💡 Approach

In order to make a deep copy, we must visit every node and edge of the graph 
and copy them. A common way to traverse a graph is using a Depth-First Search 
(DFS) or a Breadth-First Search (BFS) algorithm. The DFS implementation is 
more intuitive but it is recursive and the BFS is more complex to implement 
but it is iterative. For this problem, we choose the iterative approach 
because it is usually faster and uses less memory, and the code will not get 
too complex to understand. To keep track of visited nodes and to store the 
nodes copies, we could use a Hash Table because of it is amortized O(1) search 
and insert operations. Although, in this problem, the node identifier (val) 
set are not sparse and we know beforehand it is maximum size (100). So, we can 
use an array to store the copies and to do faster lockups (O(1) direct 
indexing).

## ⚡ Complexity Analysis

In all analysis, `V = number of nodes` and `E = number of edges`.

- **Time Complexity**: O(V + E)

At the start, we initiate in constant time (same length regardless of the 
graph) the array that will store the copies and serve as a visit check. After 
that, the BFS main loop will iterate through all nodes and "visit" their 
neighbors edges. In total, we had visited every edge. So:

`O(1) + O(V)*O(neighbors E) = O(V + E)`

- **Space Complexity**: O(V)

Using a queue to determine the order the nodes are visited requires `V` extra 
space. Besides we are making a deep copy of the graph, because those copies 
are "the returned values" we do not count them as an extra space. So:

`O(V)`

## 💻 Code

See the [solution.py](solution.py) file.
