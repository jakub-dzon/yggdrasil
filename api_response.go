package yggdrasil

import (
	"fmt"
	"net/http"
	"strings"

	pb "github.com/redhatinsights/yggdrasil/protocol"
)

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

func (res *APIresponse) Export(directite string) *pb.APIResponse {
	return &pb.APIResponse{
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
