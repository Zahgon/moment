# Agent Instructions

## Project Context

This is the maintained Moment.js 2.x codebase, ported to Go. It provides a
dependency-free date and time API with a large locale catalog and a long
compatibility history. The Go package keeps the JavaScript library's behavior,
including its quirks, so that results match Moment.js 2.31.0 byte for byte.
Favor small, conservative changes that preserve public behavior.

## Working In The Repository

The module is `github.com/moment/moment` and uses only the Go standard library.
Core behavior lives in the root package (`create.go`, `format.go`, `units.go`,
`duration.go`, `locale.go` and their neighbors). The JavaScript runtime semantics
the results depend on live in `internal/js` (numbers, strings, `Date` and local
time) and `internal/jsre` (ECMAScript regular expressions on UTF-16). Locale
definitions live in `locale/`, one file per locale. `cmd/momentcase` runs a case
file against the library and prints its results.

Recorded cases live under `cmd/momentcase/testdata/`; add regression cases near
the behavior being changed.

## Validation

Run `gofmt -l .` (it must print nothing), `go vet ./...` and `go test ./...`
before finishing. Time-zone dependent behavior must also pass with `TZ` set to
the zones of `.github/workflows/timezones.yml`.

## Code Conventions

Match surrounding code instead of introducing a new style. Moment has no runtime
dependencies; do not add one without an explicit design decision. Locale changes
require locale-specific cases and the evidence-based review described in
`CONTRIBUTING.md`.

Keep `AGENTS.md` focused on durable instructions for agents. Contributor-facing
guidance belongs in `CONTRIBUTING.md`, and release procedure belongs in
`RELEASING.md`; update those files when the corresponding behavior changes.

## Reviewing Locale Changes

When reviewing locale changes or locale pull requests, load and follow the
`locale-review` skill in `.agents/skills/locale-review/SKILL.md`. Treat that
skill as the authoritative locale-review workflow.
