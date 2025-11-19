# 🥋 code-dojo

Welcome to **Code Dojo** — a space for sharpening the mind through consistent practice in data structures, algorithms, and problem-solving.

> "The way is in training."
> 
> — Miyamoto Musashi

Just as a martial artist trains daily in the dojo to master their craft, this repository is my training ground for mastering computer science fundamentals and developing strong problem-solving skills through deliberate, focused practice.


## 🎯 Purpose

This is a long-term commitment to continuous improvement through:
- **Regular Practice** — Solving problems consistently to build and maintain skills
- **Deep Understanding** — Going beyond solutions to truly grasp core computer science concepts
- **Thoughtful Reflection** — Analyzing approaches, learning from mistakes, and refining techniques
- **Knowledge Building** — Creating a personal reference of patterns, solutions, and insights


## 📁 Structure

Each problem has its own folder with:
- `solution.*` — My implementation
- `README.md` — Problem statement, thought process and complexity analysis

Problems come from:
- [LeetCode](https://leetcode.com/)
- [HackerRank](https://www.hackerrank.com/)
- [Coderbyte](https://coderbyte.com/)
- Other sources


## 🏷️ Tags

Problems are organized by topic using tag indexes in `/tags/`.
- [Backtracking](tags/backtracking.md)
- [Combinatorics](tags/combinatorics.md)

## 🛠️ Tools

### Dojo CLI

A command-line tool to automate repository management:

- **Add new problems** with automatic ID assignment and folder structure
- **Generate solution files** with language-specific templates
- **Sync tag indexes** automatically from problem metadata
- **Validate** problem metadata for completeness
- **Update** root README with current tags

```bash
# Install the CLI
make install

# Create a new problem
dojo add --title "Two Sum" --tags "array,hash-table" --source "LeetCode" --ext py

# Sync all tags and update README
dojo sync

# Validate all problems
dojo validate
```

See [tools/dojo-cli/README.md](tools/dojo-cli/README.md) for full documentation.
