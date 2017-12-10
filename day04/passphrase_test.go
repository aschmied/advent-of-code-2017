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

func TestHasDuplicateWords(t *testing.T) {
    assertHasNoDuplicateWords(t, NewPassphrase(""))
    assertHasNoDuplicateWords(t, NewPassphrase("the"))
    assertHasNoDuplicateWords(t, NewPassphrase("the quick brown fox"))
    assertHasDuplicateWords(t, NewPassphrase("the quick brown fox jumps over the"))
    assertHasDuplicateWords(t, NewPassphrase("the the"))
}

func assertHasNoDuplicateWords(t *testing.T, passphrase passphrase) {
    assertHasDuplicateWordsValue(t, passphrase, false)
}

func assertHasDuplicateWords(t *testing.T, passphrase passphrase) {
    assertHasDuplicateWordsValue(t, passphrase, true)
}

func assertHasDuplicateWordsValue(t *testing.T, passphrase passphrase, expected bool) {
    actual := passphrase.HasDuplicateWords()
    if actual != expected {
        t.Errorf("Incorrect value for HasDuplicateWords(). Expected %v but got %v", expected, actual)
    }
}

func TestHasAnnagrams(t *testing.T) {
    assertHasNoAnnagrams(t, NewPassphrase(""))
    assertHasNoAnnagrams(t, NewPassphrase("the"))
    assertHasNoAnnagrams(t, NewPassphrase("the quick brown fox"))
    assertHasAnnagrams(t, NewPassphrase("the quick brown ckuiq fox"))
}

func TestHasAnnagramsSubsets(t *testing.T) {
    assertHasNoAnnagrams(t, NewPassphrase("quick quicker"))
}

func assertHasAnnagrams(t *testing.T, passphrase passphrase) {
    assertHasAnnagramsValue(t, passphrase, true)
}

func assertHasNoAnnagrams(t *testing.T, passphrase passphrase) {
    assertHasAnnagramsValue(t, passphrase, false)
}

func assertHasAnnagramsValue(t *testing.T, passphrase passphrase, expected bool) {
    actual := passphrase.HasAnnagrams()
    if actual != expected {
        t.Errorf("Incorrect value for HasAnnagrams(). Expected %v but got %v", expected, actual)
    }
}
