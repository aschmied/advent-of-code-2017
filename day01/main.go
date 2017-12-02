package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    input := readInput("input")
    fmt.Println(SumIntsThatRepeat(input))
}

func readInput(path string) string {
    bytes, err := ioutil.ReadFile("input")
    if err != nil {
        fmt.Printf("Error reading %s: %s", path, err)
        os.Exit(1)
    }
    return string(bytes)
}
