package testutil

import (
    "math"
    "testing"
)

// EqualFloat checks if two floats are equal within a tolerance.
func EqualFloat(t *testing.T, got, want, tolerance float64) {
    t.Helper()
    if math.Abs(got-want) > tolerance {
        t.Errorf("got %v, want %v (tolerance %v)", got, want, tolerance)
    }
}

// MustFail runs a function expecting an error.
func MustFail(t *testing.T, fn func() error) {
    t.Helper()
    if err := fn(); err == nil {
        t.Errorf("expected error but got nil")
    }
}
