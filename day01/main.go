package main

func SumIntsThatRepeat(s string) int {
    if len(s) == 0 {
        return 0
    }

    sum := 0
    for index, char0 := range s[:len(s) - 1] {
        char1 := int32(s[index + 1])
        if char0 == char1 {
            parsed := int(char0 - '0')
            sum += parsed
        }
    }

    if s[0] == s[len(s) - 1] {
        sum += int(s[0] - '0')
    }

    return sum
}

func main() {
    
}
