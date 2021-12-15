package yggdrasil

import (
	"fmt"
	"net/http"
	"strings"

	pb "github.com/redhatinsights/yggdrasil/protocol"
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

// An APIresponse represents an unexpected response from an HTTP method call.
type APIresponse struct {
	Code   int
	Body   []byte
	URL    string
	Method string
}

func (res *APIresponse) GetBody() string {
	return string(res.Body)
}

func (res *APIresponse) Export(directite string) pb.APIResponse {
	return pb.APIResponse{
		StatusCode: int64(res.Code),
		Body:       res.GetBody(),
		Directive:  directite,
		Metadata: map[string]string{
			"URL":    res.URL,
			"Method": res.Method,
		},
	}
}

func (res APIresponse) Error() string {
	v := fmt.Sprintf("unexpected response: %v - %v", res.Code, http.StatusText(res.Code))
	if res.Body != nil && len(res.Body) > 0 {
		v += fmt.Sprintf(" (%v)", strings.TrimSpace(string(res.Body)))
	}
	return v
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
