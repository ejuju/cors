This package does not enforce the policy set by the package user on incoming requests. It simply ensures browsers can make cross-origin requests to the server without being stopped by CORS errors.
I've seen other implementations of the Go CORS package that will actually for example, prevent requests from unallowed origins from being handled. But I chose not to go for this because of several assumptions detailed below:

Assumptions:

- CORS is meant to protect browser users not the server, proper mechanisms should be implemented server-side to prevent unallowed requests
- Origin: Checking the origin is not a reliable mechanism (it can be chanegd by the client)
- Methods: Checking the method can be (and is usually) handled by the router
- Allow-Headers: Allowing only certain headers will not prevent them from being received (but maybe they could be stripped off requests headers when not allowed?)
- Expose-Headers: Allowing only certain header to be exposed in the CORS policy will not prevent you from sending them by mistake
