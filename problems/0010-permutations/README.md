---
id: 10
title: Permutations
tags: [array, backtracking]
source: LeetCode
---

# Permutations

**📚 Source**: LeetCode

**🏷️ Tags**: [Array](../../tags/array.md), [Backtracking](../../tags/backtracking.md)

## 📋 Problem Statement

Given an array `nums` of distinct integers, return all the possible 
permutations. You can return the answer in **any order**.

**Constraints:**

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- All the integers of `nums` are **unique**.

## 💡 Approach

Since the input array will have only unique elements, the returned array must 
contain all possible permutations, and it can be in any order, using a 
backtracking approach is better than using a lexicograph permutation one 
because it does not require a sorted array to work.

Starting with the full array, for each of the remaining elements, if it is in 
the permutation array already, continue. Otherwise, append it into the 
permutation array, mark it as visiting, and repeat this process on the 
remaining elements. Once it returns, undo the changes (mark the element as not 
visiting and pop it from the permutation array). Then, when the permutation 
array has the same length of the full array, copy it into the answer array.

Instead of creating a new array to build the permutation in it, we can use 
the original array (or a copy of it) to hold both the current state of the 
permutation and the remaining elements. Every time we "add" an element into the 
permutation, we reduce the number of remaining elements by one, and swap this 
element to a position that "belongs" to the permutation. Then, repeat that with 
the remaining ones. Once it returns, undo the change by just swapping it back 
to its original position.

Although this is a good approach for this scenario, using the Heap's 
algorithm will be faster, as it minimizes movement by doing less elements 
swaps. Therefore, let's use it.

## ⚡ Complexity Analysis

In this analysis, `n == nums.length`

- **Time Complexity**: O(n! * n)

Before starting the recursive calls, we make a copy of the input array because 
the algorithm do swaps in-place. Then, for each permutation built, we copy it 
into the output array. So,

`O(n) + O(n! * n) = O(n! * n)`

- **Space Complexity**: O(n)

In order to build the permutations, we need an array to store its current state 
as we swap elements. Because it is a recursive approach, the call stack will 
growth linearly with the length of the permutation. So, 

`O(n) + O(n) = 2 * O(n) = O(n)`

## 💻 Code

See the [solution.py](solution.py) file.
