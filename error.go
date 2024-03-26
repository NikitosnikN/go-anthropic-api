package go_anthropic_api

import "fmt"

type APIErrorDetails struct {
	Type    string
	Message string
}

type APIError struct {
	Type         string          `json:"type"`
	ErrorDetails APIErrorDetails `json:"error"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.ErrorDetails.Message)
}
