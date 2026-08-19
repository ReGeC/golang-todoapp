package core_http_middleware

import (
	"net/http"
	"slices"
)


type Middleware func(http.Handler) http.Handler

func ChainMiddleware(
	h http.Handler,
	middlewares ...Middleware,
) http.Handler {
	if len(middlewares) == 0 {
		return h
	}

	for _, m := range slices.Backward(middlewares) {
		h = m(h)
	}

	return h
}