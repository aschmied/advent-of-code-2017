package grid_memory

import "errors"
import "math"

type gridMemory struct {
    length int
    capacity int
    grid [][]int
    origin Coord
    walker gridWalker
}

func New(capacity int) gridMemory {
    sideLength := sideLength(capacity)
    origin := Coord{sideLength / 2, sideLength / 2}

    grid := make([][]int, sideLength)
    for i, _ := range grid {
        grid[i] = make([]int, sideLength)
    }

    return gridMemory{0, capacity, grid, origin, newGridWalker()}
}

func sideLength(capacity int) int {
    candidate := int(math.Ceil(math.Sqrt(float64(capacity))))
    if candidate % 2 == 0 {
        return candidate + 1
    } else {
        return candidate
    }
}

func (self *gridMemory) Push(value int) error {
    if self.length >= self.capacity {
        return errors.New("Out of grid memory")
    }

    coord := self.walker.Step()
    arrayCoord := self.arrayCoord(coord)
    self.grid[arrayCoord.Row][arrayCoord.Col] = value
    self.length += 1
    return nil
}

func (self gridMemory) Top() int {
    arrayCoord := self.arrayCoord(self.Coord())
    return self.grid[arrayCoord.Row][arrayCoord.Col]
}

func (self gridMemory) Coord() Coord {
    return self.walker.CurrentCoord()
}

func (self gridMemory) Get(coord Coord) int {
    arrayCoord := self.arrayCoord(coord)
    if arrayCoord.Row >= len(self.grid) || arrayCoord.Col >= len(self.grid[0]) {
        return 0
    }
    return self.grid[arrayCoord.Row][arrayCoord.Col]
}

func (self gridMemory) SumInPatch(coord Coord) int {
    sum := 0
    for i := -1; i < 2; i++ {
        for j := -1; j < 2; j++ {
            patchOffset := Coord{i, j}
            sum += self.Get(coord.Add(patchOffset))
        }
    }
    return sum
}

func (self gridMemory) arrayCoord(coord Coord) Coord {
    arrayRow := self.origin.Row + coord.Row
    arrayCol := self.origin.Col + coord.Col
    return Coord{arrayRow, arrayCol}
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
