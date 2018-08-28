package main

import (
	"encoding/json"
)

func bodyParse(request Request) (*PalitraParams, error) {
	palitraParams := &PalitraParams{}

	if err := json.Unmarshal([]byte(request.Body), palitraParams); err != nil {
		return nil, err
	}

	return palitraParams, nil
}
