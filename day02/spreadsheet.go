package main

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
