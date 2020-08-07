package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expectedResult := 3
	actualResult := add(1, 2)
	if actualResult != expectedResult {
		t.Errorf("Called add(1, 2) expected %d got %d", expectedResult, actualResult)
	}
}
