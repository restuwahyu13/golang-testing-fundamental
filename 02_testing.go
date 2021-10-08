package main

import "fmt"

type MathBasicx struct {
	X, Y int
}

func Additionx(x, y int) int {
	data := MathBasicx{X: x, Y: y}
	result := data.X + data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}

func Substractx(x, y int) int {
	data := MathBasicx{X: x, Y: y}
	result := data.X - data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}

func Dividex(x, y int) int {
	data := MathBasicx{X: x, Y: y}
	result := data.X / data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}

func Multiplicationx(x, y int) int {
	data := MathBasicx{X: x, Y: y}
	result := data.X * data.Y
	defer fmt.Sprintf("your result is %d", result)
	return result
}
