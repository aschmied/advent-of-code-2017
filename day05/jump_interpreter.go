package main

type Instruction int

type Program []Instruction

type jumpInterpreter struct {
    program Program
    cycles int
    programCounter int
}

func NewJumpInterpreter(program Program) jumpInterpreter {
    return jumpInterpreter{program, 0, 0}
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
        self.program[self.programCounter] += 1
        self.programCounter += int(currentInstruction)
        self.cycles += 1
    }
}
