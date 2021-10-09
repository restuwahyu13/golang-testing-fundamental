package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddx(tt *testing.T) {

	tt.Run("Should be addition value success", func(t *testing.T) {
		ctx := Addition(5, 5)
		assert.Equal(t, ctx, 10)
	})

	tt.Run("Should be addition value is integer", func(t *testing.T) {
		ctx := Addition(5, 5)
		assert.NotEqual(t, reflect.TypeOf(ctx), reflect.TypeOf("Hello Wordl"))
	})
}

func TestSubtractx(tt *testing.T) {

	tt.Run("Should be substract value success", func(t *testing.T) {
		ctx := Substract(5, 2)
		assert.Equal(t, ctx, 3)
	})

	tt.Run("Should be substract value is integer", func(t *testing.T) {
		ctx := Substract(5, 2)
		assert.NotEqual(t, reflect.TypeOf(ctx), reflect.TypeOf("Hello Wordl"))
	})
}

func TestDividex(tt *testing.T) {

	tt.Run("Should be divide value success", func(t *testing.T) {
		ctx := Divide(10, 2)
		assert.Equal(t, ctx, 3)
	})

	tt.Run("Should be divide value is integer", func(t *testing.T) {
		ctx := Divide(10, 2)
		assert.NotEqual(t, reflect.TypeOf(ctx), reflect.TypeOf("Hello Wordl"))
	})
}

func TestMultiplicationx(tt *testing.T) {

	tt.Run("Should be multiplication value success", func(t *testing.T) {
		ctx := Multiplication(10, 2)
		assert.Equal(t, 20, ctx)
	})

	tt.Run("Should be multiplication value is integer", func(t *testing.T) {
		ctx := Multiplication(10, 2)
		assert.NotEqual(t, reflect.TypeOf(ctx), reflect.TypeOf("Hello Wordl"))
	})
}
