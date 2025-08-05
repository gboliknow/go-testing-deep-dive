package math_test

import (
    "example.com/testing/math"
    "testing"
)

func TestAdd_External(t *testing.T) {
    got := math.Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("math.Add(2, 3) failed: expected %d, got %d", want, got)
    }
}