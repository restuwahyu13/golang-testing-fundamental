package main

import (
	"reflect"

	"golang.org/x/text/date"
)

func InstanceOf(data interface{}) bool {

	switch reflect.TypeOf(data) {
	case reflect.TypeOf("Hello World"):
		return true
	case reflect.TypeOf(date.Indochina):
		return true
	case reflect.TypeOf(true):
		return true
	default:
		return false
	}
}
