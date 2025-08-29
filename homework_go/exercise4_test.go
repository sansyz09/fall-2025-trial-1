package main

import "testing"

func TestMax(t *testing.T) {
    result := Max(10, 5)
    expected := 10
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
    
    result = Max(3, 7)
    expected = 7
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}