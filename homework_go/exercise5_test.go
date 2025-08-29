package main

import "testing"

func TestStringLength(t *testing.T) {
    result := StringLength("hello")
    expected := 5
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
    
    result = StringLength("")
    expected = 0
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}