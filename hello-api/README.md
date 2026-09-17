# Hello API

A tiny HTTP API using only Go's standard library. It is intentionally separate
from the cache experiment at the repository root, so its workflow is easy to
read without cache-policy noise.

Run it locally:

```bash
go run ./cmd/hello-api
```

Then, in another terminal:

```bash
curl http://localhost:8080/
curl http://localhost:8080/health
curl 'http://localhost:8080/greet?name=Tuneesh'
curl 'http://localhost:8080/greet?name=Tuneesh&shout=true'
```

Expected responses are `{"service":"hello-api"}`, `{"status":"ok"}`, and
`{"message":"hello, Tuneesh"}`. Add `shout=true` to receive
`{"message":"HELLO, TUNEESH"}`.

## CI artifacts

The GitHub Actions build job creates two downloadable executables:

- `hello-api-linux-amd64` for Linux x64 servers.
- `hello-api-macos-arm64` for an Apple-Silicon Mac.

They are separate because a compiled Go binary targets one operating system
and CPU architecture at a time.
