package customErrors

type RequestErrorEnum string

const (
	BAD_REQUEST RequestErrorEnum = "BAD_REQUEST"
	NOT_FOUND   RequestErrorEnum = "NOT_FOUND"
)
