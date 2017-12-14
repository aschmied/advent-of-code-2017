package main

type Instruction int

type Program []Instruction

type instructionMutator func (Instruction) Instruction

type jumpInterpreter struct {
    program Program
    cycles int
    programCounter int
    mutator instructionMutator
}

func NewJumpInterpreter(program Program) jumpInterpreter {
    mutator := func (instruction Instruction) Instruction {
        return instruction + 1
    }
    return jumpInterpreter{program, 0, 0, mutator}
}

func (self jumpInterpreter) Program() Program {
    return self.program
}

func (self jumpInterpreter) Cycles() int {
    return self.cycles
}

func (self *jumpInterpreter) Execute() {
    for self.programCounter >= 0 && self.programCounter < len(self.program) {
        currentInstruction := self.program[self.programCounter]
        self.program[self.programCounter] = self.mutator(currentInstruction)
        self.programCounter += int(currentInstruction)
        self.cycles += 1
    }
}
