---
id: 9
title: Evaluate Reverse Polish Notation
tags: [array, math, stack]
source: LeetCode
---

# Evaluate Reverse Polish Notation

**📚 Source**: LeetCode

**🏷️ Tags**: [Array](../../tags/array.md), [Math](../../tags/math.md), [Stack](../../tags/stack.md)

## 📋 Problem Statement

You are given an array of strings `tokens` that represents an arithmetic 
expression in a Reverse Polish Notation.

Evaluate the expression. Return an integer that represents the value of the 
expression.

**Note** that:

- The valid operators are `"+"`, `"-"`, `"*"`, and `"/"`.
- Each operand may be an integer or another expression.
- The division between two integers always **truncates toward zero**.
- There will not be any division by zero.
- The input represents a valid arithmetic expression in a reverse polish 
notation.
- The answer and all the intermediate calculations can be represented in a 
**32-bit** integer.

**Constraints:**

- `1 <= tokens.length <= 10^4`
- `tokens[i]` is either an operator (`"+"`, `"-"`, `"*"`, or `"/"`), or an integer in the range `[-200, 200]`.

## 💡 Approach

In Reverse Polish notation, an operator acts on the last two operands or on the 
last operand and the last result (in that order \<last operand> \<operator> 
\<last result>). For this pattern, we need a data structure that maintains the 
processing order of the elements and allows efficient access to the last two 
(LIFO). This can be done using a stack.

With the help of a stack, we iterate through the tokens list checking if it is 
an operator or an operand. If it is an operator, pop the last two elements from 
the stack, compute the result, and push it onto the stack. Otherwise, push it 
onto the stack. In the end, the only element present in the stack will be the 
final result.  

## ⚡ Complexity Analysis

- **Time Complexity**: O(n)

We iterate through the tokens list doing constant work (computing the result 
of an operation, and pushing and popping elements from the stack) once. So,

`O(n) * 3*O(1) = O(n)`

- **Space Complexity**: O(n)

In a valid Reverse Polish notation, there will be two operands per operator. 
Thus, since the stack contains only operands and/or calculation results, its 
size will be the total number of operands at most. So,

`O(2/3 * n) = O(n)`

## 💻 Code

See the [solution.py](solution.py) file.
