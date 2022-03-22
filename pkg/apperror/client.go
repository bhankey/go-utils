package apperror

type Code int

const (
	InvalidAuthorization Code = iota + 1
	InvalidParams
	PermissionDenied
	NotFound
	AlreadyExist
	Canceled
	Timeout
	Unavailable
	Internal
)

type ClientError struct {
	Code       Code
	Message    string
	ErrorToLog error
}

func NewClientError(errorType ClientErrorType, err error) ClientError {
	clientError, ok := errorsMap[errorType]
	if !ok {
		clientError = errorsMap[Common]
	}

	clientError.ErrorToLog = err

	return clientError
}

func (err ClientError) Error() string {
	return err.Message
}

func (err ClientError) Unwrap() error {
	return err.ErrorToLog
}
