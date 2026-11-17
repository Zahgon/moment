package moment

import (
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

var (
	extendedISORegex = jsre.MustCompile(`^\s*((?:[+-]\d{6}|\d{4})-(?:\d\d-\d\d|W\d\d-\d|W\d\d|\d\d\d|\d\d))`+
		`(?:(T| )(\d\d(?::\d\d(?::\d\d(?:[.,]\d+)?)?)?)([+-]\d\d(?::?\d\d)?|\s*Z)?)?$`, "")
	basicISORegex = jsre.MustCompile(`^\s*((?:[+-]\d{6}|\d{4})(?:\d\d\d\d|W\d\d\d|W\d\d|\d\d\d|\d\d|))`+
		`(?:(T| )(\d\d(?:\d\d(?:\d\d(?:[.,]\d+)?)?)?)([+-]\d\d(?::?\d\d)?|\s*Z)?)?$`, "")
	tzRegex = jsre.MustCompile(`Z|[+-]\d\d(?::?\d\d)?`, "")

	isoDates = []isoFormat{
		{"YYYYYY-MM-DD", jsre.MustCompile(`[+-]\d{6}-\d\d-\d\d`, ""), true},
		{"YYYY-MM-DD", jsre.MustCompile(`\d{4}-\d\d-\d\d`, ""), true},
		{"GGGG-[W]WW-E", jsre.MustCompile(`\d{4}-W\d\d-\d`, ""), true},
		{"GGGG-[W]WW", jsre.MustCompile(`\d{4}-W\d\d`, ""), false},
		{"YYYY-DDD", jsre.MustCompile(`\d{4}-\d{3}`, ""), true},
		{"YYYY-MM", jsre.MustCompile(`\d{4}-\d\d`, ""), false},
		{"YYYYYYMMDD", jsre.MustCompile(`[+-]\d{10}`, ""), true},
		{"YYYYMMDD", jsre.MustCompile(`\d{8}`, ""), true},
		{"GGGG[W]WWE", jsre.MustCompile(`\d{4}W\d{3}`, ""), true},
		{"GGGG[W]WW", jsre.MustCompile(`\d{4}W\d{2}`, ""), false},
		{"YYYYDDD", jsre.MustCompile(`\d{7}`, ""), true},
		{"YYYYMM", jsre.MustCompile(`\d{6}`, ""), false},
		{"YYYY", jsre.MustCompile(`\d{4}`, ""), false},
	}
	isoTimes = []isoFormat{
		{"HH:mm:ss.SSSS", jsre.MustCompile(`\d\d:\d\d:\d\d\.\d+`, ""), true},
		{"HH:mm:ss,SSSS", jsre.MustCompile(`\d\d:\d\d:\d\d,\d+`, ""), true},
		{"HH:mm:ss", jsre.MustCompile(`\d\d:\d\d:\d\d`, ""), true},
		{"HH:mm", jsre.MustCompile(`\d\d:\d\d`, ""), true},
		{"HHmmss.SSSS", jsre.MustCompile(`\d\d\d\d\d\d\.\d+`, ""), true},
		{"HHmmss,SSSS", jsre.MustCompile(`\d\d\d\d\d\d,\d+`, ""), true},
		{"HHmmss", jsre.MustCompile(`\d\d\d\d\d\d`, ""), true},
		{"HHmm", jsre.MustCompile(`\d\d\d\d`, ""), true},
		{"HH", jsre.MustCompile(`\d\d`, ""), true},
	}
	aspNetJSONRegex = jsre.MustCompile(`^\/?Date\((-?\d+)`, "i")
	rfc2822         = jsre.MustCompile(`^(?:(Mon|Tue|Wed|Thu|Fri|Sat|Sun),?\s)?(\d{1,2})\s(Jan|Feb|Mar|Apr|May|Jun|Jul|`+
		`Aug|Sep|Oct|Nov|Dec)\s(\d{2,4})\s(\d\d):(\d\d)(?::(\d\d))?\s(?:(UT|GMT|[ECMP][SD]T)|([Zz])|([+-]\d{4}))$`, "")
	obsOffsets = map[string]float64{"UT": 0, "GMT": 0, "EDT": -4 * 60, "EST": -5 * 60, "CDT": -5 * 60,
		"CST": -6 * 60, "MDT": -6 * 60, "MST": -7 * 60, "PDT": -7 * 60, "PST": -8 * 60}
	rfcCommentRegex  = jsre.MustCompile(`\([^()]*\)|[\n\t]`, "g")
	rfcSpacesRegex   = jsre.MustCompile(`(\s\s+)`, "g")
	rfcLeadingRegex  = jsre.MustCompile(`^\s\s*`, "")
	rfcTrailingRegex = jsre.MustCompile(`\s\s*$`, "")
)

type isoFormat struct {
	format    string
	re        *jsre.Regexp
	allowTime bool
}

func group(m *jsre.Match, i int) (string, bool) {
	if m == nil || i >= len(m.Groups) || !m.Defined[i] {
		return "", false
	}
	return m.Groups[i], true
}

// configFromISO parses an ISO 8601 string.
func configFromISO(config *Moment) {
	s := js.ToString(config.i)
	match := extendedISORegex.Exec(s)
	if match == nil {
		match = basicISORegex.Exec(s)
	}
	if match == nil {
		config.isValid = false
		return
	}
	getParsingFlags(config).Set("iso", true)
	datePart, _ := group(match, 1)
	dateFormat, allowTime := "", false
	found := false
	for _, f := range isoDates {
		if f.re.Exec(datePart) != nil {
			dateFormat, allowTime, found = f.format, f.allowTime, true
			break
		}
	}
	if !found {
		config.isValid = false
		return
	}
	timeFormat, hasTime := "", false
	if t, ok := group(match, 3); ok && t != "" {
		sep, ok := group(match, 2)
		if !ok || sep == "" {
			sep = " "
		}
		for _, f := range isoTimes {
			if f.re.Exec(t) != nil {
				timeFormat, hasTime = sep+f.format, true
				break
			}
		}
		if !hasTime {
			config.isValid = false
			return
		}
	}
	if !allowTime && hasTime {
		config.isValid = false
		return
	}
	tzFormat := ""
	if tz, ok := group(match, 4); ok && tz != "" {
		if tzRegex.Exec(tz) == nil {
			config.isValid = false
			return
		}
		tzFormat = "Z"
	}
	config.f = dateFormat + timeFormat + tzFormat
	configFromStringAndFormat(config)
}

func untruncateYear(yearStr string) float64 {
	year := js.ParseInt(yearStr)
	if year <= 49 {
		return 2000 + year
	} else if year <= 999 {
		return 1900 + year
	}
	return year
}

func preprocessRFC2822(s string) string {
	s = rfcCommentRegex.Replace(s, " ")
	s = rfcSpacesRegex.Replace(s, " ")
	s = rfcLeadingRegex.Replace(s, "")
	return rfcTrailingRegex.Replace(s, "")
}

func indexOfString(list []string, s string) float64 {
	for i, v := range list {
		if v == s {
			return float64(i)
		}
	}
	return -1
}

// configFromRFC2822 parses an RFC 2822 date-time.
func configFromRFC2822(config *Moment) {
	match := rfc2822.Exec(preprocessRFC2822(js.ToString(config.i)))
	if match == nil {
		config.isValid = false
		return
	}
	str := func(i int) string { s, _ := group(match, i); return s }
	parsed := []float64{untruncateYear(str(4)), indexOfString(defaultLocaleMonthsShort, str(3)),
		js.ParseInt(str(2)), js.ParseInt(str(5)), js.ParseInt(str(6))}
	if sec, ok := group(match, 7); ok && sec != "" {
		parsed = append(parsed, js.ParseInt(sec))
	}
	if weekday, ok := group(match, 1); ok && weekday != "" {
		provided := indexOfString(defaultLocaleWeekdaysShort, weekday)
		actual := js.NewLocalDate(parsed[0], parsed[1], parsed[2]).Day(false)
		if provided != actual {
			getParsingFlags(config).Set("weekdayMismatch", true)
			config.isValid = false
			return
		}
	}
	config.a = make([]Value, len(parsed))
	for i, v := range parsed {
		config.a[i] = v
	}
	var offset float64
	if obs, ok := group(match, 8); ok && obs != "" {
		offset = obsOffsets[obs]
	} else if mil, ok := group(match, 9); ok && mil != "" {
		offset = 0
	} else {
		hm := js.ParseInt(str(10))
		m := js.Mod(hm, 100)
		h := (hm - m) / 100
		offset = h*60 + m
	}
	config.tzm = offset
	config.d = createUTCDate(parsed...)
	config.d.SetMinutes(true, config.d.Minutes(true)-offset)
	getParsingFlags(config).Set("rfc2822", true)
}

// configFromString tries ASP.NET JSON dates, ISO 8601 and RFC 2822, then falls back to the Date parser unless
// parsing is strict.
func configFromString(config *Moment) {
	if matched := aspNetJSONRegex.Exec(js.ToString(config.i)); matched != nil {
		config.d = js.NewDate(js.ToNumber(matched.Groups[1]))
		return
	}
	configFromISO(config)
	if js.StrictEquals(config.isValid, false) {
		config.isValid = js.Undef
	} else {
		return
	}
	configFromRFC2822(config)
	if js.StrictEquals(config.isValid, false) {
		config.isValid = js.Undef
	} else {
		return
	}
	if js.Truthy(config.strict) {
		config.isValid = false
	} else {
		createFromInputFallback(config)
	}
}
