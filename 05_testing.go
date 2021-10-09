package main

import "math"

type CalculateMeta struct {
	X, Y float64
}

func Calculate(input *CalculateMeta) float64 {
	var schema CalculateMeta
	schema.X = input.X
	schema.Y = input.Y

	return math.Pow(schema.X, schema.Y)
}
