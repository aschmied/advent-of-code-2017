package main

import (
    "errors"
    "strconv"
    "strings"
)

type Sheet struct {
    Rows [][]int
}

func ParseSheet(text string) Sheet {
    lines := strings.Split(strings.TrimSpace(text), "\n")
    rows := make([][]int, len(lines))

    for i, line := range lines {
        rows[i] = parseRow(line)
    }

    return Sheet{rows}
}

func parseRow(line string) []int {
    tokens := strings.FieldsFunc(strings.TrimSpace(line), func (c rune) bool {
        return c == '\t' || c == ' '
    })
    row := []int{}

    for _, token := range tokens {
        cell, err := parseToken(token)
        if err != nil {
            continue
        }

        row = append(row, cell)
    }

    return row
}

func parseToken(token string) (int, error) {
    trimmed := strings.TrimSpace(token)
    if len(trimmed) == 0 {
        return 0, errors.New("empty token")
    }

    parsed, err := strconv.Atoi(trimmed)
    if err != nil {
        return 0, err
    }

    return parsed, nil
}

func CalculateChecksumForSheet(sheet Sheet) int {
    checksum := 0
    for _, row := range sheet.Rows {
        checksum += CalculateChecksumForRow(row)
    }
    return checksum
}

func CalculateChecksumForRow(row []int) int {
    if len(row) == 0 {
        return 0
    }

    min := min(row)
    max := max(row)
    return max - min
}

func min(values []int) int {
    min := values[0]
    for _, value := range values {
        if value < min {
            min = value
        }
    }
    return min
}

func max(values []int) int {
    max := values[0]
    for _, value := range values {
        if value > max {
            max = value
        }
    }
    return max
}
