package apperror

type ClientErrorType int

const (
	Common ClientErrorType = iota
	WrongAuthorization
	WrongRequest
	WrongAuthToken
	NoClient
	WrongOneTimeCode
)

// thread safe, cus nobody writes.
var errorsMap = map[ClientErrorType]ClientError{
	Common:             errSomethingWentWrong,
	WrongRequest:       errWrongRequest,
	WrongAuthorization: errWrongAuthorization,
	WrongAuthToken:     errWrongAuthToken,
	NoClient:           errNoClient,
	WrongOneTimeCode:   errWrongOneTimeCode,
}

var errWrongAuthorization = ClientError{
	Code:    InvalidAuthorization,
	Message: "wrong password or email",
}

var errSomethingWentWrong = ClientError{
	Code:    Internal,
	Message: "something went wrong",
}

var errWrongRequest = ClientError{
	Code:    InvalidParams,
	Message: "wrong request",
}

var errWrongAuthToken = ClientError{
	Code:    InvalidAuthorization,
	Message: "wrong auth token",
}

var errNoClient = ClientError{
	Code:    InvalidParams,
	Message: "client doesn't exist",
}

var errWrongOneTimeCode = ClientError{
	Code:    InvalidParams,
	Message: "wrong one-time code",
}
