package math

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ConfigError is a custom error for config issues
type ConfigError struct {
	Path string
	Err  error
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("config error at %s: %v", e.Path, e.Err)
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

// InvalidInputError is a custom error for invalid inputs
type InvalidInputError struct {
	Value int
}

func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %d", e.Value)
}

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func IsPositive(a int) bool {
	return a > 0
}

func Multiply(a, b int) int {
	return a * b
}

func SafeMultiply(a, b int) (int, error) {
	result := a * b
	if a != 0 && result/a != b {
		return 0, fmt.Errorf("multiplication overflow")
	}
	return result, nil
}

func SafeMultiply2(a, b int) (int, error) {
	if b == 0 {
		return 0, nil
	}
	if a > math.MaxInt32/b || a < math.MinInt32/b {
		return 0, fmt.Errorf("multiplication overflow")
	}
	result := a * b
	return result, nil
}

func MultiplyWithConfig(a int, configPath string) (int, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return 0, &ConfigError{Path: configPath, Err: fmt.Errorf("read config: %w", err)}
	}
	multiplier, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, &ConfigError{Path: configPath, Err: fmt.Errorf("parse multiplier: %w", err)}
	}
	return a * multiplier, nil
}

func MultiplySlice(numbers []int, configPath string) ([]int, error) {
	if len(numbers) == 0 {
		return nil, &InvalidInputError{Value: 0}
	}
	multiplier, err := MultiplyWithConfig(1, configPath)
	if err != nil {
		return nil, fmt.Errorf("get multiplier: %w", err)
	}
	result := make([]int, len(numbers))
	for i, n := range numbers {
		if multiplier > 0 {
			if n > math.MaxInt32/multiplier || n < math.MinInt32/multiplier {
				return nil, fmt.Errorf("overflow at index %d: %w", i, &InvalidInputError{Value: n})
			}
		} else if multiplier < 0 {
			if n < math.MaxInt32/multiplier || n > math.MinInt32/multiplier {
				return nil, fmt.Errorf("overflow at index %d: %w", i, &InvalidInputError{Value: n})
			}
		}
		result[i] = n * multiplier
	}
	return result, nil
}

func MultiplyMap(values map[string]int, configPath string) (map[string]int, error) {
	if len(values) == 0 {
		return nil, &InvalidInputError{Value: 0}
	}
	multiplier, err := MultiplyWithConfig(1, configPath)
	if err != nil {
		return nil, fmt.Errorf("get multiplier: %w", err)
	}
	result := make(map[string]int)
	for k, v := range values {
		if multiplier > 0 {
			if v > math.MaxInt32/multiplier || v < math.MinInt32/multiplier {
				return nil, fmt.Errorf("overflow for key %s: %w", k, &InvalidInputError{Value: v})
			}
		} else if multiplier < 0 {
			if v < math.MaxInt32/multiplier || v > math.MinInt32/multiplier {
				return nil, fmt.Errorf("overflow for key %s: %w", k, &InvalidInputError{Value: v})
			}
		}
		result[k] = v * multiplier
	}
	return result, nil
}
