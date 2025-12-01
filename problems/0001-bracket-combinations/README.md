---
id: 1
title: Bracket Combinations
tags: [combinatorics, backtracking]
source: Coderbyte
---

# Bracket Combinations

**📚 Source**: Coderbyte

**🏷️ Tags**: [Combinatorics](../../tags/combinatorics.md), [Backtracking](../../tags/backtracking.md)

## 📋 Problem Statement

Have the function `BracketCombinations(num)` read `num` which will be an 
integer greater than or equal to zero, and return the number of valid 
combinations that can be formed with `num` pairs of parentheses. For example, 
if the input is 3, then the possible combinations of 3 pairs of parentheses, 
namely: ()()(), are ()()(), ()(()), (())(), ((())) and (()()). There are 5 
total combinations when the input is 3, so your program should return 5.

## 💡 Approach 1: Generate all combinations

The first thought that come is: lets generate all combinations by permutating
the brackets and count the valid ones. But, we do not need to generate all of 
them, specialy because there is repeated elements. And if we could check, 
while constructing the permutation, if the actual state can lead to a valid 
bracket combination in order to skip that path and try another? It is possible 
by using the backtracking technique, keep a counter of `(`, and when we get an 
`)` just check if the `(` counter is greater than zero. So, every time we 
reach the end of a path, we count as a valid combination.

## ⚡ Complexity Analysis

In all analysis `n = 2 * num`

- **Time Complexity**: O(n!)

At start, we create the initial state with all the brackets in linear time. 
After that, while generating all the permutations in factorial time (even 
skipping the invalid paths), check if a permutation is valid in constant time. 
So:

`O(n) + O(n!)*O(1) = O(n) + O(n!) = O(n!)`

- **Space Complexity**: O(n)

At start, we use `n` space to store the brackets initial state, and a single 
space for `(` and valid combinations counters. During the backtrack, the 
maximum size of the call stack will be `n` (when we reach a valid path) if we 
"create" the array containing the remaining elements to be passed to the next 
recursive call with constant space (we can use the initial array and every 
time we need to "remove" an element, just swap him with the element in the 
last position, pass the "new" size to the next recursive call, and swap back 
before trying the next element). So:

`O(n) + 2*O(1) + O(n)*O(1) = O(n) + O(1) + O(n) = O(n)`

## 💡 Approach 2: Using combinatorics

Looking for the results of `num` ranging from 1 to 10 (`1, 2, 5, 14, 42, 132, 429, 1430, 4862, 16796`) we can think that the result could be following some 
sort of sequence. In fact, it is. The result is a [Catalan Number](https://en.wikipedia.org/wiki/Catalan_number) and it is value is determined by the 
following formula:

$$
C_n = \frac{(2n)!}{(n+1)!\,n!} \quad 
\text{for } n \ge 0
$$

Although now we can get our result by just calculating this formula (which 
is faster than the previous approach), depending on the programming language 
you are going to implementing it, that will lead to numeric problems because 
factorial grows extremely fast and the number quickly will not be able to be 
represented in memory. Using the alternative formula bellow, we maintain the 
performance and reduce the factorial problem (that formula grows slower than the factorial one).

$$
C_0 = 1 \quad and \quad
C_n = \frac{2\,(2n - 1)}{n + 1} \, C_{n - 1} \quad 
\text{for } n > 0
$$

## ⚡ Complexity Analysis

- **Time Complexity**: O(n)

In order to compute the formula, we only have to do some constant operations 
(sum, subtraction, multiplication, and division) `n` times. So:

`O(n) * 5*O(1) = O(n)` 

- **Space Complexity**: O(1)

Implenting it iteratively only requires a constant amount of memory to help the
formula calculation.

## 💻 Code

See the [solution.cpp](solution.cpp) file.
