package main

import "reflect"
import "testing"

func TestParseSheet(t *testing.T) {
    assertParsedSheet(t, "", MakeSheet())
    assertParsedSheet(t, "0", MakeSheet(MakeRow(0)))
    assertParsedSheet(t, "0\n", MakeSheet(MakeRow(0)))
    assertParsedSheet(t, "0\n0\n", MakeSheet(MakeRow(0), MakeRow(0)))
    assertParsedSheet(t, "0 0\n", MakeSheet(MakeRow(0, 0)))
    assertParsedSheet(t, "-10  20 14  2 \n 1", MakeSheet(MakeRow(-10, 20, 14, 2), MakeRow(1)))
    assertParsedSheet(t, "1\t2\n3\t4", MakeSheet(MakeRow(1, 2), MakeRow(3, 4)))
}

func assertParsedSheet(t *testing.T, input string, expected Sheet) {
    actual := ParseSheet(input)
    if !reflect.DeepEqual(actual.Rows, expected.Rows) {
        t.Errorf("Error on %v. Expected %v but got %v", input, expected, actual)
    }
}

func TestCalculateChecksumForRow(t *testing.T) {
    assertChecksumForRow(t, MakeRow(), 0)
    assertChecksumForRow(t, MakeRow(55), 0)
    assertChecksumForRow(t, MakeRow(2, 10), 8)
    assertChecksumForRow(t, MakeRow(-2, 2), 4)
    assertChecksumForRow(t, MakeRow(-22, -10), 12)
    assertChecksumForRow(t, MakeRow(3, -4, 57, 22, -100, 20, 100, 3), 200)
}

func assertChecksumForRow(t *testing.T, input Row, expected int) {
    actual := CalculateChecksumForRow(input)
    if actual != expected {
        t.Errorf("Error on %v. Expcted %v but got %v", input, expected, actual)
    }
}

func TestSumOfMultipliers(t *testing.T) {
    assertSumOfMultipliers(t, MakeSheet(MakeRow(5, 5)), 1)
    assertSumOfMultipliers(t, MakeSheet(MakeRow(2, 3, 5, 4, 7)), 2)
    assertSumOfMultipliers(t, MakeSheet(MakeRow(2, 3, 5, 7, 2)), 1)
    assertSumOfMultipliers(t, MakeSheet(MakeRow(2, 3, 5, 4, 7), MakeRow(2, 3, 5, 7, 2)), 3)
}

func assertSumOfMultipliers(t *testing.T, input Sheet, expected int) {
    actual := SumOfMultipliers(input)
    if actual != expected {
        t.Errorf("Error on %v. Expected %v but got %v", input, expected, actual)
    }
}
