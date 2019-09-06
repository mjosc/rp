# rp

A reverse proxy proof of concept with an emphasis on the [ModifyResponse](https://golang.org/pkg/net/http/httputil/#ReverseProxy) and [ErrorHandler](https://golang.org/pkg/net/http/httputil/#ReverseProxy) functions.

## Getting Started

Run `make rp` to run the server. This will startup the proxy server as well as the mock servers representing various external services. You can configure the ports for each server in `cmd/main.go`. All servers will run on localhost.