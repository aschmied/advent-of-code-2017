package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    input := readInput("input")
    sheet := ParseSheet(input)
    checksum := CalculateChecksumForSheet(sheet)
    fmt.Println(checksum)
}

func readInput(path string) string {
    bytes, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatalf("Error reading %s: %s\n", path, err)
    }
    
    return string(bytes)
}
