package shared

import "net/http"

type Goodbye interface {
	http.Handler
}
