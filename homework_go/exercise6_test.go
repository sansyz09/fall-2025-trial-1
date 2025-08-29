package main

import "testing"

func TestFactorial(t *testing.T) {
    result := Factorial(5)
    expected := 120
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
    
    result = Factorial(0)
    expected = 1
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}