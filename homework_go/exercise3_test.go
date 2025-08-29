package main

import "testing"

func TestIsEven(t *testing.T) {
    if !IsEven(4) {
        t.Error("4 should be even")
    }
    if IsEven(5) {
        t.Error("5 should be odd")
    }
}