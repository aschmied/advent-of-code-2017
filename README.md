## Advent of Code 2017

This repo contains my solutions to the [Advent of Code 2017](http://adventofcode.com/2017) puzzles. My goal this year is to learn Go while solving the puzzles. Feedback welcome.

## Example

### Using Your Go Workspace

If you have a go workspace, then clone the repo into your src directory:

    cd $(go env GOPATH)/src
    mkdir -p github.com/aschmied && cd github.com/aschmied
    git clone git@github.com:aschmied/advent-of-code-2017.git

Then change to a solution directory, install, and run:

    cd advent-of-code-2017/day01
    go install && $(go env GOPATH)/bin/day01

The input file paths are hardcoded to the working directory in the solutions, so you need to execute them in the source directory.

### With `go run`

If you don't have a go workspace then you can use `go run`. The drawback is that you need to explicitly list all source files:

    cd day01
    go run main.go sum_ints_where_offset_char_matches.go

## Tests

Each solution has tests. Run them like:

    cd day01
    go test
