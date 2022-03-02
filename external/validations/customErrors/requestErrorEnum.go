package customErrors

type RequestErrorEnum string

const (
	BAD_REQUEST    RequestErrorEnum = "BAD_REQUEST"
	INTERNAL_ERROR RequestErrorEnum = "INTERNAL_ERROR"
)
