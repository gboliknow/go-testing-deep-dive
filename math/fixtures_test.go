package math

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestMain sets up a package-level temporary directory
func TestMain(m *testing.M) {
	// Setup: Create a package-level temp directory
	tempDir, err := os.MkdirTemp("", "math-test")
	if err != nil {
		panic("failed to create temp dir: " + err.Error())
	}
	defer os.RemoveAll(tempDir) // Teardown: Clean up after all tests
	os.Exit(m.Run())
}

// createConfigFile is a helper to create a config file in t.TempDir()
func createConfigFile(t *testing.T, data string) string {
	t.Helper()
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.txt")
	if data != "" {
		if err := os.WriteFile(configPath, []byte(data), 0644); err != nil {
			t.Fatalf("failed to write config file: %v", err)
		}
	}
	return configPath
}

// assertError checks if an error matches the expected state and message
func assertError(t *testing.T, gotErr error, wantErr bool, wantErrMsg string) {
	t.Helper()
	if (gotErr != nil) != wantErr {
		t.Fatalf("error = %v, wantErr %v", gotErr, wantErr)
	}
	if wantErr && !strings.Contains(gotErr.Error(), wantErrMsg) {
		t.Errorf("error message = %q, want to contain %q", gotErr.Error(), wantErrMsg)
	}
}

// multiplyWithConfigTestCase is a test case struct
type multiplyWithConfigTestCase struct {
	name       string
	input      int
	configData string
	expected   int
	wantErr    bool
	wantErrMsg string
}

// newMultiplyWithConfigTestCase is a factory for test cases
func newMultiplyWithConfigTestCase(name, configData string, input, expected int, wantErr bool, wantErrMsg string) multiplyWithConfigTestCase {
	return multiplyWithConfigTestCase{
		name:       name,
		input:      input,
		configData: configData,
		expected:   expected,
		wantErr:    wantErr,
		wantErrMsg: wantErrMsg,
	}
}

func TestMultiplyWithConfig1(t *testing.T) {
	tests := []multiplyWithConfigTestCase{
		newMultiplyWithConfigTestCase("ValidConfig", "3", 5, 15, false, ""),
		newMultiplyWithConfigTestCase("InvalidConfig", "invalid", 5, 0, true, "parse multiplier: strconv.Atoi: parsing \"invalid\": invalid syntax"),
		newMultiplyWithConfigTestCase("MissingConfig", "", 5, 0, true, "read config: open "),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configPath := createConfigFile(t, tt.configData)
			got, err := MultiplyWithConfig(tt.input, configPath)
			assertError(t, err, tt.wantErr, tt.wantErrMsg)
			if got != tt.expected {
				t.Errorf("MultiplyWithConfig(%d, %q) = %d, want %d", tt.input, configPath, got, tt.expected)
			}
		})
	}
}

// TestMultiplyWithConfig uses t.TempDir() and per-test setup
func TestMultiplyWithConfig(t *testing.T) {
	tests := []struct {
		name       string
		input      int
		configData string
		expected   int
		wantErr    bool
	}{
		{
			name:       "ValidConfig",
			input:      5,
			configData: "3",
			expected:   15,
			wantErr:    false,
		},
		{
			name:       "InvalidConfig",
			input:      5,
			configData: "invalid",
			expected:   0,
			wantErr:    true,
		},
		{
			name:       "MissingConfig",
			input:      5,
			configData: "", // File doesn't exist
			expected:   0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Per-test setup: Create a temp directory and config file
			tempDir := t.TempDir()
			configPath := filepath.Join(tempDir, "config.txt")

			// Write config data if provided
			if tt.configData != "" {
				if err := os.WriteFile(configPath, []byte(tt.configData), 0644); err != nil {
					t.Fatalf("failed to write config file: %v", err)
				}
			}

			// Run the test
			got, err := MultiplyWithConfig(tt.input, configPath)
			if (err != nil) != tt.wantErr {
				t.Fatalf("MultiplyWithConfig(%d, %q) error = %v, wantErr %v", tt.input, configPath, err, tt.wantErr)
			}
			if got != tt.expected {
				t.Errorf("MultiplyWithConfig(%d, %q) = %d, want %d", tt.input, configPath, got, tt.expected)
			}

			// Teardown: t.TempDir() automatically cleans up
		})
	}
}
