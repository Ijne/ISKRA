package response

const (
	StatusOK    = "ok"
	StatusError = "error"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func Ok() Response {
	return Response{Status: StatusOK}
}
