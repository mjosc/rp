package shared

import "net/http"

type InternalServerError interface {
	http.Handler
}

type BadRequestError interface {
	http.Handler
}
