# Dojo CLI

A command-line tool to manage the Code Dojo repository.

## Features

- **Create new problems** with automatic ID assignment
- **Parse problem metadata** from YAML frontmatter
- **Generate tag indexes** automatically
- **Validate** problem metadata for completeness

## Installation

### Option 1: Install to $GOPATH/bin (recommended)

From the repository root:

```bash
make install
```

Then run from anywhere:

```bash
dojo add
dojo sync
```

### Option 2: Build local binary

```bash
make build
./bin/dojo add
```

### Option 3: Run directly with go run

```bash
go run ./tools/dojo-cli/cmd/dojo add
```

## Usage

### Add a New Problem

Interactive mode:

```bash
dojo add
```

With flags:

```bash
dojo add --title "Two Sum" --tags "hashmap,array" --source "LeetCode" --ext py
```

This will:
- Automatically assign the next available problem ID
- Create a folder `problems/XXXX-problem-name/`
- Generate a `README.md` with YAML frontmatter template
- Create a `solution.{ext}` file with basic comment header

Supported extensions: `py`, `cpp`, `c`, `java`, `js`, `ts`, `go`, `rs`, `rb`, `php`, `swift`, `kt`, `scala`, and more.

### Sync Tag Indexes

Parse all problems and regenerate tag index files:

```bash
dojo sync
```

Sync a specific tag only:

```bash
dojo sync --tag hashmap
```

This will:
- Scan all `problems/*/README.md` files
- Extract metadata from YAML frontmatter
- Generate/update `tags/*.md` files with problem lists
- Update root `README.md` with current list of tags (sorted alphabetically)

### Validate Problems

Check all problems for valid metadata:

```bash
dojo validate
```

This will report any problems missing required fields (id, title, tags, source).

## Problem Structure

Each problem README should have YAML frontmatter:

```markdown
---
id: 1
title: Two Sum
tags: [array, hash-table]
source: LeetCode
---

# Two Sum

**📚 Source**: LeetCode

**🏷️ Tags**: [array](../../tags/array.md), [hash-table](../../tags/hash-table.md)

## 📋 Problem Statement

Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

## 💡 Approach

Use a hash table to store complements.

## ⚡ Complexity Analysis

- **Time Complexity**: O(n)
- **Space Complexity**: O(n)

## 💻 Code

See `solution.*` files in this directory.
```

## Development

### Project Structure

```
tools/dojo-cli/
├── cmd/
│   └── dojo/
│       └── main.go           # Entry point
├── internal/
│   ├── models/
│   │   └── problem.go        # Problem data model
│   ├── parser/
│   │   └── metadata.go       # YAML frontmatter parser
│   └── generator/
│       ├── problem.go        # Problem folder generator
│       ├── tags.go           # Tag index generator
│       └── readme.go         # Root README updater
├── pkg/
│   └── cli/
│       ├── root.go           # Root command
│       ├── add.go            # Add problem command
│       ├── sync.go           # Sync tags command
│       └── validate.go       # Validate command
├── templates/
│   ├── problem_readme.tmpl   # Problem README template
│   └── tag_index.tmpl        # Tag index template
└── go.mod
```

### Run Tests

```bash
make test
```

### Clean Build Artifacts

```bash
make clean
```

## License

Part of the Code Dojo repository.
