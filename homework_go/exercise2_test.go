package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(5, 3)
    expected := 8
    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}