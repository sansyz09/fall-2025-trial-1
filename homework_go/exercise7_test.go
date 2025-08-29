package main

import "testing"

func TestSumArray(t *testing.T) {
    result := SumArray([]int{1, 2, 3, 4, 5})
    expected := 15
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
    
    result = SumArray([]int{})
    expected = 0
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}