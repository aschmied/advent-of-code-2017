package main

import "testing"

import "reflect"

func TestParse(t *testing.T) {
    assertWords(t, NewPassphrase(""), []string{})
    assertWords(t, NewPassphrase("the"), []string{"the"})
    assertWords(t, NewPassphrase("the quick brown fox jumps over the lazy dog"),
        []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"})
    assertWords(t, NewPassphrase("&3\t_ 88930(#&@^  \t 4 \n"), []string{"&3", "_", "88930(#&@^", "4"})
}

func TestImmutability(t *testing.T) {
    passphrase := NewPassphrase("foo bar")
    w := passphrase.Words()
    w[0] = "baz"
    assertWords(t, passphrase, []string{"foo", "bar"})
}

func assertWords(t *testing.T, passphrase passphrase, expected []string) {
    actual := passphrase.Words()
    if !reflect.DeepEqual(actual, expected) {
        t.Errorf("Incorrect value for Words(). Expected %v but got %v", expected, actual)
    }
}

func TestValid(t *testing.T) {
    assertValid(t, NewPassphrase(""))
    assertValid(t, NewPassphrase("the"))
    assertValid(t, NewPassphrase("the quick brown fox"))
    assertInalid(t, NewPassphrase("the quick brown fox jumps over the"))
    assertInalid(t, NewPassphrase("the the"))
}

func assertValid(t *testing.T, passphrase passphrase) {
    assertValidValue(t, passphrase, true)
}

func assertInalid(t *testing.T, passphrase passphrase) {
    assertValidValue(t, passphrase, false)
}

func assertValidValue(t *testing.T, passphrase passphrase, expected bool) {
    actual := passphrase.Valid()
    if actual != expected {
        t.Errorf("Incorrect value for Valid(). Expected %v but got %v", expected, actual)
    }
}
