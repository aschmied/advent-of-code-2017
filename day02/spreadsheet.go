package main

import (
    "log"
    "strconv"
    "strings"
)

type Sheet struct {
    Rows []Row
}

func MakeSheet(rows ...Row) Sheet {
    if len(rows) == 0 {
        return MakeSheet(MakeRow())
    }
    return Sheet{rows}
}

type Row struct {
    Cells []int
}

func MakeRow(cells ...int) Row {
    if len(cells) == 0 {
        return Row{[]int{}}
    }
    return Row{cells}
}

func ParseSheet(text string) Sheet {
    trimmedText := strings.TrimSpace(text)
    lines := strings.Split(trimmedText, "\n")
    rows := make([]Row, len(lines))

    for i, line := range lines {
        rows[i] = parseLine(line)
    }

    return Sheet{rows}
}

func parseLine(line string) Row {
    tokens := tokenizeLine(line)
    cells := make([]int, len(tokens))

    for i, token := range tokens {
        cells[i] = parseToken(token)
    }

    return Row{cells}
}

func tokenizeLine(line string) []string {
    return strings.FieldsFunc(line, func (r rune) bool {
        return r == '\t' || r == ' '
    })
}

func parseToken(token string) int {
    parsed, err := strconv.Atoi(token)
    if err != nil {
        log.Fatalf("Error parsing token %v: %v\n", token, err)
    }

    return parsed
}

func CalculateChecksumForSheet(sheet Sheet) int {
    checksum := 0
    for _, row := range sheet.Rows {
        checksum += CalculateChecksumForRow(row)
    }
    return checksum
}

func CalculateChecksumForRow(row Row) int {
    if len(row.Cells) == 0 {
        return 0
    }

    min := min(row.Cells)
    max := max(row.Cells)
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
