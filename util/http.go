package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/dorman99/point-of-sales/common"
)

func ParseRequestBody(r *http.Request, response interface{}) interface{} {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln("error read body", r.URL)
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Fatalln("error", r.URL, err)
	}

	return response
}

func ParseResponseBody(data interface{}) []byte {
	responsePayload := common.ResponseData{
		Success: "success!",
		Data:    data,
	}
	response, _ := json.Marshal(responsePayload)
	return response
}
