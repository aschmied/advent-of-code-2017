package main

import "reflect"
import "testing"

func TestParseSheet(t *testing.T) {
    cases := []struct {
        Input string
        Expected Sheet
    } {
        {"", Sheet{[]Row{Row{[]int{}}}}},
        {"0", Sheet{[]Row{Row{[]int{0}}}}},
        {"0\n", Sheet{[]Row{Row{[]int{0}}}}},
        {"0\n0\n", Sheet{[]Row{Row{[]int{0}}, Row{[]int{0}}}}},
        {"0 0\n", Sheet{[]Row{Row{[]int{0, 0}}}}},
        {"-10  20 14  2 \n 1", Sheet{[]Row{Row{[]int{-10, 20, 14, 2}}, Row{[]int{1}}}}},
        {"1\t2\n3\t4", Sheet{[]Row{Row{[]int{1, 2}}, Row{[]int{3, 4}}}}},
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
        {Row{[]int{}}, 0},
        {Row{[]int{55}}, 0},
        {Row{[]int{2, 10}}, 8},
        {Row{[]int{-2, 2}}, 4},
        {Row{[]int{-22, -10}}, 12},
        {Row{[]int{3, -4, 57, 22, -100, 20, 100, 3}}, 200},
    }

    for _, c := range cases {
        actual := CalculateChecksumForRow(c.Input)
        if actual != c.Expected {
            t.Errorf("Error on %v. Expcted %v but got %v", c.Input, c.Expected, actual)
        }
    }
}
