package model

import "encoding/json"

type ErrorBody struct {
	StatusCode int    `json:"name"`
	Message    string `json:"message"`
}

type ErrorResponse struct {
	Body ErrorBody `json:"error"`
}

func GetErrorResponseFromJson(b []byte) (*ErrorResponse, error) {
	var errorResponse ErrorResponse

	e := json.Unmarshal(b, &errorResponse)

	return &errorResponse, e
}
