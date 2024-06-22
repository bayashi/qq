# qq

Quick Questions and answers

<a href="https://github.com/bayashi/qq/actions" title="qq CI"><img src="https://github.com/bayashi/qq/workflows/main/badge.svg" alt="qq CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/qq" title="qq report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/qq" alt="qq report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/qq" title="Go qq package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/qq.svg" alt="Go Reference: qq"></a>

`qq` provides a command `qq` to answer quick questions:

* gRPC Status Code
* HTTP Status Code
* JSON-RPC Status Code
* MIME TYPE (Content-Type)

## Usage

To see all list

```
$ qq grpc
$ qq http
$ qq jsonrpc
$ qq mimetype
```

To see the details for each

```
$ qq grpc 4
4 DeadlineExceeded
The deadline expired before the operation could complete. For operations that change the state of the system, this error may be returned even if the operation has completed successfully.  For example, a successful response from a server could have been delayed long enough for the deadline to expire.
HTTP Mapping: 504 Gateway Timeout

$ qq http 403
403 Forbidden
The 403 (Forbidden) status code indicates that the server understood the request but refuses to fulfill it. A server that wishes to make public why the request has been forbidden can describe that reason in the response content (if any).

$ qq mimetype json
json application/json

$ qq jsonrpc -32700
-32700 Parse error
Invalid JSON was received by the server. An error occurred on the server while parsing the JSON text.
```

To search

```
$ qq grpc auth
16 Unauthenticated

$ qq http 30
300 Multiple Choices
301 Moved Permanently
302 Found
303 See Other
304 Not Modified
305 Use Proxy
306 (Unused)
307 Temporary Redirect
308 Permanent Redirect
```

## Custom Dictionary

You can put your custom dictionary in [XDG Base Directory](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html).

Example: `err.yaml`

```yaml
# int
100:
  subject: "No Impact Error"
  description: "No Impact Error description"
101:
  subject: "Really Fatal Error"
  description: "Really Fatal Error description"
```

The first line of YAML file **MUST** be a comment as `# int` or `# string` depends on your dictionary key type.

```
$ qq err 101
101 Really Fatal Error
Really Fatal Error description
```

## Installation

There are several ways to install `qq`:

### Go manual

If you have golang environment:

```cmd
go install github.com/bayashi/qq@latest
```

### Mac brew

```cmd
brew tap bayashi/tap
brew install bayashi/tap/qq
```

### Binary install

Download binary from here: https://github.com/bayashi/qq/releases


## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi

## See Also

* https://metacpan.org/pod/App::httpstatus
* https://github.com/xaicron/jsonrpcstatus
* https://metacpan.org/pod/App::contenttype
