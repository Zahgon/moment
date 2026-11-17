package moment

import (
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

var (
	formattingTokens = jsre.MustCompile(`(\[[^\[]*\])|(\\e)|(\\)?(eHHmm|[Hh]mm(ss)?|Mo|MM?M?M?|Do|DDDo|DD?D?D?|ddd?d?|`+
		`do?|w[o|w]?|W[o|W]?|Qo?|N{1,5}|YYYYYY|YYYYY|YYYY|YY|y{2,4}|yo?|gg(ggg?)?|GG(GGG?)?|e|E|a|A|hh?|HH?|kk?|`+
		`mm?|ss?|S{1,9}|x|X|zz?|ZZ?|.)`, "g")
	localFormattingTokens = jsre.MustCompile(`(\[[^\[]*\])|(\\)?(LTS|LT|LL?L?L?|l{1,4})`, "g")
)

// tokenSpans splits UTF-16 text into the matches of formattingTokens; characters no alternative matches (line
// terminators) are dropped, as String.prototype.match with that expression does.
func tokenSpans(s []uint16) [][]uint16 {
	var out [][]uint16
	for from := 0; from <= len(s); {
		pos := formattingTokens.ExecUTF16(s, from)
		if pos == nil {
			break
		}
		out = append(out, s[pos[0]:pos[1]])
		if pos[1] == pos[0] {
			from = pos[1] + 1
		} else {
			from = pos[1]
		}
	}
	return out
}

// formatTokens is format.match(formattingTokens) as strings.
func formatTokens(format string) []string {
	spans := tokenSpans(js.UTF16(format))
	out := make([]string, len(spans))
	for i, s := range spans {
		out[i] = js.FromUTF16(s)
	}
	return out
}

// removeFormattingTokens turns a literal token into its text: brackets are unwrapped, backslashes dropped.
func removeFormattingTokens(tok []uint16) []uint16 {
	bracketed := false
	for i := 0; i+1 < len(tok); i++ {
		if tok[i] == '[' {
			bracketed = true
			break
		}
	}
	if bracketed {
		out := tok
		if len(out) > 0 && out[0] == '[' {
			out = out[1:]
		}
		if len(out) > 0 && out[len(out)-1] == ']' {
			out = out[:len(out)-1]
		}
		return out
	}
	out := make([]uint16, 0, len(tok))
	for _, c := range tok {
		if c != '\\' {
			out = append(out, c)
		}
	}
	return out
}

// formatMoment renders m with format after expanding the locale's long date formats.
func formatMoment(m *Moment, format string) string {
	if !m.IsValid() {
		return m.LocaleData().InvalidDate()
	}
	format = expandFormat(format, m.LocaleData())
	var out []uint16
	for _, tok := range tokenSpans(js.UTF16(format)) {
		if fn, ok := formatTokenFunctions[js.FromUTF16(tok)]; ok {
			out = append(out, js.UTF16(js.ToString(fn(m, format)))...)
		} else {
			out = append(out, removeFormattingTokens(tok)...)
		}
	}
	return js.FromUTF16(out)
}

// expandFormat replaces the long date format tokens (LT, LTS, L, LL, ...) with the locale's formats, up to six
// levels deep.
func expandFormat(format string, locale *Locale) string {
	for i := 5; i >= 0; i-- {
		localFormattingTokens.LastIndex = 0
		if !localFormattingTokens.Test(format) {
			break
		}
		format = localFormattingTokens.ReplaceFunc(format, func(match *jsre.Match) string {
			if f, ok := locale.LongDateFormat(match.Groups[0]).(string); ok && f != "" {
				return f
			}
			return match.Groups[0]
		})
	}
	localFormattingTokens.LastIndex = 0
	return format
}

// Format is moment#format: the default format when inputString is falsy, postformatted by the locale.
func (m *Moment) Format(inputString Value) string {
	if !js.Truthy(inputString) {
		if m.IsUTC() {
			inputString = DefaultFormatUtc
		} else {
			inputString = DefaultFormat
		}
	}
	return m.LocaleData().Postformat(formatMoment(m, js.ToString(inputString)))
}

// String is moment#toString.
func (m *Moment) String() string {
	return m.Clone().SetLocale("en").Format("ddd MMM DD YYYY HH:mm:ss [GMT]ZZ")
}

// ToISOString is moment#toISOString(keepOffset); it returns nil for an invalid moment.
func (m *Moment) ToISOString(keepOffset Value) (Value, error) {
	if !m.IsValid() {
		return nil, nil
	}
	utc := !js.StrictEquals(keepOffset, true)
	target := m
	if utc {
		target = m.Clone().UTC(js.Undef)
	}
	if y := target.Year(); y < 0 || y > 9999 {
		if utc {
			return formatMoment(target, "YYYYYY-MM-DD[T]HH:mm:ss.SSS[Z]"), nil
		}
		return formatMoment(target, "YYYYYY-MM-DD[T]HH:mm:ss.SSSZ"), nil
	}
	if utc {
		s, err := m.ToDate().ISOString()
		if err != nil {
			return nil, &ThrowError{Name: "RangeError", Message: err.Error()}
		}
		return s, nil
	}
	shifted := js.NewDate(m.ValueOf() + js.ToNumber(m.UTCOffset())*60*1000)
	s, err := shifted.ISOString()
	if err != nil {
		return nil, &ThrowError{Name: "RangeError", Message: err.Error()}
	}
	return js.ReplaceFirst(s, "Z", formatMoment(target, "Z")), nil
}

// Inspect is moment#inspect.
func (m *Moment) Inspect() string {
	if !m.IsValid() {
		return "moment.invalid(/* " + js.ToString(m.i) + " */)"
	}
	fn, zone := "moment", ""
	if !m.IsLocal() {
		if js.StrictEquals(m.UTCOffset(), 0.0) {
			fn = "moment.utc"
		} else {
			fn = "moment.parseZone"
		}
		zone = "Z"
	}
	year := "YYYYYY"
	if y := m.Year(); 0 <= y && y <= 9999 {
		year = "YYYY"
	}
	return m.Format("[" + fn + `("]` + year + "-MM-DD[T]HH:mm:ss.SSS" + zone + `[")]`)
}

// ToJSON is moment#toJSON.
func (m *Moment) ToJSON() (Value, error) {
	if !m.IsValid() {
		return nil, nil
	}
	return m.ToISOString(js.Undef)
}
