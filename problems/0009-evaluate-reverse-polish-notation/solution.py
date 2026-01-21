# Evaluate Reverse Polish Notation
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0009-evaluate-reverse-polish-notation

class Solution:
    class EmptyTokenListError(Exception):
        pass

    class InvalidExpressionError(Exception):
        pass

    class InvalidTokenError(Exception):
        pass

    def evalRPN(self, tokens: list[str]) -> int:
        if len(tokens) == 0:
            raise Solution.EmptyTokenListError("The token list is empty.")
        stack = []
        for i, token in enumerate(tokens):
            match token:
                case "+" | "-" | "*" | "/":
                    if len(stack) < 2:
                        raise Solution.InvalidExpressionError(
                            f"Not enough operands for operator '{token}' "
                            f"at position {i}."
                        )
                    b = stack.pop()
                    a = stack.pop()
                    if token == "+":
                        stack.append(a + b)
                    elif token == "-":
                        stack.append(a - b)
                    elif token == "*":
                        stack.append(a * b)
                    else: # token == "/":
                        if b == 0:
                            raise Solution.InvalidExpressionError(
                                f"Division by zero at position {i}."
                            )
                        stack.append(int(a / b)) # int() truncates towards zero
                case _:
                    try:
                        value = int(token)
                        stack.append(value)
                    except ValueError:
                        # token is neither an operator nor a valid integer
                        raise Solution.InvalidTokenError(
                            f"Invalid token '{token}' at position {i}."
                        )
        if len(stack) != 1:
            raise Solution.InvalidExpressionError(
                "The number of operands and operators is unbalanced."
            )
        return stack.pop()
