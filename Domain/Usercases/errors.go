package Usercases

type ApiErr struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
	Error      string `json:"error"`
}

func NewBadRequest(message string) *ApiErr {

	return &ApiErr{
		Message:    message,
		StatusCode: 400,
		Error:      "bad_request",
	}
}

func NewInternalError(message string) *ApiErr {
	return &ApiErr{
		Message:    message,
		StatusCode: 500,
		Error:      "bad_request",
	}
}
