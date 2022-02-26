package cors

// Policy defines the various header fields that can be configured for Cross-Origin Resource Sharing
//
// AllowOrigin ("Access-Control-Allow-Origin") specifies either a single origin (which tells browsers that this origin can access the resource); or else (for requests without credentials) the "*" wildcard tells browsers to allow any origin to access the resource.
//
// AllowMethods ("Access-Control-Allow-Methods") specifies the method or methods allowed when accessing the resource. This is used in response to a preflight request. The conditions under which a request is preflighted are discussed above.
//
// AllowHeaders ("Access-Control-Allow-Headers") is used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.
//
// AllowCredentials ("Access-Control-Allow-Credentials") indicates whether or not the response to the request can be exposed when the credentials flag is true. When used as part of a response to a preflight request, this indicates whether or not the actual request can be made using credentials. Note that simple GET requests are not preflighted,
//
// CacheMaxAge ("Access-Control-Max-Age") indicates for how many seconds the results of a preflight request can be cached in the browser.
//
// ExposeHeaders ("Access-Control-Expose-Headers") adds the specified headers to the allowlist that JavaScript (such as getResponseHeader()) in browsers is allowed to access.
//
type Policy struct {
	// AllowOrigin      string   // can be either an origin (with the scheme, for ex: "https://example.com") or an asterisk ("*")
	AllowOrigin      func(string) string // takes the request origin and returns the either "*" or the sender's origin if they should be allowed to access the request
	AllowMethods     []string            // for ex: "GET", "POST", "DELETE", "PUT" (allowed by default: "OPTIONS")
	AllowHeaders     []string            // for ex: "Origin, X-Custom-Header, Authorization" (allowed by default: "Cache-Control", "Content-Language", "Content-Length", "Content-Type", "Expires", "Last-Modified", "Pragma")
	ExposeHeaders    []string            // for ex: "Content-Encoding, Authorization"
	Vary             []string            // for ex: "Content-Encoding, Authorization"
	AllowCredentials bool                // can be "true" or "false" (bool for convenience)
	CacheMaxAge      int                 // in seconds (put 5 seconds if you don't know how much you should put, should not exceed 10 minutes (= 600 seconds))
}

// Keys for the header fields
const (
	KeyAllowOrigin      = "Access-Control-Allow-Origin"
	KeyAllowMethods     = "Access-Control-Allow-Methods"
	KeyAllowHeaders     = "Access-Control-Allow-Headers"
	KeyAllowCredentials = "Access-Control-Allow-Credentials"
	KeyExposeHeaders    = "Access-Control-Expose-Headers"
	KeyRequestMethods   = "Access-Control-Request-Methods"
	KeyRequestHeaders   = "Access-Control-Request-Headers"
	KeyMaxAge           = "Access-Control-Max-Age"
	KeyOrigin           = "Origin"
	KeyVary             = "Vary"
)
