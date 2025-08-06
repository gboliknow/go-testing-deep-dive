package math

import "testing"

type safeMultiplyTestCase struct {
	name       string
	a, b       int
	expected   int
	wantErr    bool
	wantErrMsg string
}

func TestSafeMultiply_NamedStruct(t *testing.T) {
	tests := []safeMultiplyTestCase{
		{
			name:     "PositiveNumbers",
			a:        4,
			b:        5,
			expected: 20,
			wantErr:  false,
		},
		{
			name:     "NegativeNumbers",
			a:        -2,
			b:        3,
			expected: -6,
			wantErr:  false,
		},
		{
			name:     "ZeroInput",
			a:        10,
			b:        0,
			expected: 0,
			wantErr:  false,
		},
		// Optional: you could add an edge case like multiplying large numbers without expecting overflow
		{
			name:     "LargeNumbersNoOverflow",
			a:        2147483647,
			b:        2,
			expected: 4294967294,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("Starting test: %s", tt.name)
			got, err := SafeMultiply(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("SafeMultiply(%d, %d) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.wantErrMsg {
				t.Errorf("SafeMultiply(%d, %d) error message = %q, want %q", tt.a, tt.b, err.Error(), tt.wantErrMsg)
			}
			if got != tt.expected {
				t.Errorf("SafeMultiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
			t.Logf("Finished test: %s", tt.name)
		})
	}
}

func TestDivide_AnonymousStruct(t *testing.T) {
	tests := []struct {
		name       string
		a, b       int
		expected   int
		wantErr    bool
		wantErrMsg string
	}{
		{"DivideByTwo", 10, 2, 5, false, ""},
		{"DivideByZero", 10, 0, 0, true, "division by zero"},
		{"NegativeResult", -10, 2, -5, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("Starting test: %s", tt.name)
			got, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Divide(%d, %d) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.wantErrMsg {
				t.Errorf("Divide(%d, %d) error message = %q, want %q", tt.a, tt.b, err.Error(), tt.wantErrMsg)
			}
			if got != tt.expected {
				t.Errorf("Divide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
			t.Logf("Finished test: %s", tt.name)
		})
	}
}
