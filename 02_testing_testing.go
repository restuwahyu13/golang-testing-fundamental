package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddx(t *testing.T) {
	ctx := Addition(5, 5)
	assert.Equal(t, 10, ctx)
}

func TestSubtractx(t *testing.T) {
	ctx := Substract(5, 2)
	assert.Equal(t, 3, ctx)
}

func TestDividex(t *testing.T) {
	ctx := Divide(10, 2)
	assert.Equal(t, 5, ctx)
}

func TestMultiplicationx(t *testing.T) {
	ctx := Multiplication(10, 2)
	assert.Equal(t, 20, ctx)
}
