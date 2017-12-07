package grid_memory

type gridWalker struct {
    current Coord
    directions []Coord
    nextDirectionIndex int
    nextDistance int
}

func newGridWalker() gridWalker {
 current := Coord{0, 0}
    directions := []Coord {
        {0, 1},
        {-1, 0},
        {0, -1},
        {1, 0},
    }
    nextDirectionIndex := 0
    nextDistance := 1
    return gridWalker{current, directions, nextDirectionIndex, nextDistance}   
}

func (self gridWalker) currentCoord() Coord {
    return self.current
}

func (self *gridWalker) nextCoord() Coord {
    nextDirection := self.directions[self.nextDirectionIndex]
    offset := nextDirection.Multiply(self.nextDistance)
    self.current = self.current.Add(offset)

    self.nextDirectionIndex = (self.nextDirectionIndex + 1) % len(self.directions)
    if self.nextDirectionIndex % 2 == 0 {
        self.nextDistance += 1
    }

    return self.current
}
