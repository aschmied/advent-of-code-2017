package main

import (
    "reflect"
    "testing"

    "github.com/aschmied/assert"
)

func TestLargestUnique(t *testing.T) {
    banks := NewMemoryBanksController([]BankSize{2, 6, 2, 0})
    assert.Int(t, banks.Largest(), 1)
}

func TestLargestNonUnique(t *testing.T) {
    banks := NewMemoryBanksController([]BankSize{2, 1, 2, 0})
    assert.Int(t, banks.Largest(), 0)
}

func TestRedistribute(t *testing.T) {
    banks := NewMemoryBanksController([]BankSize{3, 0, 0, 0, 0})
    banks.Redistribute(0)
    assertBanks(t, banks.Banks(), []BankSize{0, 1, 1, 1, 0})
}

func TestRedistributeWrap(t *testing.T) {
    banks := NewMemoryBanksController([]BankSize{0, 0, 0, 3, 0})
    banks.Redistribute(3)
    assertBanks(t, banks.Banks(), []BankSize{1, 1, 0, 0, 1})
}

func TestRedistributeWrapMultiple(t *testing.T) {
    banks := NewMemoryBanksController([]BankSize{0, 10, 0})
    banks.Redistribute(1)
    assertBanks(t, banks.Banks(), []BankSize{3, 3, 4})
}

func TestKey(t *testing.T) {
    banks := NewMemoryBanksController([]BankSize{0, 1, 999})
    assert.String(t, banks.Key(), "0,1,999,")
}

func assertBanks(t *testing.T, actual []BankSize, expected []BankSize) {
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Expected %v but got %v", expected, actual)
    }
}
