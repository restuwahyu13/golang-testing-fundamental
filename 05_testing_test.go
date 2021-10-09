package main

import (
	. "github.com/onsi/ginkgo" // for global declaration bdd test funcinonality like describe, afterEach or more
	. "github.com/onsi/gomega" // for global declaration bdd test funcinonality like expect, equal or more
)

var _ = Describe("Calculate", func() {
	var data CalculateMeta

	// run this function before all every test is calling
	BeforeSuite(func() {
		data.X = 2
		data.Y = 2
	})

	// run this function if every test is done
	AfterEach(func() {
		data.X = 0
		data.Y = 0
	})

	Describe("Group All Test Calculate", func() {

		It("Should be calculate value using to equal", func() {
			res := Calculate(&data)
			Expect(res).To(Equal(float64(4)))
		})

		It("Should be calculate value using not equal", func() {
			res := Calculate(&data)
			Expect(res).NotTo(Equal(float64(10)))
		})

		It("Should be calculate value overwrite old data", func() {
			data.X = 10
			data.Y = 2

			res := Calculate(&data)
			Expect(res).To(Equal(float64(100)))
		})

		It("Should be calculate value using not equal", func() {
			Skip("this test not execute because, this test is skip")

			res := Calculate(&data)
			Expect(res).NotTo(Equal(float64(10)))
		})
	})
})
