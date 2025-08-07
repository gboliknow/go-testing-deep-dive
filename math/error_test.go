package math

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// equalSlices compares two int slices for equality
func equalSlices(t *testing.T, a, b []int) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// equalMaps compares two string-int maps for equality
func equalMaps(t *testing.T, a, b map[string]int) bool {
	t.Helper()
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if bV, ok := b[k]; !ok || bV != v {
			return false
		}
	}
	return true
}

// TestMultiplySlice tests slice equality and error handling
func TestMultiplySlice(t *testing.T) {
	tests := []struct {
		name       string
		numbers    []int
		configData string
		expected   []int
		wantErr    bool
		wantErrMsg string
		wantErrIs  string
	}{
		{
			name:       "ValidInput",
			numbers:    []int{1, 2, 3},
			configData: "2",
			expected:   []int{2, 4, 6},
			wantErr:    false,
		},
		{
			name:       "EmptySlice",
			numbers:    []int{},
			configData: "2",
			expected:   nil,
			wantErr:    true,
			wantErrMsg: "invalid input: 0",
			wantErrIs:  "InvalidInputError",
		},
		{
			name:       "InvalidConfig",
			numbers:    []int{1, 2},
			configData: "invalid",
			expected:   nil,
			wantErr:    true,
			wantErrMsg: "parse multiplier: strconv.Atoi: parsing \"invalid\": invalid syntax",
			wantErrIs:  "ConfigError",
		},
		{
			name:       "NegativeMultiplier",
			numbers:    []int{1, 2},
			configData: "-2",
			expected:   []int{-2, -4},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPath := filepath.Join(t.TempDir(), "config.txt")
			if tt.configData != "" {
				if err := os.WriteFile(configPath, []byte(tt.configData), 0644); err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}
			}

			got, err := MultiplySlice(tt.numbers, configPath)

			if (err != nil) != tt.wantErr {
				t.Fatalf("MultiplySlice(%v, %q) error = %v, wantErr %v", tt.numbers, configPath, err, tt.wantErr)
			}
			if tt.wantErr {
				switch tt.wantErrIs {
				case "InvalidInputError":
					var target *InvalidInputError
					if !errors.As(err, &target) {
						t.Errorf("MultiplySlice(%v, %q) error type = %T, want *InvalidInputError", tt.numbers, configPath, err)
					}
				case "ConfigError":
					var target *ConfigError
					if !errors.As(err, &target) {
						t.Errorf("MultiplySlice(%v, %q) error type = %T, want *ConfigError", tt.numbers, configPath, err)
					}
				}
				if !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("MultiplySlice(%v, %q) error message = %q, want to contain %q", tt.numbers, configPath, err.Error(), tt.wantErrMsg)
				}
			}

			if !equalSlices(t, got, tt.expected) {
				t.Errorf("MultiplySlice(%v, %q) = %v, want %v", tt.numbers, configPath, got, tt.expected)
			}
		})
	}
}

// TestMultiplyMap tests map equality and error handling
func TestMultiplyMap(t *testing.T) {
	tests := []struct {
		name       string
		values     map[string]int
		configData string
		expected   map[string]int
		wantErr    bool
		wantErrMsg string
		wantErrIs  string
	}{
		{
			name:       "ValidInput",
			values:     map[string]int{"a": 1, "b": 2},
			configData: "3",
			expected:   map[string]int{"a": 3, "b": 6},
			wantErr:    false,
		},
		{
			name:       "EmptyMap",
			values:     map[string]int{},
			configData: "2",
			expected:   nil,
			wantErr:    true,
			wantErrMsg: "invalid input: 0",
			wantErrIs:  "InvalidInputError",
		},
		{
			name:       "InvalidConfig",
			values:     map[string]int{"a": 1},
			configData: "invalid",
			expected:   nil,
			wantErr:    true,
			wantErrMsg: "parse multiplier: strconv.Atoi: parsing \"invalid\": invalid syntax",
			wantErrIs:  "ConfigError",
		},
		{
			name:       "NegativeMultiplier",
			values:     map[string]int{"a": 1, "b": 2},
			configData: "-2",
			expected:   map[string]int{"a": -2, "b": -4},
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPath := filepath.Join(t.TempDir(), "config.txt")
			if tt.configData != "" {
				if err := os.WriteFile(configPath, []byte(tt.configData), 0644); err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}
			}

			got, err := MultiplyMap(tt.values, configPath)

			if (err != nil) != tt.wantErr {
				t.Fatalf("MultiplyMap(%v, %q) error = %v, wantErr %v", tt.values, configPath, err, tt.wantErr)
			}
			if tt.wantErr {
				switch tt.wantErrIs {
				case "InvalidInputError":
					var target *InvalidInputError
					if !errors.As(err, &target) {
						t.Errorf("MultiplyMap(%v, %q) error type = %T, want *InvalidInputError", tt.values, configPath, err)
					}
				case "ConfigError":
					var target *ConfigError
					if !errors.As(err, &target) {
						t.Errorf("MultiplyMap(%v, %q) error type = %T, want *ConfigError", tt.values, configPath, err)
					}
				}
				if !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Errorf("MultiplyMap(%v, %q) error message = %q, want to contain %q", tt.values, configPath, err.Error(), tt.wantErrMsg)
				}
			}

			if !equalMaps(t, got, tt.expected) {
				t.Errorf("MultiplyMap(%v, %q) = %v, want %v", tt.values, configPath, got, tt.expected)
			}
		})
	}
}