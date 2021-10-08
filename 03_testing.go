package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func GetRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method == "GET" && req.URL.Path == "/get" {
		data := Response{Message: "Hello World From - GET"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	} else {
		data := Response{Message: "Bad Request"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	}
}

func PostRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method == "POST" && req.URL.Path == "/post" {

		data, err := io.ReadAll(req.Body)
		req.Body.Close()

		if err != nil {
			data := Response{Message: "Hello World From - POST"}
			json, _ := json.Marshal(data)
			fmt.Fprint(rw, string(json))
		} else {
			data := Response{Message: string(data)}
			json, _ := json.Marshal(data)
			fmt.Fprint(rw, string(json))
		}

	} else {
		data := Response{Message: "Bad Request"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	}
}

func main() {
	http.HandleFunc("/get", GetRequest)
	http.HandleFunc("/post", PostRequest)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
