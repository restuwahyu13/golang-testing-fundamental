package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstanceOf(tt *testing.T) {

	tt.Run("Should be instanceof is integer", func(t *testing.T) {
		ctx := InstanceOf(2021)
		assert.Equal(t, ctx, true)
	})

	tt.Run("Should be instanceof value is boolean", func(t *testing.T) {
		ctx := InstanceOf(true)
		assert.Equal(t, ctx, true)
	})

	tt.Run("Should be instanceof value is string", func(t *testing.T) {
		ctx := InstanceOf("hello world")
		assert.Equal(t, ctx, true)
	})
}
