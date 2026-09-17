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
curl 'http://localhost:8080/greet?name=Tuneesh&reverse=true'
curl http://localhost:8080/request-id
```

Expected responses are `{"service":"hello-api"}`, `{"status":"ok"}`, and
`{"message":"hello, Tuneesh"}`. Add `shout=true` to receive
`{"message":"HELLO, TUNEESH"}`. `/request-id` returns a fresh UUID, supplied
by the external `github.com/google/uuid` dependency. Add `reverse=true` to
receive `{"message":"hseenuT ,olleh"}`.

## Dependencies and cache keys

`go.mod` declares that this API needs `github.com/google/uuid v1.6.0`.
`go.sum` records the checksum Go expects when it downloads that module. The CI
workflow uses both files as its Go-cache fingerprint, so changing a dependency
causes a new cache entry instead of reusing one built for an older dependency
set.

## Cache lab

`Hello API Cache Lab` is a manually triggered workflow that compares three
options for the same vet, test, and build commands:

- `actions/setup-go`'s built-in, shared Go cache.
- Explicit `actions/cache` entries isolated by job type.
- `cloudx-io/setup-go` entries isolated by job type.

The explicit Actions cache is the apples-to-apples baseline for CloudX: it
caches the same `GOMODCACHE` and `GOCACHE` directories but gives test, vet,
and build their own keys. The lab is deliberately separate from normal CI.
This API is tiny, so use its test-command duration and the `(cached)` marker as
learning evidence, not as a performance claim. In a production workflow, pin
third-party Actions to reviewed commit SHAs instead of a moving major-version
tag.

## CI artifacts

The GitHub Actions build job creates two downloadable executables:

- `hello-api-linux-amd64` for Linux x64 servers.
- `hello-api-macos-arm64` for an Apple-Silicon Mac.

They are separate because a compiled Go binary targets one operating system
and CPU architecture at a time.

## CI smoke test

After the build job uploads the Linux binary, a separate `smoke-test` job
downloads that exact artifact, starts it, and checks `/health`. It does not
check out source code or compile Go: this demonstrates how jobs pass a real
deliverable through an artifact rather than a cache.
