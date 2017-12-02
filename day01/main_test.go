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
    }
    for _, c := range cases {
        actual := SumIntsThatRepeat(c.Input)
        if actual != c.Expected {
            t.Errorf("Error on %s. Expected %d but got %d", c.Input, c.Expected, actual)
        }
    }
}
