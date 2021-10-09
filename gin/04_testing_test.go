package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/restuwahyu13/go-supertest/supertest"
	"github.com/stretchr/testify/assert"

	"github.com/restuwahyu13/go-testing-fundamental/gin/utils"
)

// initialize router
var app = Router()

// ======================
// first testing teritory
// ======================

func TestGetRequestFirst(tt *testing.T) {

	tt.Run("Shoud be get data from GET success", func(t *testing.T) {
		assert := assert.New(t)

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(rr)
		GetRequestFirst(c)

		res := rr.Result()
		parse := utils.Parse(rr.Body.Bytes())

		assert.Equal(res.StatusCode, 200)
		assert.Equal(parse.Code, 200)
		assert.Equal(parse.Message, "Request From Get Method")
	})
}

func TestPostRequestFirst(tt *testing.T) {

	tt.Run("Shoud be get data from POST success ", func(t *testing.T) {
		assert := assert.New(t)

		rr := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(rr)
		PostRequestFirst(c)

		res := rr.Result()
		parse := utils.Parse(rr.Body.Bytes())

		assert.Equal(res.StatusCode, 200)
		assert.Equal(parse.Code, 200)
		assert.Equal(parse.Message, "Request From Post Method")
	})
}

// ======================
// second testing teritory
// ======================

func TestGetRequestSecond(tt *testing.T) {

	tt.Run("Shoud be second get data from GET success not custom httptest", func(t *testing.T) {
		assert := assert.New(t)

		req, _ := http.NewRequest(http.MethodGet, "/second", nil)
		rr := httptest.NewRecorder()

		app.ServeHTTP(rr, req)

		res := rr.Result()
		parse := utils.Parse(rr.Body.Bytes())

		assert.Equal(res.StatusCode, 200)
		assert.Equal(parse.Code, 200)
		assert.Equal(parse.Message, "Request From Get Method")
	})

	tt.Run("Shoud be second get data from GET success with custom httptest", func(t *testing.T) {
		assert := assert.New(t)

		rr, req, _ := utils.HttpRequest(http.MethodGet, "/second", nil)

		app.ServeHTTP(rr, req)

		res := rr.Result()
		parse := utils.Parse(rr.Body.Bytes())

		assert.Equal(res.StatusCode, 200)
		assert.Equal(parse.Code, 200)
		assert.Equal(parse.Message, "Request From Get Method")
	})
}

func TestPostRequestSecond(tt *testing.T) {

	tt.Run("Shoud be second get data from POST success not custom httptest", func(t *testing.T) {
		assert := assert.New(t)

		req, _ := http.NewRequest(http.MethodPost, "/second", nil)
		rr := httptest.NewRecorder()

		app.ServeHTTP(rr, req)

		res := rr.Result()
		parse := utils.Parse(rr.Body.Bytes())

		assert.Equal(res.StatusCode, 200)
		assert.Equal(parse.Code, 200)
		assert.Equal(parse.Message, "Request From Post Method")
	})

	tt.Run("Shoud be second get data from POST success with custom httptest", func(t *testing.T) {
		assert := assert.New(t)

		rr, req, _ := utils.HttpRequest(http.MethodPost, "/second", nil)

		app.ServeHTTP(rr, req)

		res := rr.Result()
		parse := utils.Parse(rr.Body.Bytes())

		assert.Equal(res.StatusCode, 200)
		assert.Equal(parse.Code, 200)
		assert.Equal(parse.Message, "Request From Post Method")
	})
}

// ======================
// third testing teritory
// ======================

func TestGetRequestThird(tt *testing.T) {

	tt.Run("Shoud be second get data from GET success using supertest", func(t *testing.T) {

		assert := assert.New(t)
		testing := supertest.NewSuperTest(app, t)

		testing.Get("/third")
		testing.Send(nil)
		testing.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			assert.Equal(res.StatusCode, 200)
			assert.Equal(parse.Code, 200)
			assert.Equal(parse.Message, "Request From Get Method")
		})
	})
}

func TestPostRequestThird(tt *testing.T) {

	tt.Run("Shoud be second get data from POST success using supertest", func(t *testing.T) {

		assert := assert.New(t)
		testing := supertest.NewSuperTest(app, t)
		payload := gin.H{}

		testing.Post("/third")
		testing.Send(payload)
		testing.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			assert.Equal(res.StatusCode, 200)
			assert.Equal(parse.Code, 200)
			assert.Equal(parse.Message, "Request From Post Method")
		})
	})

	tt.Run("Shoud be second get data from POST success using supertest with payload", func(t *testing.T) {

		assert := assert.New(t)
		testing := supertest.NewSuperTest(app, t)
		payload := gin.H{"message": "Request From Testing"}

		testing.Post("/third")
		testing.Send(payload)
		testing.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

			res := rr.Result()
			parse := utils.Parse(rr.Body.Bytes())

			assert.Equal(res.StatusCode, 200)
			assert.Equal(parse.Code, 200)
			assert.Equal(parse.Message, "Request From Testing")
		})
	})
}
