package main

import (
    "bytes"
    "strconv"
)

type BankSize int

type memoryBanksController struct {
    banks []BankSize
}

func NewMemoryBanksController(banks []BankSize) memoryBanksController {
    return memoryBanksController{banks}
}

func (self memoryBanksController) Largest() int {
    indexOfLargest := 0
    sizeOfLargest := BankSize(0)
    for index, size := range self.banks {
        if size > sizeOfLargest {
            sizeOfLargest = size
            indexOfLargest = index
        }
    }
    return indexOfLargest
}

func (self memoryBanksController) Banks() []BankSize {
    return self.banks
}

func (self *memoryBanksController) Redistribute(bankIndex int) {
    numberToRedistribute := self.banks[bankIndex]
    self.banks[bankIndex] = 0

    numberToEachBank := int(numberToRedistribute) / len(self.banks)
    numberRemainder := int(numberToRedistribute) - (numberToEachBank * len(self.banks))

    for i := 0; i < len(self.banks); i++ {
        self.banks[i] += BankSize(numberToEachBank)
    }
    for numberRemainder > 0 {
        bankIndex = (bankIndex + 1) % len(self.banks)
        self.banks[bankIndex]++
        numberRemainder--
    }
}

func (self memoryBanksController) Key() string {
    var buffer bytes.Buffer
    for _, bank := range self.banks {
        buffer.WriteString(strconv.Itoa(int(bank)))
        buffer.WriteString(",")
    }
    return buffer.String()
}
