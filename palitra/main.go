package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jkomyno/palitra"
)

const (
	errParseError     = "Can't parse errors"
	errParseData      = "Can't parse data"
	errParseReqParams = "Can't parse request params"
	errParseImageURL  = "Can't parse image from URL"

	errMultipleErrors = "Multiple errors"

	fieldMessage = "message"
)

func handleSingleError(errorMessage string, statusCode int) Response {
	errs := make(PalitraErrors)
	errs[fieldMessage] = errorMessage

	return errorResponse(errs, statusCode)
}

// Handler is the lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	palitraParams, err := bodyParse(request)
	if err != nil {
		return handleSingleError(errParseReqParams, 400), err
	}

	errs, hasErrors := palitraParams.validate()
	if hasErrors {
		return errorResponse(errs, 400), errors.New(errMultipleErrors)
	}

	img, err := getImageFromURL(palitraParams.URL)
	if err != nil {
		return handleSingleError(errParseImageURL, 500), err
	}

	palette := palitra.GetPalette(img, palitraParams.Limit, palitraParams.Width)

	response := &HandlerResponse{
		Data: palette,
	}

	return successResponse(response), nil
}

func main() {
	lambda.Start(Handler)
}
