package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {

	var response Response
	assert := assert.New(t)

	rq := httptest.NewRequest(http.MethodGet, "/get", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetRequest)

	handler.ServeHTTP(rr, rq)

	err := json.Unmarshal([]byte(rr.Body.String()), &response)

	if err != nil {
		t.Error("Parse JSON Data Error")
	}

	if response.Message == "Bad Request" {
		t.Error("Request Failed")
	}

	r := rr.Result()

	assert.Equal(r.StatusCode, 200)
	assert.Equal(r.Header.Get("content-type"), "application/json")
	assert.Equal(response.Message, "Hello World From - GET")
}

func TestPostRequest(t *testing.T) {

	var response Response
	assert := assert.New(t)

	body := "Data From TestRequest"

	rq := httptest.NewRequest(http.MethodPost, "/post", bytes.NewBufferString(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostRequest)

	handler.ServeHTTP(rr, rq)

	err := json.Unmarshal([]byte(rr.Body.String()), &response)

	if err != nil {
		t.Error("Parse JSON Data Error")
	}

	if response.Message == "Bad Request" {
		t.Error("Request Failed")
	}

	r := rr.Result()

	assert.Equal(r.StatusCode, 200)
	assert.Equal(r.Header.Get("content-type"), "application/json")
	assert.Equal(response.Message, "Data From TestRequest")
}