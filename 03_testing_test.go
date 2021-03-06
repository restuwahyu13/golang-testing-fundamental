package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(tt *testing.T) {

	tt.Run("Should be get request data success", func(t *testing.T) {

		var response Response
		assert := assert.New(t)

		rq := httptest.NewRequest(http.MethodGet, "/get", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HttpGetRequest)

		handler.ServeHTTP(rr, rq)

		err := json.Unmarshal([]byte(rr.Body.String()), &response)

		if err != nil {
			t.FailNow()
		}

		if response.Message == "Bad Request" {
			t.FailNow()
		}

		r := rr.Result()

		assert.Equal(r.StatusCode, 200)
		assert.Equal(r.Header.Get("content-type"), "application/json")
		assert.Equal(response.Message, "Hello World From - GET")
	})

	tt.Run("Should be get request data error", func(t *testing.T) {

		var response Response
		assert := assert.New(t)

		rq := httptest.NewRequest(http.MethodGet, "/getx", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HttpGetRequest)

		handler.ServeHTTP(rr, rq)

		err := json.Unmarshal([]byte(rr.Body.String()), &response)

		if err != nil {
			t.FailNow()
		}

		r := rr.Result()

		assert.Equal(r.StatusCode, 400)
		assert.Equal(r.Header.Get("content-type"), "application/json")
		assert.Equal(response.Message, "Bad Request")
	})
}

func TestPostRequest(tt *testing.T) {

	tt.Run("Should be not parsing data from body", func(t *testing.T) {

		var response Response
		assert := assert.New(t)

		rq := httptest.NewRequest(http.MethodPost, "/post", nil)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HttpPostRequest)

		handler.ServeHTTP(rr, rq)

		err := json.Unmarshal([]byte(rr.Body.String()), &response)

		if err != nil {
			t.FailNow()
		}

		if response.Message == "Bad Request" {
			t.FailNow()
		}

		r := rr.Result()

		assert.Equal(r.StatusCode, 200)
		assert.Equal(r.Header.Get("content-type"), "application/json")
		assert.Equal(response.Message, "Hello World From - POST")
	})

	tt.Run("Should be parsing data from body success", func(t *testing.T) {

		var response Response
		assert := assert.New(t)

		body := "Data From TestRequest"

		rq := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HttpPostRequest)

		handler.ServeHTTP(rr, rq)

		err := json.Unmarshal([]byte(rr.Body.String()), &response)

		if err != nil {
			t.FailNow()
		}

		if response.Message == "Bad Request" {
			t.FailNow()
		}

		r := rr.Result()

		assert.Equal(r.StatusCode, 200)
		assert.Equal(r.Header.Get("content-type"), "application/json")
		assert.Equal(response.Message, "Data From TestRequest")
	})

	tt.Run("Should be parsing data from body error", func(t *testing.T) {

		var response Response
		assert := assert.New(t)

		body := "Data From TestRequest"

		rq := httptest.NewRequest(http.MethodPost, "/postx", bytes.NewBufferString(body))
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HttpPostRequest)

		handler.ServeHTTP(rr, rq)

		err := json.Unmarshal([]byte(rr.Body.String()), &response)

		if err != nil {
			t.FailNow()
		}

		r := rr.Result()

		assert.Equal(r.StatusCode, 400)
		assert.Equal(r.Header.Get("content-type"), "application/json")
		assert.Equal(response.Message, "Bad Request")
	})
}
