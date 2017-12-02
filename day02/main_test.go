package main

import "reflect"
import "testing"

func TestParseSheet(t *testing.T) {
    cases := []struct {
        Input string
        Expected Sheet
    } {
        {"", Sheet{[][]int{{}}}},
        {"0", Sheet{[][]int{{0}}}},
        {"0\n", Sheet{[][]int{{0}}}},
        {"0\n0\n", Sheet{[][]int{{0}, {0}}}},
        {"0 0\n", Sheet{[][]int{{0, 0}}}},
        {"-10  20 14  2 \n 1", Sheet{[][]int{{-10, 20, 14, 2}, {1}}}},
        {"1\t2\n3\t4", Sheet{[][]int{{1, 2},{3, 4}}}},
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
        Input []int
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
