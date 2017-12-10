package main

type stringSet struct {
    values map[string]bool
}

func NewStringSet() stringSet {
    return stringSet{make(map[string]bool)}
}

func (self *stringSet) Add(string string) {
    self.values[string] = true
}

func (self *stringSet) Contains(string string) bool {
    _, ok := self.values[string]
    return ok
}

func (self *stringSet) Size() int {
    return len(self.values)
}
