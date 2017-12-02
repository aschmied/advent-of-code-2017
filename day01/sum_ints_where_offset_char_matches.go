package main

func SumIntsWhereOffsetCharMatches(s string, offset int) int {
    sum := runningSum{}
    for i := 0; i < len(s); i++ {
        char0 := s[i]
        char1 := s[(i + offset) % len(s)]
        sum.addIfEqual(btoi(char0), btoi(char1))
    }

    return sum.sum
}

func btoi(b byte) int {
    return int(b - '0')
}

type runningSum struct {
    sum int
}

func (s *runningSum) addIfEqual(c0 int, c1 int) {
    if c0 != c1 {
        return
    }
    s.sum += c0
}
