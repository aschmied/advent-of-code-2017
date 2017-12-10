package main

import "strings"

type passphrase struct {
    words []string
}

func NewPassphrase(string string) passphrase {
    return passphrase{strings.Fields(string)}
}

func (self passphrase) Words() []string {
    ret := make([]string, len(self.words))
    copy(ret, self.words)
    return ret
}

func (self passphrase) Valid() bool {
    usedWords := NewStringSet()
    for _, word := range self.words {
        if usedWords.Contains(word) {
            return false
        }
        usedWords.Add(word)
    }
    return true
}
