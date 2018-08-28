package main

import (
	"bytes"
	"encoding/json"
)

var defaultHeaders = map[string]string{
	"Content-Type": "application/json",
}

func errorResponse(errs PalitraErrors, statusCode int) Response {
	var body string
	response := &HandlerResponse{
		Errors: errs,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		body = errParseError
	} else {
		var buf bytes.Buffer
		json.HTMLEscape(&buf, jsonResponse)
		body = buf.String()
	}

	return Response{
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Body:            body,
		Headers:         defaultHeaders,
	}
}

func successResponse(response *HandlerResponse) Response {
	var body string

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		body = errParseData
	} else {
		var buf bytes.Buffer
		json.HTMLEscape(&buf, jsonResponse)
		body = buf.String()
	}

	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            body,
		Headers:         defaultHeaders,
	}
}
