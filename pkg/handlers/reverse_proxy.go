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
			// We've mapped the first directory of each incoming path to a third-party service. This directory
			// must be removed before proxying the request to each service.
			path := utils.RemoveDirFromPath(r.URL.Path, 0)

			r.URL.Scheme = target.Scheme
			r.URL.Host = target.Host
			r.URL.Path = path
		},
		ModifyResponse: func(res *http.Response) error {
			header := res.Header.Get(shared.HeaderModifyResponseError)
			if len(header) > 0 {
				return fmt.Errorf("ReverseProxy.ModifyResponse received %v header: %v", shared.HeaderModifyResponseError, header)
			}
			return nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			w.WriteHeader(500)

			body := fmt.Sprintf("Internal Server Error: [ source: ReverseProxy.ErrorHandler, error: %v ]", err)
			w.Write([]byte(body))
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
