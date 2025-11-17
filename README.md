# 🎄 Advent of Go Template 🎄

This repository serves as a template for solving Advent of Code puzzles in Go.

## How to use

1. You can use the [Github Template Feature](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template) or simply download the repository to your machine.
2. Run `make init` to initialize empty `input.txt` files. The files are gitignored as per the AoC ToS.
3. Solve a specific day by copying your personal input file into the respective `input.txt` of the day's directory and implement the `SolvePuzzle1` / `SolvePuzzle2`.
4. Run the specific day via `make day-<day_number>` or run the test file via your favourite IDE (make sure to use verbose mode).
5. Validate the printed results on the [Advent of Code website](https://adventofcode.com/).
