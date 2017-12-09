package grid_memory

type Coord struct {
    Row int
    Col int
}

func (self Coord) Multiply(scalar int) Coord {
    return Coord{scalar * self.Row, scalar * self.Col}
}

func (self Coord) Add(other Coord) Coord {
    return Coord{self.Row + other.Row, self.Col + other.Col}
}

func (self Coord) Equals(other Coord) bool {
    return self.Row == other.Row && self.Col == other.Col
}