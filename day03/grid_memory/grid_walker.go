package grid_memory

type gridWalker struct {
    location Coord
    directions []Coord
    currentDirectionIndex int
    currentSideLength int
    locationOfNextTurn Coord
}

func newGridWalker() gridWalker {
    location := Coord{0, 0}
    directions := []Coord {
        {0, 1},
        {-1, 0},
        {0, -1},
        {1, 0},
    }
    currentDirectionIndex := 0
    currentSideLength := 1
    locationOfNextTurn := Coord{0, currentSideLength}
    return gridWalker{
        location,
        directions,
        currentDirectionIndex,
        currentSideLength,
        locationOfNextTurn,
    }
}

func (self gridWalker) CurrentCoord() Coord {
    return self.location
}

func (self *gridWalker) Step() Coord {
    previousLocation := self.step()
    if self.location.Equals(self.locationOfNextTurn) {
        self.turn()
        self.calculateLocationOfNextTurn()
    }
    return previousLocation
}

func (self *gridWalker) step() Coord {
    currentLocation := self.location
    direction := self.directions[self.currentDirectionIndex]
    self.location = currentLocation.Add(direction)
    return currentLocation
}

func (self *gridWalker) turn() {
    self.currentDirectionIndex = (self.currentDirectionIndex + 1) % 4
    if self.currentDirectionIndex % 2 == 0 {
        self.currentSideLength += 1
    }

}

func (self *gridWalker) calculateLocationOfNextTurn() {
    direction := self.directions[self.currentDirectionIndex]
    self.locationOfNextTurn = self.location.Add(direction.Multiply(self.currentSideLength))
}
