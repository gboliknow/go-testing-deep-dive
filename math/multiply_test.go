package math

import (
	"testing"
	"time"
)

func TestMutiply_Basic(t *testing.T) {
	got := Multiply(4, 5)
	want := 20

	if got != want {
		t.Error("Mutiply(4, 5) failed: expected", want, "got", got)
	}
}
func TestMultiply_Negative(t *testing.T) {
	got := Multiply(-2, 3)
	want := -6
	if got != want {
		t.Errorf("Multiply(-2, 3) failed: expected %d, got %d", want, got)
	}
}

// TestMultiply_ZeroInput uses t.Fatal for a critical failure
func TestMultiply_ZeroInput(t *testing.T) {
	got := Multiply(10, 0)
	want := 0
	if got != want {
		t.Fatal("Multiply(10, 0) failed: expected 0")
	}
}

func TestMultiply_LargeNumbers(t *testing.T) {
	got := Multiply(1000, 1000)
	want := 1000000
	if got != want {
		t.Fatalf("Multiply(1000, 1000) failed: expected %d, got %d", want, got)
	}
}

// TestMultiply_NotFullyTested uses t.Skip to skip an incomplete test
func TestMultiply_NotFullyTested(t *testing.T) {
	t.Skip("Multiply overflow cases not yet tested")
}

// TestMultiply_Parallel uses t.Parallel for concurrent execution
func TestMultiply_Parallel(t *testing.T) {
	t.Parallel()
	got := Multiply(2, 3)
	want := 6
	if got != want {
		t.Errorf("Multiply(2, 3) failed: expected %d, got %d", want, got)
	}
}

func TestMultiply_Slow(t *testing.T) {
	t.Parallel()
	time.Sleep(500 * time.Millisecond)
	got := Multiply(3, 3)
	want := 9
	if got != want {
		t.Errorf("Multiply(2, 3) failed: expected %d, got %d", want, got)
	}
}
