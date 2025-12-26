---
id: 6
title: Time Based Key-Value Store
tags: [hash-table, string, binary-search]
source: LeetCode
---

# Time Based Key-Value Store

**📚 Source**: LeetCode

**🏷️ Tags**: [Hash Table](../../tags/hash-table.md), [String](../../tags/string.md), [Binary Search](../../tags/binary-search.md)

## 📋 Problem Statement

Design a time-based key-value data structure that can store multiple values for 
the same key at different time stamps and retrieve the key's value at a certain 
timestamp.

Implement the `TimeMap` class:

- `TimeMap()` Initializes the object of the data structure.
- `void set(String key, String value, int timestamp)` Stores the key `key` with 
the value `value` at the given time `timestamp`.
- `String get(String key, int timestamp)` Returns a value such that `set` was 
called previously, with `timestamp_prev <= timestamp`. If there are multiple 
such values, it returns the value associated with the largest `timestamp_prev`. 
If there are no values, it returns `""`.

**Constraints:**

- `1 <= key.length, value.length <= 100`
- `key` and `value` consist of lowercase English letters and digits.
- `1 <= timestamp <= 10^7`
- All the timestamps `timestamp` of `set` are strictly increasing.
- At most `2 * 10^5` calls will be made to `set` and `get`.

## 💡 Approach

Since the keys are strings, we use a hash map to store its data for efficient 
lookups. For the values and the timestamps of each key, we need a data 
structure that enable us to do efficient insertions and searchs (faster than 
linear time). We can achiev that by using a balanced binary tree (BBT) or a 
sorted array to enable the use of binary search. In this problem, we chose the 
array because its guaranted that the `timestamp_prev <= timestamp` so it will 
be naturaly sorted and we can insert new values faster than in a BBT. Using an 
array for both the values and the timestamps, we find the target index doing 
the binary search in the timestamp's array and access its value using direct 
indexing in the other array.

## ⚡ Complexity Analysis

In this analysis, `k = number of keys in TimeMap` and 
`n = number of values (value and timestamp) of a key`

- **Time Complexity**:

`set(key, value, timestamp)`: We check in amortized constant time if the key 
is present and initialize its value if needed. After that, the value and the 
timestamp are appended in constant time. So:

`O(1) + O(1) + O(1) = 3 * O(1) = O(1)`

`get(key, timestamp)`: We check in amortized constant time if the key is 
present. After that, we search for the index of the correct timestamp in 
logarithmic time. Then, we access and return the value in constant time. So:

`O(1) + O(log n) + O(1) = 2*O(1) + O(log n) = O(log n)`

Note: even though the hash calculation is linear with the length of the string 
that represents the key, we chose to consider it a constant due to its reduced 
maximum length (`100`) when compared to the maximum number of values a key can 
have (`2 * 10^5`).

- **Space Complexity**: O(k * n)

For each key, storing its values and timestamps requires linear space. So:

`O(k)*O(n) + O(k)*O(n) = 2 * O(k) * O(n) = O(k * n)`

## 💻 Code

See the [solution.py](solution.py) file.
