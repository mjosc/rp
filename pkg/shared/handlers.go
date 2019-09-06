package shared

import "net/http"

// ReverseProxy is a wrapper around http.Handler to aide FX in tracking depenencies by type.
type ReverseProxy interface {
	http.Handler
}

// HelloProxy is a reverse proxy forwarding requests to HelloService.
type HelloProxy interface {
	ReverseProxy
}

// GoodbyeProxy is a reverse proxy forwarding requests to GoodbyeService.
type GoodbyeProxy interface {
	ReverseProxy
}

// BadProxy is a reverse proxy forwarding requests to BadService which returns various error cases.
type BadProxy interface {
	ReverseProxy
}

// NoConnectionProxy is a reverse proxy forwarding requests to NoConnectionService which is literally
// a service that does not exist and therefore no connection can be made.
type NoConnectionProxy interface {
	ReverseProxy
}
