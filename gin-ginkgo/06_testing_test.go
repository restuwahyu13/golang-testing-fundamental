package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/restuwahyu13/go-testing-fundamental/gin/utils"
)

var router = Router()

var _ = Describe("All Test", func() {
	Describe("Group All Test", func() {

		It("Should be get request success - first", func() {
			rr, req, err := utils.HttpRequest(http.MethodGet, "/first", nil)
			gin.SetMode(gin.TestMode)
			router.ServeHTTP(rr, req)

			if err != nil {
				Expect(err).To(MatchError(err))
			}

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			Expect(res.StatusCode).To(Equal(200))
			Expect(parse.Message).To(Equal("Request From Get Method"))
		})

		It("Should be post request success - first", func() {
			rr, req, err := utils.HttpRequest(http.MethodPost, "/first", nil)
			gin.SetMode(gin.TestMode)
			router.ServeHTTP(rr, req)

			if err != nil {
				Expect(err).To(MatchError(err))
			}

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			Expect(res.StatusCode).To(Equal(200))
			Expect(parse.Message).To(Equal("Request From Post Method"))
		})

		It("Should be get request success - second", func() {
			rr, req, err := utils.HttpRequest(http.MethodGet, "/second", nil)
			router.ServeHTTP(rr, req)

			if err != nil {
				Expect(err).To(MatchError(err))
			}

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			Expect(res.StatusCode).To(Equal(200))
			Expect(parse.Message).To(Equal("Request From Get Method"))
		})

		It("Should be post request success - second", func() {
			rr, req, err := utils.HttpRequest(http.MethodPost, "/second", nil)
			gin.SetMode(gin.TestMode)
			router.ServeHTTP(rr, req)

			if err != nil {
				Expect(err).To(MatchError(err))
			}

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			Expect(res.StatusCode).To(Equal(200))
			Expect(parse.Message).To(Equal("Request From Post Method"))
		})

		It("Should be get request success - third", func() {
			rr, req, err := utils.HttpRequest(http.MethodGet, "/third", nil)
			gin.SetMode(gin.TestMode)
			router.ServeHTTP(rr, req)

			if err != nil {
				Expect(err).To(MatchError(err))
			}

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			Expect(res.StatusCode).To(Equal(200))
			Expect(parse.Message).To(Equal("Request From Get Method"))
		})

		It("Should be post request success - third", func() {
			rr, req, err := utils.HttpRequest(http.MethodPost, "/third", nil)
			gin.SetMode(gin.TestMode)
			router.ServeHTTP(rr, req)

			if err != nil {
				Expect(err).To(MatchError(err))
			}

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			Expect(res.StatusCode).To(Equal(200))
			Expect(parse.Message).To(Equal("Request From Post Method"))
		})
	})
})
