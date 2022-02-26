package cors

import (
	"net/http"
	"strings"
)

// Set returns a middleware function that will add the CORS headers to server responses
func Set(p Policy) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(KeyAllowHeaders, KeyOrigin) // always allow origin in request header
			w.Header().Add(KeyVary, KeyOrigin)         // always set origin in Vary header to prevent cache poisoning https://github.com/rs/cors/issues/10

			// check if request is preflight request
			if IsPreflightRequest(r) {
				// set default headers (vary headers should be set for caching, see: https://github.com/rs/cors/issues/10, https://github.com/fastify/fastify-cors/pull/45, https://textslashplain.com/2018/08/02/cors-and-vary/)
				w.Header().Add(KeyVary, KeyRequestMethods)
				w.Header().Add(KeyVary, KeyRequestHeaders)

				// validate origin
				origin := r.Header.Get(KeyOrigin)
				if origin == "" {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("no origin provided in preflight request"))
					return
				}
				allowedOrigin := p.AllowOrigin(origin)
				if allowedOrigin != origin && allowedOrigin != "*" {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("origin is not allowed"))
					return
				}

				// validate request method
				method := r.Header.Get(KeyRequestMethods)
				if !methodIsAllowed(method, p.AllowMethods) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("method is not allowed"))
					return
				}

				// validate request headers
				headers := r.Header.Get(KeyRequestHeaders)
				if !headersAreAllowed(headers, p.AllowHeaders) {
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("headers are not allowed"))
					return
				}

				// set response headers
				w.Header().Set(KeyAllowOrigin, allowedOrigin)

				// todo: https://github.com/rs/cors/blob/master/cors.go (line 317)
				w.WriteHeader(http.StatusOK)
				return
			}

			// not a preflight request
			next.ServeHTTP(w, r)
		})
	}
}

// IsPreflightRequest returns true if the request is a preflight request
func IsPreflightRequest(r *http.Request) bool {
	return r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != ""
}

//
func methodIsAllowed(method string, allowedMethods []string) bool {
	included := false
	for _, m := range allowedMethods {
		if m == method {
			included = true
		}
	}
	return included
}

//
func headersAreAllowed(headers string, allowedHeaders []string) bool {
	fields := strings.Split(headers, ",")
	for _, f := range fields {
		if !headerIsAllowed(f, allowedHeaders) {
			return false
		}
	}
	return true
}

func headerIsAllowed(header string, allowedHeaders []string) bool {
	included := false
	for _, h := range allowedHeaders {
		if header == h {
			included = true
		}
	}
	return included
}
