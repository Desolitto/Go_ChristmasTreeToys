# ChristmasTreeToys

## Overview

ChristmasTreeToys is a collection of Go packages and functions designed for working with binary trees containing toys and managing a collection of presents. This project implements:

- Checking the balance of toys within binary trees.
- Zigzag (spiral) traversal of the tree to create a "garland" sequence.
- A custom heap for presents with user-defined comparison logic.
- A classic knapsack problem solution to select the optimal set of presents.

---

## Table of Contents

- [Introduction](#introduction)  
- [Features](#features)  
- [Getting Started](#getting-started)  
- [Project Structure](#project-structure)  

---

## Introduction

In programming, working with trees and data structures for optimal element selection is a common task. This project demonstrates:

- How to verify the balance of toy counts in a binary tree.
- How to traverse a tree in a zigzag pattern.
- How to implement a heap with custom comparison logic for selecting the best presents.
- How to solve the knapsack problem using dynamic programming to maximize present value under size constraints.

---

## Features

- Checks whether the number of toys in the left and right subtrees are balanced.
- Performs zigzag (spiral) traversal of a binary tree to generate a sequence of toy presence.
- Implements a present heap sorted by value (descending) and size (ascending).
- Provides functions to select the top-N coolest presents.
- Offers an optimal solution to the knapsack problem for selecting presents based on size and value.

---

## Getting Started

### Prerequisites

- Go 1.16 or higher
- Git (for cloning the repository)

### Installation

```bash
git clone https://github.com/Desolitto/Go_ChristmasTreeToys
cd Go_ChristmasTreeToys
go mod tidy
```

---

## Project Structure

```
ChristmasTreeToys/
│
├── pkg/
│   ├── presents.go       # Present struct, PresentHeap, and present-related functions
│   ├── tree.go           # TreeNode struct and tree-related functions
├── tests/
│   ├── presents_test.go  # Tests for presents package
│   ├── tree_test.go      # Tests for tree package
└── README.md             # Project documentation
```

---

## Explanation of the Directories and Files

- **pkg/presents.go** — contains the implementation of the Present struct, PresentHeap, and functions for managing presents.
- **pkg/tree.go** — contains the TreeNode struct and functions for checking toy balance and zigzag traversal.
- **tests/** — contains unit tests for both presents and tree packages.
- **README.md** — provides project documentation and usage instructions.

---
