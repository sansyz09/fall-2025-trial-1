package main

import "testing"

func TestGreet(t *testing.T) {
    result := Greet("Alice")
    expected := "Hello, Alice!"
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}