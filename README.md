# CORS utilities for Golang

---

This CORS implementation simply responds to the browser with the right headers, but it does not prevent unallowed requests from being handled.

---

## Features

- CORS middleware for Gorilla Mux
- Easy configuration (using Policy struct)

---

## Roadmap

- Allow dynamic policy `func(r *http.Request) ...` for policy fields
- Access-Control-Request-Private-Network

## Notes

Comments, tips & tricks, etc.

### CORS-safelisted response header

A CORS-safelisted response header is an HTTP header in a CORS response that it is considered safe to expose to client scripts. Only safelisted response headers are made available to web pages.

By default, the safelist includes the following response headers (so they don't need to be included in the `Access-Control-Allow-Headers` field):

- Cache-Control
- Content-Language
- Content-Length
- Content-Type
- Expires
- Last-Modified
- Pragma

But wait... there's additional restrictions... (that only apply when the corresponding above field is not included in the `Access-Control-Allow-Headers` field)

CORS-safelisted headers must also fulfill the following requirements in order to be a CORS-safelisted request header:

- For Accept-Language and Content-Language: can only have values consisting of 0-9, A-Z, a-z, space or \*,-.;=.
- For Accept and Content-Type: can't contain a CORS-unsafe request header byte: 0x00-0x1F (except for 0x09 (HT), which is allowed), "():<>?@[\\]{}, and 0x7F (DEL).
- For Content-Type: needs to have a MIME type of its parsed value (ignoring parameters) of either application x-www-form-urlencoded, multipart/form-data, or text/plain.
- For any header: the value's length can't be greater than 128.

If you need to bypass these restrictions, then you should include the header field name (for ex: ``) in the `AllowHeaders` field on the Policy struct

---

## References

- Great explanation of CORS: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
