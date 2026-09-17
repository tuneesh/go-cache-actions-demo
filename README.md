# GitHub Actions Go-cache comparison

This is a deliberately small, repeatable comparison for Go CI caching. It is
not an application; its packages have two-second tests so the work performed is
visible in a GitHub Actions log.

## What the workflow compares

Every push starts these jobs in parallel:

| Job family | Cache policy | What it answers |
| --- | --- | --- |
| `Default cache` | GitHub's `actions/setup-go` built-in cache, shared by lint/test/build | What happens with the conventional setup? |
| `Isolated cache` | `cloudx-io/setup-go`, one cache family for each job purpose | Does a test-specific cache retain useful Go test results? |
| `Native job-isolated cache` | The same general idea written directly with `actions/cache` | Is CloudX adding a capability, or packaging a cache-key policy? |

The first two families each run lint, test, and build. That matters: GitHub
Actions jobs run on separate fresh machines. They do not share a Jenkins-style
workspace, and a GitHub cache entry is immutable once saved. The default jobs
therefore compete to save one shared cache. The isolated jobs use different key
prefixes such as `test`, `lint`, and `build`.

`go.mod` is explicitly used as the dependency fingerprint. A production project
that has one should normally use `go.sum` (or both module files) instead.

## How to run the experiment

1. Push the green baseline and let its workflow finish. This is the cold run.
2. Change only `api/api.go` and adjust `api/api_test.go` to match.
3. Push again and open the three `… test` job logs.
4. Compare the `time go test` output and the package suffixes:
   - `2.00s` means that package's test actually ran.
   - `(cached)` means Go reused that package's previous test result.

For this three-package demo, a useful warm-run result is that `api` runs again
while `billing` and `database` are cached. Judge the test-command duration and
the packages that ran, not the entire job duration: runner startup, action
downloads, Go installation, and cache transfer are fixed overhead.

## What this demo can and cannot show

It can show cache isolation and correct Go test-cache invalidation. It cannot
validate a vendor's broad speedup claim: that depends on a real repository's
number of packages, test duration, change patterns, cache size, and runner
network speed. The native job demonstrates that CloudX is primarily a
convenient implementation of an explicit Actions cache policy, rather than a
new kind of Go compiler cache.

Treat a single warm run as evidence to investigate, not a correctness
guarantee. This repository's Actions history deliberately includes safety runs
that change an assertion and confirm the test executes. When evaluating this on
a production repository, use Go's `GODEBUG=gocachetest=1` diagnostic for a
short trial and verify that a changed package is not reported as cached.

For production, pin third-party Actions to reviewed commit SHAs rather than a
moving tag, and treat all caches as performance-only. Pass compiled deliverables
between jobs with artifacts, never via a cache.
