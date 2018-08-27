package main

import (
	"encoding/json"
	_ "log"
	_ "net/http"
	_ "net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	basePostRequest = Request{
		HTTPMethod: "POST",
		Path:       "/",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	testImage = "https://images.unsplash.com/photo-1527259216948-b0c66d6fc31f?ixlib=rb-0.3.5&ixid=eyJhcHBfaWQiOjEyMDd9&s=0634d6a5dc2b01309181c1002c5d4e51&auto=format&fit=crop&w=750&q=80"
)

func TestHandler(t *testing.T) {
	t.Run("Bad requests", func(t *testing.T) {
		t.Run("should return an error", func(t *testing.T) {
			_, err := Handler(Request{})
			assert.True(t, err != nil)
		})

		t.Run("should return 400 as HTTP status code", func(t *testing.T) {
			resp, _ := Handler(Request{})
			assert.Equal(t, resp.StatusCode, 400)
		})
	})

	t.Run("Error parse handling", func(t *testing.T) {
		t.Run("parse error message should be in structured properly", func(t *testing.T) {
			response, _ := Handler(Request{})
			handlerResponse := &HandlerResponse{}
			err := json.Unmarshal([]byte(response.Body), handlerResponse)
			assert.Nil(t, err)
			assert.Equal(t, handlerResponse.Errors[fieldMessage], errParseReqParams)
			assert.Equal(t, `application/json`, response.Headers["Content-Type"])
		})
	})

	t.Run("Validating params", func(t *testing.T) {
		t.Run("It should validate missing params", func(t *testing.T) {
			request := basePostRequest
			request.Body = `{ "width": 5, "limit": 10 }`
			response, err := Handler(request)
			assert.Equal(t, err.Error(), errMultipleErrors)

			handlerResponse := &HandlerResponse{}
			err = json.Unmarshal([]byte(response.Body), handlerResponse)
			assert.Nil(t, err)
			assert.Nil(t, handlerResponse.Data)
			assert.Equal(t, len(handlerResponse.Errors), 2)
			assert.Equal(t, handlerResponse.Errors[fieldLimit], errValidateLimit)
			assert.Equal(t, handlerResponse.Errors[fieldURL], errValidateURL)

			t.Log("response", response)
		})
	})
}
