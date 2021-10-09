package main

import (
	"reflect"
	"testing"
)

func TestAddition(tt *testing.T) {

	tt.Run("Should be addition value success", func(t *testing.T) {
		ctx := Addition(5, 5)

		if ctx != 10 {
			t.FailNow()
		}
	})

	tt.Run("Should be addition value is integer", func(t *testing.T) {
		ctx := Addition(5, 5)

		if reflect.TypeOf(ctx) == reflect.TypeOf("Hello Wordl") {
			t.FailNow()
		}
	})
}

func TestSubtract(tt *testing.T) {

	tt.Run("Should be substract value success", func(t *testing.T) {
		ctx := Substract(5, 2)

		if ctx != 3 {
			t.FailNow()
		}
	})

	tt.Run("Should be substract value is integer", func(t *testing.T) {
		ctx := Substract(5, 2)

		if reflect.TypeOf(ctx) == reflect.TypeOf("Hello Wordl") {
			t.FailNow()
		}
	})
}

func TestDivide(tt *testing.T) {

	tt.Run("Should be divide value success", func(t *testing.T) {
		ctx := Divide(10, 2)

		if ctx != 5 {
			t.FailNow()
		}
	})

	tt.Run("Should be divide value is integer", func(t *testing.T) {
		ctx := Divide(10, 2)

		if reflect.TypeOf(ctx) == reflect.TypeOf("Hello Wordl") {
			t.FailNow()
		}
	})
}

func TestMultiplication(tt *testing.T) {

	tt.Run("Should be multiplication value success", func(t *testing.T) {
		ctx := Multiplication(10, 2)

		if ctx != 20 {
			t.FailNow()
		}
	})

	tt.Run("Should be multiplication value is integer", func(t *testing.T) {
		ctx := Multiplication(10, 2)

		if reflect.TypeOf(ctx) == reflect.TypeOf("Hello Wordl") {
			t.FailNow()
		}
	})
}
