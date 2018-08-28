package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/jkomyno/palitra"
)

// PalitraParams is the type of the request body
type PalitraParams struct {
	// URL is the url of image to inspect
	URL string `json:"url"`

	// Limit is the amount of the nearest CSS3 colors to detect
	Limit int `json:"limit"`

	// Width is the width of the resized image
	Width uint `json:"width"`
}

// PalitraErrors is a dictionary of validation/parsing errors displayed in the response
type PalitraErrors map[string]string

// HandlerResponseData is the palette list with percentages returned by palitra on success
type HandlerResponseData []palitra.ColorPercentageT

// HandlerResponse is the type of the handler response
type HandlerResponse struct {
	Errors PalitraErrors       `json:"errors,omitempty"`
	Data   HandlerResponseData `json:"data,omitempty"`
}

// Response is the response to be returned by API Gateway for the request
type Response events.APIGatewayProxyResponse

// Request contains data coming from the API Gateway proxy
type Request events.APIGatewayProxyRequest
