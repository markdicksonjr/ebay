package ebay

import "net/http"

type Response struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
}
