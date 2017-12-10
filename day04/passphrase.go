package main

import (
    "sort"
    "strings"
)

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

func (self passphrase) HasDuplicateWords() bool {
    usedWords := NewStringSet()
    for _, word := range self.words {
        if usedWords.Contains(word) {
            return true
        }
        usedWords.Add(word)
    }
    return false
}

func (self passphrase) HasAnnagrams() bool {
    usedLettersets := NewStringSet()
    for _, word := range self.words {
        letterset := self.letterset(word)
        if usedLettersets.Contains(letterset) {
            return true
        }
        usedLettersets.Add(letterset)
    }
    return false
}

func (self passphrase) letterset(word string) string {
    letters := strings.Split(word, "")
    sort.Strings(letters)
    return strings.Join(letters, "")
}
