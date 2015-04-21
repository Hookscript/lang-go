# Go Reference

See also Hookscript's [general documentation](http://docs.hookscript.com/)

## HTTP Request

An incoming HTTP request is represented by a [*http.Request](http://golang.org/pkg/net/http/#Request) value.
If your script needs access to the request, include a `*http.Request` parameter on your `Hook` function. For
example,

```go
func Hook(r *http.Request) { ... }
```

Some fields of the `Request` deserve clarification:

  * `RemoteAddr` is not populated.  Instead use the `X-Forwarded-For` header. See #1
  * `TLS` is not populated, even if the original request arrived via an encrypted connection

## HTTP Response

A script doesn't need an explicit response variable.  Anything sent to `os.Stdout` becomes the
HTTP response body.  However, if you want to modify HTTP response headers, you may include
a [*http.Response](http://golang.org/pkg/net/http/#Response) parameter on your `Hook` function.  For
example,

```go
func Hook(res *http.Response) { ... }
```
