package grid_memory

import "testing"

func TestEmpty(t *testing.T) {
    gridMemory := New(0)
    err := gridMemory.Push(0)
    assertErrorNotNil(t, err)
}

func TestCapacityOfOne(t *testing.T) {
    gridMemory := New(1)
    err := gridMemory.Push(0)
    assertErrorIsNil(t, err)
    err = gridMemory.Push(0)
    assertErrorNotNil(t, err)
}

func TestLargeCapacity(t *testing.T) {
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
        assertCurrentCoord(t, gridMemory, expectedCoordsBeforePush[i], i)
        err := gridMemory.Push(i)
        assertErrorIsNil(t, err)
    }
}

func TestGet(t *testing.T) {
    gridMemory := New(4)
    gridMemory.Push(1)
    assertValueAt(t, gridMemory, Coord{0, 0}, 1)
    assertValueAt(t, gridMemory, Coord{0, 1}, 0)
    assertValueAt(t, gridMemory, Coord{-99, 99}, 0)
}

func TestSumInPatch(t *testing.T) {
    gridMemory := New(9)
    for i := 1; i <= 9; i++ {
        gridMemory.Push(i)
    }
    assertSumInPatch(t, gridMemory, Coord{0, 0}, 45)
    assertSumInPatch(t, gridMemory, Coord{0, 1}, 27)
    assertSumInPatch(t, gridMemory, Coord{0, 2}, 14)
    assertSumInPatch(t, gridMemory, Coord{0, 3}, 0)
}

func assertErrorNotNil(t *testing.T, err error) {
    if err == nil {
        t.Error("Expected error not to be nil")
    }
}

func assertErrorIsNil(t *testing.T, err error) {
    if err != nil {
        t.Errorf("Expected error to be nil: %v", err)
    }
}

func assertCurrentCoord(t *testing.T, gridMemory gridMemory, expected Coord, index int) {
    actual := gridMemory.Coord()
    if actual.Row != expected.Row || actual.Col != expected.Col {
        t.Errorf("Coord incorrect before pushing %v. Expected (%v, %v) but got (%v, %v)",
            index, expected.Row, expected.Col, actual.Row, actual.Col)
    }
}

func assertValueAt(t *testing.T, gridMemory gridMemory, location Coord, expected int) {
    actual := gridMemory.Get(location)
    if actual != expected {
        t.Errorf("Coord at (%v, %v) is incorrect. Expected %v but got %v", location.Row, location.Col, expected, actual)
    }
}

func assertSumInPatch(t *testing.T, gridMemory gridMemory, patchOrigin Coord, expected int) {
    actual := gridMemory.SumInPatch(patchOrigin)
    if actual != expected {
        t.Errorf("Patch surrounding (%v, %v) has incorrect sum. Expected %v but got %v", patchOrigin.Row, patchOrigin.Col, expected, actual)
    }
}