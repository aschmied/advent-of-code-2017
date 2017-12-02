package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func main() {
    input := readInput("input")
    fmt.Println(SumIntsWhereOffsetCharMatches(input, 1))
    fmt.Println(SumIntsWhereOffsetCharMatches(input, len(input) / 2))
}

func readInput(path string) string {
    bytes, err := ioutil.ReadFile("input")
    if err != nil {
        fmt.Printf("Error reading %s: %s", path, err)
        os.Exit(1)
    }
    return strings.TrimSpace(string(bytes))
}
