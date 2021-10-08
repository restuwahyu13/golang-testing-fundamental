package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ctx := Addition(5, 5)

	if ctx != 10 {
		t.Errorf("Test is fail: %d", ctx)
	}
}

func TestSubtract(t *testing.T) {
	ctx := Substract(5, 2)

	if ctx != 3 {
		t.Errorf("Test is fail: %d", ctx)
	}
}

func TestDivide(t *testing.T) {
	ctx := Divide(10, 2)

	if ctx != 5 {
		t.Errorf("Test is fail: %d", ctx)
	}
}

func TestMultiplication(t *testing.T) {
	ctx := Multiplication(10, 2)

	if ctx != 20 {
		t.Errorf("Test is fail: %d", ctx)
	}
}
