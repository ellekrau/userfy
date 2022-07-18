package authmiddleware

type Response struct {
	Message string `json:"message"`
}

func newResponse(message string) Response {
	return Response{
		Message: message,
	}
}
