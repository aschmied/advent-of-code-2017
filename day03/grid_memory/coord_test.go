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

func TestEquals(t *testing.T) {
    assertTrue(t, Coord{1, 2}.Equals(Coord{1, 2}))
    assertFalse(t, Coord{1, 2}.Equals(Coord{3, 2}))
    assertFalse(t, Coord{1, 2}.Equals(Coord{1, 3}))
}

func assertCoordsEqual(t *testing.T, actual Coord, expected Coord) {
    if actual.Row != expected.Row || actual.Col != expected.Col {
        t.Errorf("Expected (%v, %v) but got (%v, %v)",
            expected.Row, expected.Col, actual.Row, actual.Col)
    }
}

func assertTrue(t *testing.T, actual bool) {
    assertBool(t, actual, true)
}

func assertFalse(t *testing.T, actual bool) {
    assertBool(t, actual, false)
}

func assertBool(t *testing.T, actual bool, expected bool) {
    if actual != expected {
        t.Errorf("Expected %v but got %v", expected, actual)
    }
}