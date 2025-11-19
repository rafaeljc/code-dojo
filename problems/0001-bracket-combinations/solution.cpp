// Bracket Combinations
// Source: Coderbyte
// Problem: https://github.com/rafaeljc/code-dojo/tree/main/problems/0001-bracket-combinations

enum Error {
    InvalidInput = -1,
};

int BracketCombinations(int num) {
    if (num < 0) {
        return Error::InvalidInput;
    }
    // Assuming that the zero pairs of brackets does not produce any valid
    // combinations
    if (num == 0) {
        return 0;
    }
    if (num == 1) {
        return 1;
    }
    int combinations = 1; // C_0 = 1
    for (int n = 1; n <= num; n++) {
        combinations *= 2 * (2*n - 1);
        combinations /= (n + 1);
    }
    return combinations;
}
