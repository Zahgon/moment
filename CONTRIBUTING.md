# Contributing to Moment.js

Moment.js has been a Linux Foundation-hosted project since joining the JS
Foundation in October 2016. The JS Foundation later merged with the Node.js
Foundation to form the [OpenJS Foundation](https://openjsf.org/), which hosts
the project today.

## Submitting Issues

Before submitting an issue, check the existing issues and make sure it has not
already been reported. Bug reports should include a minimal reproduction with
the exact actual and expected output.

Moment.js is in maintenance mode and is not accepting new features. Use the
provided issue forms for reproducible defects and locale corrections. For usage
questions and other issue categories, follow the links in the issue chooser.

Maintainers assess reports individually and may choose not to fix a valid defect
when compatibility risk, established behavior, severity, or available capacity
outweighs the benefit of a change. Submitting a report does not guarantee that a
code change will be made. Technical modernization is planned and initiated by
maintainers; unsolicited modernization proposals may be closed.

## Read Before Submitting Pull Requests

- **Pull requests to the `master` branch will be closed.** Please submit all
  pull requests to the `develop` branch.
- **You will be required to sign the OpenJS Foundation CLA before your pull
  request can be merged.** The EasyCLA bot will prompt you, or you can
  [sign it now](https://api.easycla.lfx.linuxfoundation.org/v2/repository-provider/github/sign/6354452/235430945/1).
- **Locale translations will not be merged without regression cases.** See
  the recorded cases under `cmd/momentcase/testdata/` for examples.

## Code Organization

The library is a Go module that uses only the standard library.

- `*.go` at the root: the moment core (creation and parsing, formatting, units,
  durations, locales)
- `internal/js`: the JavaScript number, string and `Date` semantics the results
  depend on
- `internal/jsre`: ECMAScript regular expressions matched on UTF-16 code units
- `locale/*.go`: locale definitions, one file per locale
- `cmd/momentcase`: the case runner and its recorded cases

Follow the style of the surrounding code and keep it `gofmt`-formatted.

## Setting Up the Development Environment

To contribute, fork the library. You need [git](http://git-scm.com/) and
[Go](https://go.dev/dl/) 1.23 or newer.

```bash
git clone https://github.com/moment/moment.git
cd moment
git checkout develop  # all patches against develop branch, please!
gofmt -l .            # must print nothing
go vet ./...
go test ./...
```

## Changing Locale Files

Locale changes are reviewed first against the latest available data in the
official [Unicode CLDR repository](https://github.com/unicode-org/cldr), which is
the project's preferred baseline. A change that differs from CLDR may still be
accepted when strong evidence shows that CLDR is stale, inaccurate, or does not
reflect the locale's documented convention.

Please describe the convention behind the change and include supporting sources
when available. Useful references include government or institutional style
guides, national standards, language-academy dictionaries, official calendars,
and other evidence of contemporary usage. Sources are requested but not
required; maintainers will independently validate the proposed behavior.
Maintainers are not necessarily native speakers. Locale contributions are
reviewed by a maintainer using the AI-assisted workflow defined in the
[`locale-review` skill](.agents/skills/locale-review/SKILL.md). AI output is not
treated as evidence; review conclusions remain grounded in CLDR and
authoritative external sources.

If the same behavior was previously discussed or a similar change was declined,
address the prior decision and explain what new context or evidence supports
reconsidering it.

Make locale changes in `locale/<locale>.go` and add or update the matching
cases under `cmd/momentcase/testdata/` for each distinct form affected.

## Development Commands

- `go test ./...` Run the tests.
- `go test -cover ./...` Run the tests with a coverage summary.
- `go vet ./...` and `gofmt -l .` Perform static and formatting checks.
- `go test -run '^$' -bench . .` Run the benchmarks.
- `git config core.hooksPath .githooks` Install safeguards against development
  on `master`.
- `go build ./cmd/momentcase` Build the case runner.

Release preparation is documented in [RELEASING.md](RELEASING.md).
