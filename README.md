# Go cache Actions demo

This repository compares GitHub's `actions/setup-go` cache with
`cloudx-io/setup-go` under three parallel Go jobs: lint, test, and build.

The three package tests each take two seconds when uncached. Push a change to
`api/` and inspect the `Default cache — test` and `Isolated cache — test` job
logs to compare their restored Go test caches.
