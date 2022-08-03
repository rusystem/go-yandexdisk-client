package responses

import (
	"fmt"
)

type ErrorResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
	Error       string `json:"error"`
}

type SuccessResponse struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}

func (e *ErrorResponse) Info() string {
	return fmt.Sprintf("message: %s, description: %s, error: %s\n", e.Message, e.Description, e.Error)
}
