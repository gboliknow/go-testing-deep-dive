package math

import (
	"runtime"
	"testing"
)

func TestAdd_Basic(t *testing.T) {
	got := Add(9, 3)
	want := 12

	if got != want {
		t.Error("Add( 2, 3)  failed: expected ", want, "got", got)
	}
}

func TestSubtract_basic(t *testing.T) {
	got := Subtract(5, 3)
	want := 2
	if got != want {
		t.Errorf("Subtract(5,3) failed : expected %d , got %d", want, got)
	}
}

func TestDivide_ByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Fatal("Divide (10, 0) should return an error")
	}
}

func TestDivide_Basic(t *testing.T) {
	got, err := Divide(8, 2)
	want := 4

	if err != nil {
		t.Fatalf("Divide (8, 2) failed with error %v", err)
	}
	if got != want {
		t.Fatalf("Divide (8,2 ) failed : expected %d , got %d", want, got)
	}
}

func TestMutiply_NotImplemented(t *testing.T) {
	t.Skip("Multiply is not yet implemented")
}

func TestAdd_Parallel(t *testing.T) {
	t.Parallel()

	got := Add(1, 1)
	want := 2
	if got != want {
		t.Errorf("Add(1,1) failed : expected %d , got %d", want, got)
	}

}

// Scenario 7: Multiple Assertions (t.Error)
func TestAdd_Zero(t *testing.T) {
	got := Add(0, 0)
	want := 0
	if got != want {
		t.Error("Add(0, 0) failed: expected", want, "got", got)
	}
	if !IsPositive(got) && got != 0 {
		t.Error("Add(0, 0) should return a non-negative number")
	}
}

// Scenario 8: Edge Case (t.Errorf)
func TestSubtract_Negative(t *testing.T) {
	got := Subtract(-1, -1)
	want := 0
	if got != want {
		t.Errorf("Subtract(-1, -1) failed: expected %d, got %d", want, got)
	}
}

// Scenario 9: Invalid Input (t.Fatal)
func TestIsPositive_Negative(t *testing.T) {
	if IsPositive(-5) {
		t.Fatal("IsPositive(-5) should return false")
	}
}

// Scenario 10: Skip Based on Condition (t.Skip)
func TestAdd_PlatformSpecific(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on Windows")
	}
	got := Add(10, 10)
	want := 20
	if got != want {
		t.Errorf("Add(10, 10) failed: expected %d, got %d", want, got)
	}
}

// Scenario 11: Parallel Edge Case (t.Parallel)
func TestAdd_LargeNumbers(t *testing.T) {
	t.Parallel()
	got := Add(1000000, 2000000)
	want := 3000000
	if got != want {
		t.Errorf("Add(1000000, 2000000) failed: expected %d, got %d", want, got)
	}
}

// Scenario 12: Mixed Use (t.Error + t.Fatal)
func TestDivide_Mixed(t *testing.T) {
	got, err := Divide(6, 2)
	want := 3
	if err != nil {
		t.Fatal("Divide(6, 2) failed with unexpected error:", err)
	}
	if got != want {
		t.Error("Divide(6, 2) failed: expected", want, "got", got)
	}
}

//Add a table-driven test to math_test.go to practice naming and subtests:

func TestAdd_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{name: "PositiveNumbers", a: 2, b: 3, expected: 5},
		{name: "NegativeNumbers", a: -1, b: -1, expected: -2},
		{name: "ZeroInput", a: 0, b: 0, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add (%d, %d) failed :expected  %d , got %d", tt.a, tt.b, tt.expected, got)
			}
		})
	}

}

