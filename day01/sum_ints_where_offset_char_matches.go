package main

func SumIntsWhereOffsetCharMatches(s string, offset int) int {
    accumulate := makeAccumulator()

    for i := 0; i < len(s); i++ {
        char0 := s[i]
        char1 := s[(i + offset) % len(s)]
        if char0 == char1 {
            accumulate(btoi(char0))
        }
    }

    return accumulate(0)
}

func btoi(b byte) int {
    return int(b - '0')
}

func makeAccumulator() func(int) int {
    var runningSum int
    return func(next int) int {
        runningSum += next
        return runningSum
    }
}
