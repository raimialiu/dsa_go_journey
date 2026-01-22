# DSA Go Journey

A personal Data Structures and Algorithms learning repository in Go, tracking my journey to mastery over the next four years (starting January 2026).

## About

This project documents my genuine progress in understanding algorithmic problem-solving. All implementations are written by hand as part of deliberate practice—no AI-generated content.

## Project Structure

```
dsa_go_journey/
├── main.go              # Entry point with example usage
├── problems/            # Algorithm implementations by category
│   └── arrays/          # Array-based problems
│       ├── two_sum.go
│       └── two_sum_test.go
└── go.mod
```

## Implemented Algorithms

### Arrays
- **Two Sum** - Find indices of two numbers that add up to a target value

## Requirements

- Go 1.24+

## Usage

### Run the main program
```bash
go run main.go
```

### Run tests
```bash
go test ./...
```

### Run tests with verbose output
```bash
go test ./... -v
```

### Run benchmarks
```bash
go test -bench=. ./problems/arrays
```

## Example

```go
package main

import (
    "fmt"
    "dsa_go_journey/problems/arrays"
)

func main() {
    numbers := []int{2, 7, 11, 15}
    target := 9
    result := arrays.TwoSums(numbers, target)
    fmt.Println(result) // Output: [0 1]
}
```

## Roadmap

This repository will grow to include:
- More array algorithms
- Linked lists
- Trees and graphs
- Sorting and searching
- Dynamic programming
- And more...

## License

Personal learning project.
