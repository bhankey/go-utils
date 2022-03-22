package apperror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (err ClientError) GetGRPCError() error {
	grpcErr := status.New(err.GetGRPCode(), err.Message)
	// grpcErr.WithDetails() TODO think about details https://cloud.google.com/apis/design/errors#error_model

	return grpcErr.Err() // nolint: wrapcheck, nolintlint
}

func (err ClientError) GetGRPCode() codes.Code {
	return map[Code]codes.Code{
		InvalidAuthorization: codes.Unauthenticated,
		InvalidParams:        codes.InvalidArgument,
		PermissionDenied:     codes.PermissionDenied,
		NotFound:             codes.NotFound,
		AlreadyExist:         codes.AlreadyExists,
		Canceled:             codes.Canceled,
		Timeout:              codes.DeadlineExceeded,
		Unavailable:          codes.Unavailable,
		Internal:             codes.Internal,
	}[err.Code]
}

func NewClientErrorFromGRPC(err *status.Status) ClientError {
	return ClientError{
		Code:    getCodeFromGRPC(err.Code()),
		Message: err.Message(),
	}
}

func getCodeFromGRPC(grpcCode codes.Code) Code {
	return map[codes.Code]Code{
		codes.Unauthenticated:  InvalidAuthorization,
		codes.InvalidArgument:  InvalidParams,
		codes.PermissionDenied: PermissionDenied,
		codes.NotFound:         NotFound,
		codes.AlreadyExists:    AlreadyExist,
		codes.Canceled:         Canceled,
		codes.DeadlineExceeded: Timeout,
		codes.Unavailable:      Unavailable,
		codes.Internal:         Internal,
	}[grpcCode]
}
