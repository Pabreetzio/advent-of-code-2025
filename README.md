# Advent of Code 2025

Solutions for [Advent of Code 2025](https://adventofcode.com/2025) in Go.

## Structure

```
.
├── common/          # Shared utilities
│   └── input.go     # Input file parsing utilities
├── day01/           # Day 1 solution
│   ├── main.go
│   └── input.txt
├── day02/           # Day 2 solution
│   ├── main.go
│   └── input.txt
└── ...
```

## Running Solutions

To run a specific day's solution:

```bash
go run ./day01
```

Or build and run:

```bash
cd day01
go build
./day01
```

## Adding a New Day

1. Create a new folder for the day (e.g., `day03`)
2. Add `main.go` with your solution
3. Add `input.txt` with your puzzle input
4. Use the `common` package to read the input file
