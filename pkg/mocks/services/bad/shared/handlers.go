package shared

import "net/http"

type InternalServerError interface {
	http.Handler
}

type BadRequestError interface {
	http.Handler
}

// ModifyResponseError adds a header to the response indicating to the proxy service to return an error
// from ReverseProxy.ModifyResponse, thus triggering a call to ReverseProxy.ErrorHandler.
type ModifyResponseError interface {
	http.Handler
}
