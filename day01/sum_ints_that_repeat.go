package main

func SumIntsThatRepeat(s string) int {
    if len(s) == 0 {
        return 0
    }

    sum := runningSum{}
    for index, char0 := range s[:len(s) - 1] {
        char1 := s[index + 1]
        sum.addIfEqual(btoi(byte(char0)), btoi(char1))
    }
    sum.addIfEqual(btoi(s[0]), btoi(s[len(s) - 1]))

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
