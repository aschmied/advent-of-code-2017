package main

import "testing"

func TestEmptyString(t *testing.T) {
    s := NewStringSet()
    assertSize(t, s, 0)
    s.Add("")
    assertSize(t, s, 1)
}

func TestNonEmptyStrings(t *testing.T) {
    s := NewStringSet()
    s.Add("foo")
    assertSize(t, s, 1)
    s.Add("bar")
    assertSize(t, s, 2)
    s.Add("bar")
    assertSize(t, s, 2)

    assertContains(t, s, "foo")
    assertContains(t, s, "bar")
    assertNotContains(t, s, "")
    assertNotContains(t, s, "baz")
}

func assertSize(t *testing.T, set stringSet, expected int) {
    actual := set.Size()
    if actual != expected {
        t.Errorf("Incorrect size of stringSet. Expected %v but got %v", expected, actual)
    }
}

func assertContains(t *testing.T, set stringSet, string string) {
    contains := set.Contains(string)
    if !contains {
        t.Errorf("Expected stringSet to contain %v but it did not", string)
    }
}

func assertNotContains(t *testing.T, set stringSet, string string) {
    contains := set.Contains(string)
    if contains {
        t.Errorf("Expected stringSet not to contain %v but it does", string)
    }
}
