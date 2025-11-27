# Clone Graph
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0002-clone-graph

import queue
from typing import Optional

class Node:
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

class Solution:
    _MIN_NODE_VALUE = 1
    _MAX_NODE_VALUE = 100
    _MAX_NUM_NODES = _MAX_NODE_VALUE - _MIN_NODE_VALUE + 1

    def _getIndex(self, node: Node) -> int:
        return node.val - Solution._MIN_NODE_VALUE

    def cloneGraph(self, node: Optional["Node"]) -> Optional["Node"]:
        if not node:
            return None
        copies: list[Optional[Node]] = [
            None for _ in range(Solution._MAX_NUM_NODES)
        ]
        q = queue.Queue()
        q.put(node)
        node_idx = self._getIndex(node)
        copies[node_idx] = Node(node.val)
        while not q.empty():
            n = q.get()
            n_copy = copies[self._getIndex(n)]
            assert n_copy is not None  # to satisfy type checker
            for neighbor in n.neighbors:
                neighbor_idx = self._getIndex(neighbor)
                neighbor_copy = copies[neighbor_idx]
                if not neighbor_copy:
                    q.put(neighbor)
                    neighbor_copy = Node(neighbor.val)
                    copies[neighbor_idx] = neighbor_copy
                n_copy.neighbors.append(neighbor_copy)
        return copies[node_idx]
