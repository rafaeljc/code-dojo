---
id: 7
title: Minimum Window Substring
tags: [hash-table, string, sliding-window]
source: LeetCode
---

# Minimum Window Substring

**📚 Source**: LeetCode

**🏷️ Tags**: [Hash Table](../../tags/hash-table.md), [String](../../tags/string.md), [Sliding Window](../../tags/sliding-window.md)

## 📋 Problem Statement

Given two strings `s` and `t` of lengths `m` and `n` respectively, return the 
minimum window substring of `s` such that every character in `t` (including 
duplicates) is included in the window. If there is no such substring, return 
the empty string `""`.

It is guaranted that the answer is unique.

**Constraints:**

- `m == s.length`
- `n == t.length`
- `1 <= m, n <= 10^5`
- `s` and `t` consist of uppercase and lowercase English letters.

## 💡 Approach

First, we count the frequency of characters in `t` to be able to check if a 
window of `s` have all characters of `t`. Now, we could create windows of `s` 
(with sizes starting from `n` until `m`) and slide them while checking if any 
have all characters of `t`. If one of them (the first) have all characters, 
return that window. After trying every window size, if it does not have 
returned yet, return an empty string. But, that is inefficient because we are 
visiting every char multiple times since we are trying every window size.

Instead, we can grow the window by moving its end to the right until we have 
all characters of `t` in it. Then, we shrink it by moving its begin to the 
right until we have the minimum possible number of `t` characters in it and 
keeping the window valid. Repeat that until we reach the end of `s` while 
tracking the minimum window found. Every time we move the window, we do not 
need to count its characters frequency from scrach. We just update it as we add 
or remove the letters. In order to do that efficiently, instead of using a hash 
map, we will use an array for faster access through direct indexing because the 
number of possible characters is constant and small (English letters). 

## ⚡ Complexity Analysis

- **Time Complexity**: O(n + m)

At the start, we count the frequency of characters of `t` in linear time. After 
that, we slide the window in linear time because the window's pointers had only 
moved to the right. So,

`O(n) + O(m) = O(n + m)`

- **Space Complexity**: O(1)

In order to slide the window while keeping track of the frequency of characters 
in it, we need a fixed size array to store them. Also, we need some variables 
to hold data for the sliding window. So,

`O(1) + O(1) = O(1)`

## 💻 Code

See the [solution.py](solution.py) file.
