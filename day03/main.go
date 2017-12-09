package main

import "github.com/aschmied/advent-of-code-2017/day03/grid_memory"

import "fmt"
import "math"

const input int = 361527

func main() {
    gridMemory := grid_memory.New(input)
    for i := 1; i < input; i++ {
        gridMemory.Push(i)
    }
    coord := gridMemory.Coord()
    manhattanDistanceToOrigin := math.Abs(float64(coord.Row)) + math.Abs(float64(coord.Col))
    fmt.Println(manhattanDistanceToOrigin)
}
