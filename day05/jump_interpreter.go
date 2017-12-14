package main

type Instruction int

type Program []Instruction

type instructionMutator func (Instruction) Instruction

type jumpInterpreter struct {
    program Program
    cycles int
    mutator instructionMutator
}

func NewJumpInterpreter(program Program) jumpInterpreter {
    mutator := func (instruction Instruction) Instruction {
        return instruction + 1
    }
    return jumpInterpreter{program, 0, mutator}
}

func (self jumpInterpreter) Program() Program {
    return self.program
}

func (self jumpInterpreter) Cycles() int {
    return self.cycles
}

func (self *jumpInterpreter) Execute() {
    index := 0
    for index >= 0 && index < len(self.program) {
        currentInstruction := self.program[index]
        self.program[index] = self.mutator(currentInstruction)
        index += int(currentInstruction)
        self.cycles += 1
    }
}
