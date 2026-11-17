# [Moment.js](https://momentjs.com/)

[![Go Reference][go-reference-image]][go-reference-url]
[![MIT License][license-image]][license-url]
[![Coverage Status][coveralls-image]][coveralls-url]
[![FOSSA Status][fossa-badge-image]][fossa-badge-url]

A date library for parsing, validating, manipulating, and formatting dates,
written in Go. It reproduces Moment.js 2.31.0, including its 139 locales and the
JavaScript `Date` behavior its results depend on.

> ⚠️ Moment.js is a legacy project, now in **maintenance mode**. New features
> are not being accepted. For more details, see the
> [latest project status update](https://momentjs.com/news/).

_Thank you._

Moment.js is a project of the [OpenJS Foundation](https://openjsf.org/), which
is part of the [Linux Foundation](https://www.linuxfoundation.org/).

## Resources

- [Documentation](https://momentjs.com/docs/)
- [Changelog](CHANGELOG.md)
- [Stack Overflow](https://stackoverflow.com/questions/tagged/momentjs)
- [Security Policy](SECURITY.md)
- [Threat Model](THREAT_MODEL.md)

## Installation

```bash
go get github.com/moment/moment
```

## Usage

The package mirrors the JavaScript API: `moment.New` is `moment(...)`,
`moment.CreateUTC` is `moment.utc(...)`, `moment.CreateDuration` is
`moment.duration(...)`, and moments, durations and locales have methods named
after their JavaScript counterparts. Arguments and results are JavaScript values
(see `moment.Value`). Import `github.com/moment/moment/locale` to make the
locales available.

```go
import (
	"fmt"

	"github.com/moment/moment"
	_ "github.com/moment/moment/locale"
)

func main() {
	m, err := moment.New("2010-10-20 4:30", "YYYY-MM-DD HH:mm")
	if err != nil {
		// What moment.js throws while parsing comes back as a *moment.ThrowError.
		panic(err)
	}
	fmt.Println(m.Format("LLLL"))
	fmt.Println(m.SetLocale("fr").Format("LLLL"))
}
```

Local time follows the `TZ` environment variable and the system time-zone
database, with the rules Node.js applies to times skipped or repeated by a
transition.

The `momentcase` command (`go build ./cmd/momentcase`) runs a case file, a JSON
program of calls, and prints one JSON line per result; it is the harness of the
recorded regression cases.

## License

Moment.js is an OpenJS Foundation project and is freely distributable under the
terms of the [MIT license][license-url].

<!-- markdownlint-disable MD033 -->
<a href="https://openjsf.org/">
  <picture>
    <source
      media="(prefers-color-scheme: dark)"
      srcset="https://raw.githubusercontent.com/openjs-foundation/artwork/main/openjs_foundation/openjs_foundation-logo-horizontal-color-dark_background.svg">
    <img
      src="https://raw.githubusercontent.com/openjs-foundation/artwork/main/openjs_foundation/openjs_foundation-logo-horizontal-color.svg"
      width="200"
      alt="OpenJS Foundation">
  </picture>
</a>
<!-- markdownlint-enable MD033 -->

Copyright [OpenJS Foundation](https://openjsf.org/) and Moment.js contributors.
All rights reserved. The [OpenJS Foundation](https://openjsf.org/) has registered
trademarks and uses trademarks. For a list of trademarks of the
[OpenJS Foundation](https://openjsf.org/), please see our
[Trademark Policy](https://trademark-policy.openjsf.org/) and
[Trademark List](https://trademark-list.openjsf.org/). Trademarks and logos not
indicated on the
[list of OpenJS Foundation trademarks](https://trademark-list.openjsf.org/) are
trademarks&trade; or registered&reg; trademarks of their respective holders. Use
of them does not imply any affiliation with or endorsement by them.

[The OpenJS Foundation](https://openjsf.org/) |
[AI Coding Assistants Policy][ai-policy] |
[Bylaws](https://bylaws.openjsf.org/) |
[Code of Conduct](https://code-of-conduct.openjsf.org/) |
[Cookie Policy](https://www.linuxfoundation.org/cookies/) |
[Privacy Policy](https://privacy-policy.openjsf.org/) |
[Terms of Use](https://terms-of-use.openjsf.org/) |
[Trademark List](https://trademark-list.openjsf.org/) |
[Trademark Policy](https://trademark-policy.openjsf.org/)

[license-image]: https://img.shields.io/badge/license-MIT-blue.svg?style=flat
[license-url]: LICENSE
[go-reference-image]: https://pkg.go.dev/badge/github.com/moment/moment.svg
[go-reference-url]: https://pkg.go.dev/github.com/moment/moment
[coveralls-image]: https://coveralls.io/repos/moment/moment/badge.svg?branch=develop
[coveralls-url]: https://coveralls.io/r/moment/moment?branch=develop
[fossa-badge-image]: https://app.fossa.io/api/projects/custom%2B20126%2Fgithub.com%2Fmoment%2Fmoment.svg?type=shield
[fossa-badge-url]: https://app.fossa.io/projects/custom%2B20126%2Fgithub.com%2Fmoment%2Fmoment?ref=badge_shield
[ai-policy]: https://ai-coding-assistants-policy.openjsf.org/
