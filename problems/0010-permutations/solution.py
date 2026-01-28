# Permutations
# Source: LeetCode
# Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0010-permutations

class Solution:
    def _buildPermutations(
        self, nums: list[int], k: int, permutations: list[list[int]]
    ) -> None:
        if k == 1:
            # Append a copy of the current permutation to the result list
            permutations.append(nums[:])
            return
        self._buildPermutations(nums, k - 1, permutations)
        for i in range(k - 1):
            if k % 2 == 0:
                # Swap the i-th and (k-1)-th elements
                nums[i], nums[k - 1] = nums[k - 1], nums[i]
            else:
                # Swap the first and (k-1)-th elements
                nums[0], nums[k - 1] = nums[k - 1], nums[0]
            self._buildPermutations(nums, k - 1, permutations)
        return

    def permute(self, nums: list[int]) -> list[list[int]]:
        size = len(nums)
        if size == 0:
            return []
        permutations = []
        # Use a copy of 'nums' to avoid modifying the original list
        self._buildPermutations(nums[:], size, permutations)
        return permutations
