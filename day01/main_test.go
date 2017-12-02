package main

import "testing"

func TestScanAndSum(t *testing.T) {
    cases := []struct {
        Input string
        Expected int
    } {
        {"", 0},
        {"1", 1},
        {"1234323", 0},
        {"112", 1},
        {"11233", 4},
        {"1121", 2},
        {"1122", 3},
        {"1111", 4},
        {"1234", 0},
        {"91212129", 9},
    }
    for _, c := range cases {
        actual := SumIntsWhereOffsetCharMatches(c.Input, 1)
        if actual != c.Expected {
            t.Errorf("Error on %s. Expected %d but got %d", c.Input, c.Expected, actual)
        }
    }
}
