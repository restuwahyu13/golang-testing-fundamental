package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func HttpGetRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method == "GET" && req.URL.Path == "/get" {
		data := Response{Message: "Hello World From - GET"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		data := Response{Message: "Bad Request"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	}
}

func HttpPostRequest(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	if req.Method == "POST" && req.URL.Path == "/post" {

		data, _ := io.ReadAll(req.Body)
		req.Body.Close()

		if len(string(data)) < 1 {
			data := Response{Message: "Hello World From - POST"}
			json, _ := json.Marshal(data)
			fmt.Fprint(rw, string(json))
		} else {
			data := Response{Message: string(data)}
			json, _ := json.Marshal(data)
			fmt.Fprint(rw, string(json))
		}

	} else {
		rw.WriteHeader(http.StatusBadRequest)
		data := Response{Message: "Bad Request"}
		json, _ := json.Marshal(data)
		fmt.Fprint(rw, string(json))
	}
}

// func main() {
// 	http.HandleFunc("/get", HttpGetRequest)
// 	http.HandleFunc("/post", HttpPostRequest)

// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }
