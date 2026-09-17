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
curl http://localhost:8080/health
curl 'http://localhost:8080/greet?name=Tuneesh'
```

Expected responses are `{"status":"ok"}` and `{"message":"hello, Tuneesh"}`.
