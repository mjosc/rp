package shared

import "net/http"

type ReverseProxy interface {
	http.Handler
}

type HelloProxy interface {
	ReverseProxy
}

type GoodbyeProxy interface {
	ReverseProxy
}
