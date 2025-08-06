package math

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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
		return 0, fmt.Errorf("read config: %w", err)
	}
	multiplier, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, fmt.Errorf("parse multiplier: %w", err)
	}
	return a * multiplier, nil
}
