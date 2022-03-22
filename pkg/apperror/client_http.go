package apperror

import "net/http"

func (err ClientError) GetHTTPCode() int {
	return map[Code]int{
		InvalidAuthorization: http.StatusUnauthorized,
		InvalidParams:        http.StatusBadRequest,
		PermissionDenied:     http.StatusForbidden,
		NotFound:             http.StatusNotFound,
		AlreadyExist:         http.StatusConflict,
		Canceled:             http.StatusRequestTimeout,
		Timeout:              http.StatusRequestTimeout,
		Unavailable:          http.StatusServiceUnavailable,
		Internal:             http.StatusInternalServerError,
	}[err.Code]
}
