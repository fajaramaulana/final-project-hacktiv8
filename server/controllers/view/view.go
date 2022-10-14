package view

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Payload map[string]string `json:"payload,omitempty"`
}

func ErrorValidation(status int, message string, payload map[string]string) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}
}

func Error(status int, message string) *Response {
	return &Response{
		Status:  status,
		Message: message,
	}
}
