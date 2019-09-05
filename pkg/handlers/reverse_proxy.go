package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/mjosc/rp/pkg/utils"

	"github.com/mjosc/rp/pkg/shared"
)

func NewReverseProxy(dst string) (shared.ReverseProxy, error) {
	target, err := url.ParseRequestURI(dst)
	if err != nil {
		return nil, err
	}
	inner := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.URL.Path = utils.RemoveDirFromPath(r.URL.Path, 0)

			fmt.Println(r.URL)
		},
	}
	return &reverseProxy{
		Inner: inner,
	}, nil
}

type reverseProxy struct {
	Inner *httputil.ReverseProxy
}

func (p *reverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.Inner.ServeHTTP(w, r)
}
