package utils

import (
	"encoding/json"
	"log"
)

type ResponseMeta struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Parse(data []byte) ResponseMeta {
	var res ResponseMeta

	err := json.Unmarshal(data, &res)

	if err != nil {
		log.Fatalf("Parse json data error %s", err)
	}

	return res
}
