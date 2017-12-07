package grid_memory

import "testing"

func TestAdd(t *testing.T) {
    c1 := Coord{1, 2}
    c2 := Coord{2, 3}
    c3 := c1.Add(c2)
    assertCoordsEqual(t, c1, Coord{1, 2})
    assertCoordsEqual(t, c2, Coord{2, 3})
    assertCoordsEqual(t, c3, Coord{3, 5})
}

func TestMultiply(t *testing.T) {
    c1 := Coord{1, 2}
    c2 := c1.Multiply(2)
    assertCoordsEqual(t, c1, Coord{1, 2})
    assertCoordsEqual(t, c2, Coord{2, 4})
}

func assertCoordsEqual(t *testing.T, actual Coord, expected Coord) {
    if actual.Row != expected.Row || actual.Col != expected.Col {
        t.Errorf("Expected (%v, %v) but got (%v, %v)",
            expected.Row, expected.Col, actual.Row, actual.Col)
    }
}
