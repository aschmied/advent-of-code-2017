package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func main() {
    program := readProgram("input")

    interpreter := NewJumpInterpreter(program)
    interpreter.Execute()

    upDownInterpreter := NewUpDownJumpInterpreter(program)
    upDownInterpreter.Execute()

    fmt.Println(interpreter.Cycles())
    fmt.Println(upDownInterpreter.Cycles())
}

func readProgram(filename string) Program {
    lines := readLines(filename)
    return parseLines(lines)
}

func readLines(filename string) []string {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Errorf("Error reading %v: %v", filename, err)
    }
    text := string(bytes)
    return strings.Split(strings.TrimSpace(text), "\n")
}

func parseLines(lines []string) Program {
    program := Program{}
    for _, line := range lines {
        instruction := parseInstruction(line)
        program = append(program, instruction)
    }
    return program
}

func parseInstruction(string string) Instruction {
    integer, err := strconv.Atoi(strings.TrimSpace(string))
    if err != nil {
        fmt.Errorf("Error parsing %v to int: %v", string, err)
    }
    return Instruction(integer)
}
