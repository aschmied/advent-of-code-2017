package main

import "github.com/aschmied/advent-of-code-2017/day03/grid_memory"

import "fmt"
import "math"

const input int = 361527

func main() {
    fmt.Println(solvePartOne())
    fmt.Println(solvePartTwo())
}

func solvePartOne() int {
    gridMemory := grid_memory.New(input)
    for i := 1; i < input; i++ {
        gridMemory.Push(i)
    }
    coord := gridMemory.Coord()
    manhattanDistanceFromOrigin := math.Abs(float64(coord.Row)) + math.Abs(float64(coord.Col))
    return int(manhattanDistanceFromOrigin)
}

func solvePartTwo() int {
    gridMemory := grid_memory.New(input)
    gridMemory.Push(1)

    for true {
        patchOrigin := gridMemory.Coord()
        sum := gridMemory.SumInPatch(patchOrigin)
        if sum > input {
            return sum
        }
        gridMemory.Push(sum)
    }

    return -1
}