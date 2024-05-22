
# Palindrome Detection Using Manacher's Algorithm

This project implements a modified version of Manacher's Algorithm in Go to find all palindromic substrings within a given string.

## Files
- `main.go`: Contains the Go code implementing the algorithm.

## Overview
The program consists of several key functions:

### `preprocess(original string) string`
Transforms the input string by inserting special characters (`#`) between each character of the original string to handle even and odd length palindromes uniformly. It also adds boundary markers (`^` and `$`) to prevent out-of-bound errors during the palindrome expansion checks.

### `findAllPalindromes(input string) []string`
Implements the main logic of Manacher's Algorithm to find and store all palindromic substrings. It uses a transformed version of the input string and processes it to calculate palindromic radii, adjust centers and boundaries, and store the results.

### `min(a, b int) int`
A helper function to compute the minimum of two integers, useful for determining the initial radius based on known boundaries and mirrored positions.

## Usage
Run the program by executing the `main.go` file. The input string is defined within the `main` function. The program outputs all palindromic substrings found in the input string.

## Example
Given the input "aabbbaa", the expected output will be:
```
Palindromes found: [aa bb bbb aa]
```

This example illustrates the algorithm's ability to find multiple palindromic substrings, including overlapping and nested palindromes.
