package customErrors

type ErrorCodeEnum string

const (
	NAME_IS_MANDATORY  ErrorCodeEnum = "NAME_IS_MANDATORY"
	PRICE_IS_MANDATORY ErrorCodeEnum = "PRICE_IS_MANDATORY"
	LOGIN_ERROR        ErrorCodeEnum = "LOGIN_ERROR"
)
