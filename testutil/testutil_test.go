package testutil

import (
	"testing"
)

func TestFloatEqual(t *testing.T) {
	if FloatEqE10(float64(3.14), float64(3.141)) == true {
		t.Errorf("error")
	}
	if FloatEqE10(float64(3.14), float64(3.14)) == false {
		t.Errorf("error")
	}
}
