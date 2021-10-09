package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCalculate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculate Suite")
}
