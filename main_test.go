package main

import (
	"testing"

	"github.com/nolka/goincrementor/types"
)

func TestIncrementorDefaultValue(t *testing.T) {
	inc := types.NewIncrementor()
	if inc.GetNumber() != 0 {
		t.Error("Default counter value must be equal to 0")
	}
}

func TestIncrementorZeroValue(t *testing.T) {
	inc := types.NewIncrementor()
	err := inc.SetMaxValue(0)
	if err != nil {
		t.Errorf("Zero value is allowed to use as max value, but got error: %s", err)
	}
}

func TestIncrementValueByOne(t *testing.T) {
	inc := types.NewIncrementor()

	inc.IncrementNumber()
	if inc.GetNumber() != 1 {
		t.Error("Value must be equal to 1")
	}
}

func TestIncrementMaxValueAndReset(t *testing.T) {
	inc := types.NewIncrementor()

	inc.SetMaxValue(2)

	if inc.GetNumber() != 0 {
		t.Error("Counter value must not be changed if SetMaxValue() called")
	}

	inc.IncrementNumber()
	if inc.GetNumber() != 1 {
		t.Error("Value must be equal to 1")
	}

	inc.IncrementNumber()
	inc.IncrementNumber()
	if inc.GetNumber() != 0 {
		t.Error("Value must reseted to 0 because max value is 2")
	}
}

func TestResetToZeroIfMaxValueLessThanCounterValue(t *testing.T) {
	inc := types.NewIncrementor()

	inc.IncrementNumber()
	inc.IncrementNumber()

	inc.SetMaxValue(1)
	inc.IncrementNumber()

	if inc.GetNumber() != 0 {
		t.Error("Counter value must be reseted to zero because MaxValue is less than current counter value")
	}
}
