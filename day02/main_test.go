package main

import "testing"

func TestCalculateChecksumForRow(t *testing.T) {
    cases := []struct {
        Input [] int
        Expected int
    } {
        {[]int{}, 0},
        {[]int{55}, 0},
        {[]int{2, 10}, 8},
        {[]int{-2, 2}, 4},
        {[]int{-22, -10}, 12},
        {[]int{3, -4, 57, 22, -100, 20, 100, 3}, 200},
    }

    for _, c := range cases {
        actual := CalculateChecksumForRow(c.Input)
        if actual != c.Expected {
            t.Errorf("Error on %v. Expcted %v but got %v", c.Input, c.Expected, actual)
        }
    }
}
