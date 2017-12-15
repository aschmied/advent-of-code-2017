package main

import (
    "fmt"
    "github.com/aschmied/set"
)

func main() {
    input := []BankSize{14, 0, 15, 12, 11, 11, 3, 5, 1, 6, 8, 4, 9, 1, 8, 4}
    banks := NewMemoryBanksController(input)

    stepsToRepeatAState := 0
    states := set.Strings()
    for !states.Contains(banks.Key()) {
        states.Add(banks.Key())
        banks.Redistribute(banks.Largest())
         stepsToRepeatAState++
    }

    targetState := banks.Key()
    cycleLength := 1
    banks.Redistribute(banks.Largest())
    for targetState != banks.Key() {
        cycleLength += 1
        banks.Redistribute(banks.Largest())
    }

    fmt.Println(stepsToRepeatAState)
    fmt.Println(cycleLength)
}
