package utils

type HttpResponse struct {
	Code    int
	Message string
}

func ResponseBuilder(code int, message string) HttpResponse {
	return HttpResponse{
		code,
		message,
	}
}
