package main

import (
    "reflect"
    "testing"
)

func TestProgram(t *testing.T) {
    program := Program{0, 1, 2, 3}
    assertProgram(t, NewJumpInterpreter(program).Program(), program)
}

func TestExecuteJumpBackward(t *testing.T) {
    interpreter := NewJumpInterpreter(Program{-1, 0})
    interpreter.Execute()
    assertInt(t, interpreter.Cycles(), 1)
    assertProgram(t, interpreter.Program(), Program{0, 0})
}

func TestExecuteJumpForward(t *testing.T) {
    interpreter := NewJumpInterpreter(Program{1, 1})
    interpreter.Execute()
    assertInt(t, interpreter.Cycles(), 2)
    assertProgram(t, interpreter.Program(), Program{2, 2})
}

func TestComplexProgram(t *testing.T) {
    interpreter := NewJumpInterpreter(Program{0, -1, -2, -999})
    interpreter.Execute()
    assertInt(t, interpreter.Cycles(), 7)
    assertProgram(t, interpreter.Program(), Program{4, 0, -1, -998})
}

func assertInt(t *testing.T, actual int, expected int) {
    if expected != actual {
        t.Errorf("Expected %d but got %d", expected, actual)
    }
}

func assertProgram(t *testing.T, actual Program, expected Program) {
    if !reflect.DeepEqual(expected, actual) {
        t.Errorf("Expected %v but got %v", expected, actual)
    }
}
