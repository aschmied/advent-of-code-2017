package grid_memory

import "testing"

func TestSpiraling(t *testing.T) {
    expectedCoordsBeforePush := []Coord {
        {0, 0},
        {0, 1},
        {-1, 1},
        {-1, 0},
        {-1, -1},
        {0, -1},
        {1, -1},
        {1, 0},
        {1, 1},
        {1, 2},
        {0, 2},
        {-1, 2},
        {-2, 2},
        {-2, 1},
        {-2, 0},
    }
    
    n := len(expectedCoordsBeforePush)
    gridMemory := New(n)
    for i := 0; i < n; i++ {
        assertCurrentCoord(t, gridMemory, expectedCoordsBeforePush[i])
        err := gridMemory.Push(i)
        assertErrorIsNil(t, err)
    }
}

func assertErrorIsNil(t *testing.T, err error) {
    if err != nil {
        t.Errorf("Expected error to be nil: %v", err)
    }
}

func assertCurrentCoord(t *testing.T, gridMemory gridMemory, expected Coord) {
    actual := gridMemory.Coord()
    if actual.Row != expected.Row || actual.Col != expected.Col {
        t.Errorf("Coord incorrect after pushing %v. Expected (%v, %v) but got (%v, %v)",
            gridMemory.Top(), expected.Row, expected.Col, actual.Row, actual.Col)
    }
}
