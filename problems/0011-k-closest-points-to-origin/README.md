---
id: 11
title: K Closest Points to Origin
tags: [array, math, geometry, sorting, heap, quickselect]
source: LeetCode
---

# K Closest Points to Origin

**📚 Source**: LeetCode

**🏷️ Tags**: [Array](../../tags/array.md), [Math](../../tags/math.md), [Geometry](../../tags/geometry.md), [Sorting](../../tags/sorting.md), [Heap](../../tags/heap.md), [Quickselect](../../tags/quickselect.md)

## 📋 Problem Statement

Given an array of `points` where `points[i] = [x_i, y_i]` represents a point on 
the **X-Y** plane and an integer `k`, return the `k` closest points to the 
origin `(0, 0)`.

The distance between two points on the **X-Y** plane is the Euclidean distance 
(i.e. $\sqrt{(x_1 - x_2)^2 + (y_1 - y_2)^2}$).

You may return the answer in **any order**. The answer is **guaranteed** to be 
**unique** (except for the order that it is in).

**Constraints:**

- `1 <= k <= points.length <= 10^4`
- `-10^4 <= x_i, y_i <= 10^4`

## 💡 Approach

A simple way to solve this problem is: build a new array containing the 
distance from each point to the origin in addition to its coordinates, sort 
this array by distance, and create the returning array using the first K 
points. It works, but it is inefficient.

Since we are only interested in the K closest points, we can reduce the extra 
space required by using a max heap to store only K elements. Now, we can sort 
the points while calculating their distance. Then, we build the return array 
using the elements in the heap.

Although the latter approach is already good, when K is not small, using 
Quickselect will be faster (linear runtime in most cases depending on the pivot 
choices). Note that we would still need to sort after Quickselect if the output 
had to be sorted.

Instead of using only one of these approaches, we will choose which one to use 
based on the value of K. If it is small, run the max heap approach. Otherwise, 
run the Quickselect. This way, we have a good balance asuming that the values 
of K are evenly distributed.

## ⚡ Complexity Analysis

In this analysis, `n == points.length`

- **Time Complexity**: O(n * log k)

When the algorithm branches to the max heap solution, we perform logarithmic 
heap operation for each point and then build the output in linear time. If the 
execution branches to Quickselect, in addition to running linearly most of the 
time, its execution time can sometimes (rarely) be quadratic. Furthermore, if 
in-place operations are not allowed, we must copy the input array before 
starting the Quickselect. So,

Max heap: `O(n * log k) + O(k) = O(n * log k)`

Quickselect: `O(n) + O(n) + O(k) = O(n) + O(k)`

Overall: `O(max(max heap, quickselect)) = O(n * log k)`

- **Space Complexity**: O(n)

If the execution follows the max heap branch, only extra linear space will be 
needed to store K elements in the heap. Otherwise, space will be needed to 
store a full copy of the input array. So,

Max heap: `O(k)`

Quickselect: `O(n)`

Overall: `O(max(max heap, quickselect)) = O(n)`

## 💻 Code

See the [solution.py](solution.py) file.
