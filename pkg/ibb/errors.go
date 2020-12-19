package ibb

import "errors"

const (
	errParsing       = "failed to parse json"
	errReqDataF      = "request data: code = %d, status = %s"
	errCreateRequest = "failed to create request"
	errDoRequest     = "failed to perform request"
)

var (
	// ErrRequestFailed - response status != 200
	ErrRequestFailed = errors.New("failed to perform http request")
)
