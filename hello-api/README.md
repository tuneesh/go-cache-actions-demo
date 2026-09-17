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
