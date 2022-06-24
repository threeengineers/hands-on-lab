package feature_test

import (
	"testing"

	"github.com/threeengineers/hands-on-lab/02-unit-testing/feature"
)

func TestCalculatorAddShouldReturnSuccess(t *testing.T) {
	calculator := new(feature.Calculator)
	result := calculator.Add(1, 2)

	if result != 3 {
		t.Errorf("Expected %d, got %d", 3, result)
	}
}

func TestCalculatorSubtractShouldReturnSuccess(t *testing.T) {
	calculator := new(feature.Calculator)
	result := calculator.Subtract(1, 2)

	if result != -1 {
		t.Errorf("Expected %d, got %d", -1, result)
	}
}
