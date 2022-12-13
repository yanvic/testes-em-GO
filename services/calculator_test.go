package services_test

import (
	"projeto go/services"
	"testing"
)

func TestSum(t *testing.T) {
	if services.Sum(2, 2) != 4 {
		t.Error("2 + 2 must be 4")
	}
}

func TestMulti(t *testing.T) {
	if services.Multiply(2, 2) != 4 {
		t.Error("2 * 2 must be 4")
	}
}
