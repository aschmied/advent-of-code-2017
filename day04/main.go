package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    lines := readInput("input")
    numberValidPassphrases := 0
    for _, line := range lines {
        passphrase := NewPassphrase(line)
        if !passphrase.HasDuplicateWords() {
            numberValidPassphrases += 1
        }
    }
    fmt.Println(numberValidPassphrases)
}

func readInput(filename string) []string {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Errorf("Error reading %v: %v", filename, err)
    }
    text := string(bytes)
    return strings.Split(strings.TrimSpace(text), "\n")
}
