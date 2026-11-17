// moment.js
// version : 2.31.0
// authors : Tim Wood, Iskren Chernev, Moment.js contributors
// license : MIT
// momentjs.com

// Package moment parses, validates, manipulates and formats dates. It is a port of moment.js 2.31.0 that keeps
// the library's JavaScript-facing behaviour, including its quirks: values are JavaScript values (see Value), local
// time follows the TZ environment variable with the rules of Node.js, and results match the original byte for
// byte.
package moment

import (
	"time"

	"github.com/moment/moment/internal/js"
)

// Version is moment.version.
const Version = "2.31.0"

// Value is a JavaScript value as the API accepts and returns it: nil (null), js.Undefined, bool, float64, string,
// []any or []string (arrays), *js.Object (plain objects), *js.Date, *js.Func, *Moment, *Duration or *Locale.
type Value = any

// Hooks are the replaceable globals of moment: the clock, the default formats, the two-digit year rule and the
// receivers of deprecation messages and console warnings.
var (
	// Now returns the current time in milliseconds (moment.now).
	Now = func() float64 { return float64(time.Now().UnixMilli()) }
	// DefaultFormat is moment.defaultFormat.
	DefaultFormat Value = "YYYY-MM-DDTHH:mm:ssZ"
	// DefaultFormatUtc is moment.defaultFormatUtc.
	DefaultFormatUtc Value = "YYYY-MM-DDTHH:mm:ss[Z]"
	// SuppressDeprecationWarnings stops deprecation messages from reaching Warn.
	SuppressDeprecationWarnings = false
	// DeprecationHandler receives every deprecation (name is nil for deprecated functions) when set.
	DeprecationHandler func(name Value, msg string)
	// Warn receives what moment.js writes with console.warn.
	Warn = func(string) {}
)

// ParseTwoDigitYear is moment.parseTwoDigitYear: 00 to 68 are 2000 to 2068, 69 to 99 are 1969 to 1999.
func ParseTwoDigitYear(input Value) float64 {
	n := toInt(input)
	if n > 68 {
		return n + 1900
	}
	return n + 2000
}

// ThrowError is a JavaScript exception raised by the library: Name is "Error" for moment's own errors and the
// engine's error name (TypeError, RangeError) for operations JavaScript rejects.
type ThrowError struct {
	Name    string
	Message string
}

func (e *ThrowError) Error() string { return e.Name + ": " + e.Message }

func typeError(msg string) *ThrowError { return &ThrowError{Name: "TypeError", Message: msg} }

// deprecation reports a deprecated function: the handler hears every call, the warning is written once.
type deprecation struct {
	msg       string
	firstTime bool
}

func newDeprecation(msg string) *deprecation { return &deprecation{msg: msg, firstTime: true} }

func (d *deprecation) fire() {
	if DeprecationHandler != nil {
		DeprecationHandler(nil, d.msg)
	}
	if d.firstTime {
		warnDeprecation(d.msg)
		d.firstTime = false
	}
}

var deprecationsSeen = map[string]bool{}

// deprecateSimple reports a named deprecation (moment's deprecateSimple).
func deprecateSimple(name, msg string) {
	if DeprecationHandler != nil {
		DeprecationHandler(name, msg)
	}
	if !deprecationsSeen[name] {
		warnDeprecation(msg)
		deprecationsSeen[name] = true
	}
}

func warnDeprecation(msg string) {
	if !SuppressDeprecationWarnings {
		Warn("Deprecation warning: " + msg)
	}
}

var (
	fallbackDeprecation = newDeprecation("value provided is not in a recognized RFC2822 or ISO format. moment " +
		"construction falls back to js Date(), which is not reliable across all browsers and versions. Non " +
		"RFC2822/ISO date formats are discouraged. Please refer to http://momentjs.com/guides/#/warnings/js-date/ " +
		"for more info.")
	langDeprecation      = newDeprecation("moment.lang is deprecated. Use moment.locale instead.")
	langDataDeprecation  = newDeprecation("moment.langData is deprecated. Use moment.localeData instead.")
	momentLangDeprecated = newDeprecation("moment().lang() is deprecated. Instead, use moment().localeData() to " +
		"get the language configuration. Use moment().locale() to change languages.")
	minDeprecation    = newDeprecation("moment().min is deprecated, use moment.max instead. http://momentjs.com/guides/#/warnings/min-max/")
	maxDeprecation    = newDeprecation("moment().max is deprecated, use moment.min instead. http://momentjs.com/guides/#/warnings/min-max/")
	datesDeprecation  = newDeprecation("dates accessor is deprecated. Use date instead.")
	monthsDeprecation = newDeprecation("months accessor is deprecated. Use month instead")
	yearsDeprecation  = newDeprecation("years accessor is deprecated. Use year instead")
	zoneDeprecation   = newDeprecation("moment().zone is deprecated, use moment().utcOffset instead. " +
		"http://momentjs.com/guides/#/warnings/zone/")
	dstShiftedDeprecation = newDeprecation("isDSTShifted is deprecated. See " +
		"http://momentjs.com/guides/#/warnings/dst-shifted/ for more information")
	toIsoStringDeprecation = newDeprecation("toIsoString() is deprecated. Please use toISOString() instead " +
		"(notice the capitals)")
)

// createFromInputFallback hands strings no other parser accepts to the JavaScript Date parser.
func createFromInputFallback(config *Moment) {
	fallbackDeprecation.fire()
	s := js.ToString(config.i)
	if js.Truthy(config.useUTC) {
		s += " UTC"
	}
	config.d = js.ParseDate(s)
}
