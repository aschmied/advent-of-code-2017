package grid_memory

import "errors"
import "math"

type gridMemory struct {
    length int
    capacity int
    grid [][]int
    walker gridWalker
}

func New(capacity int) gridMemory {
    sideLength := int(math.Ceil(math.Sqrt(float64(capacity))))
    indexOffset := sideLength / 2

    grid := make([][]int, sideLength)
    for i, _ := range grid {
        grid[i] = make([]int, sideLength)
    }

    return gridMemory{0, capacity, grid, newGridWalker()}
}

func (self *gridMemory) Push(value int) error {
    if self.length >= self.capacity {
        return errors.New("Grid memory full")
    }

    coord := self.walker.nextCoord()
    self.grid[coord.Row][coord.Col] = value
    return nil
}

func (self gridMemory) Top() int {
    coord := self.Coord()
    return self.grid[coord.Row][coord.Col]
}

func (self gridMemory) Coord() Coord {
    return self.walker.currentCoord()
}





// gridMemory = GridMemory{input}
// for i := 1; i < input + 1; i++ {
//     err = gridMemory.push(i)
// }

// manhattanDistanceFromOne = |gridMemory.Row| + |gridMemory.Col|
// fmt.Println(manhattanDistanceFromOne)

// gridMemory = GridMemory{input}
// for i := 1; gridMemory.currentPatchSum(3) < input; i++ {
//     err = gridMemory.push(i)
// }

// manhattanDistanceFromOne = |gridMemory.Row| + gridMemory.Col|
