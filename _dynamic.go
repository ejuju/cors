package cors

import "net/http"

// DynamicPolicy allows users of this package to dynamically set each header field based on the current request
type DynamicPolicy struct {
	AllowOriginFunc          func(r *http.Request) string //
	AllowMethodsFunc         func(r *http.Request) string //
	AllowHeadersFunc         func(r *http.Request) string //
	AllowCredentialsFunc     func(r *http.Request) string //
	PreflightCacheMaxAgeFunc func(r *http.Request) string //
	ExposeHeadersFunc        func(r *http.Request) string //
}
