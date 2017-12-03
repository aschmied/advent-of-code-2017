package main

import "reflect"
import "testing"

func TestParseSheet(t *testing.T) {
    cases := []struct {
        Input string
        Expected Sheet
    } {
        {"", MakeSheet()},
        {"0", MakeSheet(MakeRow(0))},
        {"0\n", MakeSheet(MakeRow(0))},
        {"0\n0\n", MakeSheet(MakeRow(0), MakeRow(0))},
        {"0 0\n", MakeSheet(MakeRow(0, 0))},
        {"-10  20 14  2 \n 1", MakeSheet(MakeRow(-10, 20, 14, 2), MakeRow(1))},
        {"1\t2\n3\t4", MakeSheet(MakeRow(1, 2), MakeRow(3, 4))},
    }

    for _, c := range cases {
        actual := ParseSheet(c.Input)
        if !reflect.DeepEqual(actual.Rows, c.Expected.Rows) {
            t.Errorf("Error on %v. Expected %v but got %v", c.Input, c.Expected, actual)
        }
    }
}

func TestCalculateChecksumForRow(t *testing.T) {
    cases := []struct {
        Input Row
        Expected int
    } {
        {MakeRow(), 0},
        {MakeRow(55), 0},
        {MakeRow(2, 10), 8},
        {MakeRow(-2, 2), 4},
        {MakeRow(-22, -10), 12},
        {MakeRow(3, -4, 57, 22, -100, 20, 100, 3), 200},
    }

    for _, c := range cases {
        actual := CalculateChecksumForRow(c.Input)
        if actual != c.Expected {
            t.Errorf("Error on %v. Expcted %v but got %v", c.Input, c.Expected, actual)
        }
    }
}

func TestSumOfMultipliers(t *testing.T) {
    cases := []struct {
        Input Sheet
        Expected int
    } {
        {MakeSheet(MakeRow(2, 3, 5, 4, 7)), 2},
        {MakeSheet(MakeRow(2, 3, 5, 7, 2)), 1},
        {MakeSheet(MakeRow(2, 3, 5, 4, 7), MakeRow(2, 3, 5, 7, 2)), 3},
    }

    for _, c := range cases {
        actual := SumOfMultipliers(c.Input)
        if actual != c.Expected {
            t.Error("Error on %v. Expected %v but got %v", c.Input, c.Expected, actual)
        }
    }
}

func TestFindPair(t *testing.T) {
    predicate := func (v1, v2 int) bool {
        return v1 == v2
    }
    cases := []struct {
        Input Row
        Expected int
    } {
        {MakeRow(1, 1), 1},
        {MakeRow(1, 2, 3, 4, 3, 5), 3},
    }

    for _, c := range cases {
        actual, _ := FindPair(c.Input, predicate)
        if actual != c.Expected {
            t.Errorf("Error on %v. Expcted %v but got %v", c.Input, c.Expected, actual)
        }
    }
}
