package main

import "fmt"

type MathBasic struct {
	X, Y int
}

func Addition(x, y int) int {
	data := MathBasic{X: x, Y: y}
	result := data.X + data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}

func Substract(x, y int) int {
	data := MathBasic{X: x, Y: y}
	result := data.X - data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}

func Divide(x, y int) int {
	data := MathBasic{X: x, Y: y}
	result := data.X / data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}

func Multiplication(x, y int) int {
	data := MathBasic{X: x, Y: y}
	result := data.X * data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}
