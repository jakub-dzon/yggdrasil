package yggdrasil

import (
	"fmt"
	"net/http"
)

// ErrInvalidContentType indicates an unsupported "collector" value was given
// in the upload request.
var ErrInvalidContentType = &APIresponse{
	Code: http.StatusUnsupportedMediaType,
	Body: []byte("Content type of payload is unsupported"),
}

// ErrPayloadTooLarge indicates an upload request body exceeded the size limit.
var ErrPayloadTooLarge = &APIresponse{
	Code: http.StatusRequestEntityTooLarge,
	Body: []byte("Payload too large"),
}

// ErrUnauthorized indicates an upload request without an Authentication header.
var ErrUnauthorized = &APIresponse{
	Code: http.StatusUnauthorized,
	Body: []byte("Authentication missing from request"),
}

// An InvalidValueTypeError represents an error when serializing data into an
// unsupported destination.
type InvalidValueTypeError struct {
	key string
	val interface{}
}

func (e InvalidValueTypeError) Error() string {
	return fmt.Sprintf("invalid type '%T' for key '%s'", e.val, e.key)
}

// An InvalidArgumentError represents an invalid value passed to a command line
// argument.
type InvalidArgumentError struct {
	flag, value string
}

func (e InvalidArgumentError) Error() string {
	if e.value == "" {
		return "missing value for argument '--" + e.flag + "'"
	}
	return "invalid value '" + e.value + "' for argument '" + e.flag + "'"
}
