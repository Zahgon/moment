package moment

import (
	"math"
	"sort"
	"strings"

	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

var (
	defaultLocaleMonths = []string{"January", "February", "March", "April", "May", "June", "July", "August",
		"September", "October", "November", "December"}
	defaultLocaleMonthsShort   = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	defaultLocaleWeekdays      = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	defaultLocaleWeekdaysShort = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	defaultLocaleWeekdaysMin   = []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	monthsInFormat             = jsre.MustCompile(`D[oD]?(\[[^\[\]]*\]|\s)+MMMM?`, "")
	percentD                   = jsre.MustCompile(`%d`, "i")
	percentS                   = jsre.MustCompile(`%s`, "i")
	monthsParseProperties      = []string{"monthsParse", "longMonthsParse", "shortMonthsParse", "monthsRegex",
		"monthsShortRegex", "monthsStrictRegex", "monthsShortStrictRegex"}
	weekdaysParseProperties = []string{"weekdaysParse", "fullWeekdaysParse", "shortWeekdaysParse",
		"minWeekdaysParse", "weekdaysRegex", "weekdaysShortRegex", "weekdaysMinRegex", "weekdaysStrictRegex",
		"weekdaysShortStrictRegex", "weekdaysMinStrictRegex"}

	baseConfig = js.Obj(
		"calendar", js.Obj("sameDay", "[Today at] LT", "nextDay", "[Tomorrow at] LT", "nextWeek", "dddd [at] LT",
			"lastDay", "[Yesterday at] LT", "lastWeek", "[Last] dddd [at] LT", "sameElse", "L"),
		"longDateFormat", js.Obj("LTS", "h:mm:ss A", "LT", "h:mm A", "L", "MM/DD/YYYY", "LL", "MMMM D, YYYY",
			"LLL", "MMMM D, YYYY h:mm A", "LLLL", "dddd, MMMM D, YYYY h:mm A"),
		"invalidDate", "Invalid date",
		"ordinal", "%d",
		"dayOfMonthOrdinalParse", jsre.MustCompile(`\d{1,2}`, ""),
		"relativeTime", js.Obj("future", "in %s", "past", "%s ago", "s", "a few seconds", "ss", "%d seconds",
			"m", "a minute", "mm", "%d minutes", "h", "an hour", "hh", "%d hours", "d", "a day", "dd", "%d days",
			"w", "a week", "ww", "%d weeks", "M", "a month", "MM", "%d months", "y", "a year", "yy", "%d years"),
		"months", defaultLocaleMonths,
		"monthsShort", defaultLocaleMonthsShort,
		"week", js.Obj("dow", 0.0, "doy", 6.0),
		"weekdays", defaultLocaleWeekdays,
		"weekdaysMin", defaultLocaleWeekdaysMin,
		"weekdaysShort", defaultLocaleWeekdaysShort,
		"meridiemParse", jsre.MustCompile(`[ap]\.?m?\.?`, "i"),
	)
)

func init() {
	GetSetGlobalLocale("en", js.Obj(
		"eras", []any{
			js.Obj("since", "0001-01-01", "until", math.Inf(1), "offset", 1.0, "name", "Anno Domini",
				"narrow", "AD", "abbr", "AD"),
			js.Obj("since", "0000-12-31", "until", math.Inf(-1), "offset", 1.0, "name", "Before Christ",
				"narrow", "BC", "abbr", "BC"),
		},
		"dayOfMonthOrdinalParse", jsre.MustCompile(`\d{1,2}(th|st|nd|rd)`, ""),
		"ordinal", englishOrdinal,
	))
}

// englishOrdinal is the ordinal callback of the base English locale; like every ordinal callback it receives the
// argument as it was passed and returns a JavaScript value.
func englishOrdinal(_ *Locale, number Value, _ string) Value {
	n := js.ToNumber(number)
	b := math.Mod(n, 10)
	suffix := "th"
	if js.ToInt32(math.Mod(n, 100)/10) != 1 {
		switch b {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		}
	}
	return js.ToString(number) + suffix
}

func stringList(v Value) ([]string, bool) {
	switch x := v.(type) {
	case []string:
		return x, true
	case []any:
		out := make([]string, len(x))
		for i, e := range x {
			out[i] = js.ToString(e)
		}
		return out, true
	}
	return nil, false
}

func listValue(v Value, i float64) Value {
	switch x := v.(type) {
	case []string:
		if i >= 0 && i < float64(len(x)) && i == math.Trunc(i) {
			return x[int(i)]
		}
	case []any:
		if i >= 0 && i < float64(len(x)) && i == math.Trunc(i) {
			return x[int(i)]
		}
	}
	return js.Undef
}

func listAsValue(v Value) Value {
	if s, ok := v.([]string); ok {
		return append([]string(nil), s...)
	}
	return v
}

// Calendar is Locale#calendar(key, mom, now): the calendar format for key, from a function when the locale has
// one.
func (l *Locale) Calendar(key Value, mom *Moment, now *Moment) Value {
	if f, ok := l.own.Get("calendar").(func(*Locale, string, *Moment, *Moment) string); ok {
		return f(l, js.ToString(key), mom, now)
	}
	cal, _ := l.own.Get("_calendar").(*js.Object)
	if cal == nil {
		return js.Undef
	}
	output := cal.Get(js.ToString(key))
	if !js.Truthy(output) {
		output = cal.Get("sameElse")
	}
	if f, ok := output.(func(*Moment, *Moment) string); ok {
		return f(mom, now)
	}
	return output
}

// LongDateFormat is Locale#longDateFormat(key); a lower-case key derives its format from the upper-case one.
func (l *Locale) LongDateFormat(key Value) Value {
	formats, _ := l.own.Get("_longDateFormat").(*js.Object)
	if formats == nil {
		return js.Undef
	}
	k := js.ToString(key)
	format := formats.Get(k)
	upper := formats.Get(js.ToUpper(k))
	if js.Truthy(format) || !js.Truthy(upper) {
		return format
	}
	cache, _ := l.own.Get("_longDateFormatCache").(*js.Object)
	if cache != nil {
		if entry, ok := cache.Get(k).(*js.Object); ok && js.StrictEquals(entry.Get("formatUpper"), upper) {
			return entry.Get("format")
		}
	}
	var b strings.Builder
	for _, tok := range formatTokens(js.ToString(upper)) {
		if tok == "MMMM" || tok == "MM" || tok == "DD" || tok == "dddd" {
			tok = tok[1:]
		}
		b.WriteString(tok)
	}
	if cache == nil {
		cache = js.NewObject()
		l.own.Set("_longDateFormatCache", cache)
	}
	cache.Set(k, js.Obj("formatUpper", upper, "format", b.String()))
	return b.String()
}

// InvalidDate is Locale#invalidDate.
func (l *Locale) InvalidDate() string { return js.ToString(l.own.Get("_invalidDate")) }

// Ordinal is Locale#ordinal(number, token).
func (l *Locale) Ordinal(number, token Value) Value {
	if f, ok := l.own.Get("ordinal").(func(*Locale, Value, string) Value); ok {
		return f(l, number, js.ToString(token))
	}
	return js.ReplaceFirst(js.ToString(l.own.Get("_ordinal")), "%d", js.ToString(number))
}

// Preparse is Locale#preparse.
func (l *Locale) Preparse(s string) string {
	if f, ok := l.own.Get("preparse").(func(*Locale, string) string); ok {
		return f(l, s)
	}
	return s
}

// Postformat is Locale#postformat.
func (l *Locale) Postformat(s string) string {
	if f, ok := l.own.Get("postformat").(func(*Locale, string) string); ok {
		return f(l, s)
	}
	return s
}

func (l *Locale) relativeTimeWithoutPostformat(number float64, withoutSuffix bool, key string, isFuture bool) string {
	rt, _ := l.own.Get("_relativeTime").(*js.Object)
	if rt == nil {
		return ""
	}
	output := rt.Get(key)
	if f, ok := output.(func(float64, bool, string, bool) string); ok {
		return f(number, withoutSuffix, key, isFuture)
	}
	return percentD.Replace(js.ToString(output), js.NumberToString(number))
}

// RelativeTime is Locale#relativeTime(number, withoutSuffix, key, isFuture).
func (l *Locale) RelativeTime(number, withoutSuffix, key, isFuture Value) string {
	return l.Postformat(l.relativeTimeWithoutPostformat(js.ToNumber(number), js.Truthy(withoutSuffix),
		js.ToString(key), js.Truthy(isFuture)))
}

func (l *Locale) pastFutureWithoutPostformat(diff float64, output string) string {
	rt, _ := l.own.Get("_relativeTime").(*js.Object)
	if rt == nil {
		return output
	}
	key := "past"
	if diff > 0 {
		key = "future"
	}
	format := rt.Get(key)
	if f, ok := format.(func(string) string); ok {
		return f(output)
	}
	return percentS.Replace(js.ToString(format), output)
}

// PastFuture is Locale#pastFuture(diff, output).
func (l *Locale) PastFuture(diff Value, output Value) string {
	return l.Postformat(l.pastFutureWithoutPostformat(js.ToNumber(diff), js.ToString(output)))
}

type listFunc = func(*Locale, *Moment, any) any

// Months is Locale#months(m, format): the month name of m (the genitive form where the format calls for it), or
// every name when m is nil.
func (l *Locale) Months(m *Moment, format Value) Value {
	if f, ok := l.own.Get("months").(listFunc); ok {
		return f(l, m, format)
	}
	months := l.own.Get("_months")
	if m == nil {
		if o, ok := months.(*js.Object); ok {
			return listAsValue(o.Get("standalone"))
		}
		return listAsValue(months)
	}
	if o, ok := months.(*js.Object); ok {
		test := monthsInFormat
		if re, ok := o.Get("isFormat").(*jsre.Regexp); ok {
			test = re
		}
		key := "standalone"
		if test.Test(js.ToString(format)) {
			key = "format"
		}
		return listValue(o.Get(key), m.Month())
	}
	return listValue(months, m.Month())
}

// MonthsShort is Locale#monthsShort(m, format).
func (l *Locale) MonthsShort(m *Moment, format Value) Value {
	if f, ok := l.own.Get("monthsShort").(listFunc); ok {
		return f(l, m, format)
	}
	months := l.own.Get("_monthsShort")
	if m == nil {
		if o, ok := months.(*js.Object); ok {
			return listAsValue(o.Get("standalone"))
		}
		return listAsValue(months)
	}
	if o, ok := months.(*js.Object); ok {
		key := "standalone"
		if monthsInFormat.Test(js.ToString(format)) {
			key = "format"
		}
		return listValue(o.Get(key), m.Month())
	}
	return listValue(months, m.Month())
}

func (l *Locale) cacheList(key string) ([]any, bool) {
	v := l.own.Get(key)
	switch x := v.(type) {
	case []any:
		return x, true
	case []string:
		return asArray(x), true
	case []*jsre.Regexp:
		out := make([]any, len(x))
		for i, r := range x {
			out[i] = r
		}
		return out, true
	}
	return nil, js.Truthy(v)
}

func (l *Locale) setCacheItem(key string, i int, v Value) {
	list, _ := l.cacheList(key)
	for len(list) <= i {
		list = append(list, js.Undef)
	}
	list[i] = v
	l.own.Set(key, list)
}

func cacheItem(list []any, i int) Value {
	if i < len(list) {
		return list[i]
	}
	return js.Undef
}

func indexOfValue(list []any, s string) float64 {
	for i, v := range list {
		if js.StrictEquals(v, s) {
			return float64(i)
		}
	}
	return -1
}

func compileOrNil(source, flags string) *jsre.Regexp {
	re, err := jsre.Compile(source, flags)
	if err != nil {
		return nil
	}
	return re
}

func testRegex(v Value, s string) bool {
	re, ok := v.(*jsre.Regexp)
	return ok && re.Test(s)
}

// strictParse looks name up, lower-cased, in the lower-cased lists named by order; the first hit wins.
func (l *Locale) strictParse(name string, order []string) Value {
	llc := js.ToLower(name)
	for _, key := range order {
		list, _ := l.cacheList(key)
		if i := indexOfValue(list, llc); i != -1 {
			return i
		}
	}
	return nil
}

// MonthsParse is Locale#monthsParse(monthName, format, strict): the month index, or null/undefined.
func (l *Locale) MonthsParse(monthName string, format Value, strict bool) Value {
	if js.Truthy(l.own.Get("_monthsParseExact")) {
		if _, ok := l.cacheList("_monthsParse"); !ok {
			l.own.Set("_monthsParse", []any{})
			short, long := make([]any, 12), make([]any, 12)
			for i := 0; i < 12; i++ {
				mom := utcMoment([]any{2000.0, float64(i)})
				short[i] = js.ToLower(js.ToString(l.MonthsShort(mom, "")))
				long[i] = js.ToLower(js.ToString(l.Months(mom, "")))
			}
			l.own.Set("_shortMonthsParse", short)
			l.own.Set("_longMonthsParse", long)
		}
		f := js.ToString(format)
		switch {
		case strict && f == "MMM":
			return l.strictParse(monthName, []string{"_shortMonthsParse"})
		case strict:
			return l.strictParse(monthName, []string{"_longMonthsParse"})
		case f == "MMM":
			return l.strictParse(monthName, []string{"_shortMonthsParse", "_longMonthsParse"})
		}
		return l.strictParse(monthName, []string{"_longMonthsParse", "_shortMonthsParse"})
	}
	if _, ok := l.cacheList("_monthsParse"); !ok {
		l.own.Set("_monthsParse", []any{})
		l.own.Set("_longMonthsParse", []any{})
		l.own.Set("_shortMonthsParse", []any{})
	}
	f := js.ToString(format)
	for i := 0; i < 12; i++ {
		mom := utcMoment([]any{2000.0, float64(i)})
		long, _ := l.cacheList("_longMonthsParse")
		if strict && !js.Truthy(cacheItem(long, i)) {
			l.setCacheItem("_longMonthsParse", i, compileOrNil(
				"^"+js.ReplaceFirst(js.ToString(l.Months(mom, "")), ".", "")+"$", "i"))
			l.setCacheItem("_shortMonthsParse", i, compileOrNil(
				"^"+js.ReplaceFirst(js.ToString(l.MonthsShort(mom, "")), ".", "")+"$", "i"))
		}
		parse, _ := l.cacheList("_monthsParse")
		if !strict && !js.Truthy(cacheItem(parse, i)) {
			source := "^" + js.ToString(l.Months(mom, "")) + "|^" + js.ToString(l.MonthsShort(mom, ""))
			l.setCacheItem("_monthsParse", i, compileOrNil(js.ReplaceFirst(source, ".", ""), "i"))
		}
		long, _ = l.cacheList("_longMonthsParse")
		short, _ := l.cacheList("_shortMonthsParse")
		parse, _ = l.cacheList("_monthsParse")
		switch {
		case strict && f == "MMMM" && testRegex(cacheItem(long, i), monthName):
			return float64(i)
		case strict && f == "MMM" && testRegex(cacheItem(short, i), monthName):
			return float64(i)
		case !strict && testRegex(cacheItem(parse, i), monthName):
			return float64(i)
		}
	}
	return js.Undef
}

func sortByLengthDesc(pieces []string) {
	sort.SliceStable(pieces, func(i, j int) bool { return js.Len16(pieces[i]) > js.Len16(pieces[j]) })
}

func (l *Locale) computeMonthsParse() {
	var short, long, mixed []string
	for i := 0; i < 12; i++ {
		mom := utcMoment([]any{2000.0, float64(i)})
		s := regexEscape(js.ToString(l.MonthsShort(mom, "")))
		lg := regexEscape(js.ToString(l.Months(mom, "")))
		short = append(short, s)
		long = append(long, lg)
		mixed = append(mixed, lg, s)
	}
	sortByLengthDesc(short)
	sortByLengthDesc(long)
	sortByLengthDesc(mixed)
	l.own.Set("_monthsRegex", compileOrNil("^("+strings.Join(mixed, "|")+")", "i"))
	l.own.Set("_monthsShortRegex", l.own.Get("_monthsRegex"))
	l.own.Set("_monthsStrictRegex", compileOrNil("^("+strings.Join(long, "|")+")", "i"))
	l.own.Set("_monthsShortStrictRegex", compileOrNil("^("+strings.Join(short, "|")+")", "i"))
}

func regexProp(v Value) *jsre.Regexp {
	re, _ := v.(*jsre.Regexp)
	return re
}

// monthsRegexFor implements monthsRegex and monthsShortRegex; prefix is "_months" or "_monthsShort".
func (l *Locale) monthsRegexFor(prefix string, strict bool) *jsre.Regexp {
	if js.Truthy(l.own.Get("_monthsParseExact")) {
		if !l.own.Has("_monthsRegex") {
			l.computeMonthsParse()
		}
		if strict {
			return regexProp(l.own.Get(prefix + "StrictRegex"))
		}
		return regexProp(l.own.Get(prefix + "Regex"))
	}
	if !l.own.Has(prefix + "Regex") {
		l.own.Set(prefix+"Regex", matchWord)
	}
	if js.Truthy(l.own.Get(prefix+"StrictRegex")) && strict {
		return regexProp(l.own.Get(prefix + "StrictRegex"))
	}
	return regexProp(l.own.Get(prefix + "Regex"))
}

// MonthsRegex is Locale#monthsRegex(strict).
func (l *Locale) MonthsRegex(strict bool) *jsre.Regexp { return l.monthsRegexFor("_months", strict) }

// MonthsShortRegex is Locale#monthsShortRegex(strict).
func (l *Locale) MonthsShortRegex(strict bool) *jsre.Regexp {
	return l.monthsRegexFor("_monthsShort", strict)
}

// Week is Locale#week(m): the week of the year of m in this locale's week rules.
func (l *Locale) Week(m *Moment) float64 {
	return weekOfYear(m, js.ToNumber(l.weekSetting("dow")), js.ToNumber(l.weekSetting("doy"))).week
}

// FirstDayOfWeek is Locale#firstDayOfWeek.
func (l *Locale) FirstDayOfWeek() Value { return l.weekSetting("dow") }

// FirstDayOfYear is Locale#firstDayOfYear.
func (l *Locale) FirstDayOfYear() Value { return l.weekSetting("doy") }

func shiftWeekdays(ws Value, n float64) Value {
	list, ok := stringList(ws)
	if !ok {
		return js.Undef
	}
	k := int(js.ToIntegerOrInfinity(n))
	if k < 0 {
		k = max(0, len(list)+k)
	}
	k = min(k, len(list))
	end := min(7, len(list))
	if k > end {
		k = end
	}
	out := append([]string{}, list[k:end]...)
	return append(out, list[:k]...)
}

// WeekdayArg is the first argument of the weekday list methods: nil for none, a moment, or true for the list
// sorted by the locale's first day of the week.
type WeekdayArg = Value

// Weekdays is Locale#weekdays(m, format).
func (l *Locale) Weekdays(m WeekdayArg, format Value) Value {
	if f, ok := l.own.Get("weekdays").(listFunc); ok {
		mom, _ := m.(*Moment)
		if b, isBool := m.(bool); isBool && b {
			// A weekdays function receives the JavaScript `true` (the locale-sorted list) in place of format.
			format = true
		}
		return f(l, mom, format)
	}
	weekdays := l.own.Get("_weekdays")
	if o, ok := weekdays.(*js.Object); ok {
		key := "standalone"
		if mom, isM := m.(*Moment); isM && mom != nil && testRegex(o.Get("isFormat"), js.ToString(format)) {
			key = "format"
		}
		weekdays = o.Get(key)
	}
	return l.pickWeekday(weekdays, m)
}

func (l *Locale) pickWeekday(list Value, m WeekdayArg) Value {
	if b, ok := m.(bool); ok && b {
		return shiftWeekdays(list, js.ToNumber(l.weekSetting("dow")))
	}
	if mom, ok := m.(*Moment); ok && mom != nil {
		return listValue(list, mom.Day())
	}
	return listAsValue(list)
}

// WeekdaysShort is Locale#weekdaysShort(m).
func (l *Locale) WeekdaysShort(m WeekdayArg) Value {
	return l.pickWeekday(l.own.Get("_weekdaysShort"), m)
}

// WeekdaysMin is Locale#weekdaysMin(m).
func (l *Locale) WeekdaysMin(m WeekdayArg) Value {
	return l.pickWeekday(l.own.Get("_weekdaysMin"), m)
}

func weekdayMoment(i int) *Moment {
	m := utcMoment([]any{2000.0, 1.0})
	m.GetSetDay(float64(i))
	return m
}

// WeekdaysParse is Locale#weekdaysParse(weekdayName, format, strict).
func (l *Locale) WeekdaysParse(name string, format Value, strict bool) Value {
	f := js.ToString(format)
	if js.Truthy(l.own.Get("_weekdaysParseExact")) {
		if _, ok := l.cacheList("_weekdaysParse"); !ok {
			full, short, minimal := make([]any, 7), make([]any, 7), make([]any, 7)
			for i := 0; i < 7; i++ {
				mom := weekdayMoment(i)
				minimal[i] = js.ToLower(js.ToString(l.WeekdaysMin(mom)))
				short[i] = js.ToLower(js.ToString(l.WeekdaysShort(mom)))
				full[i] = js.ToLower(js.ToString(l.Weekdays(mom, "")))
			}
			l.own.Set("_weekdaysParse", full)
			l.own.Set("_shortWeekdaysParse", short)
			l.own.Set("_minWeekdaysParse", minimal)
		}
		order := map[string][]string{
			"dddd": {"_weekdaysParse", "_shortWeekdaysParse", "_minWeekdaysParse"},
			"ddd":  {"_shortWeekdaysParse", "_weekdaysParse", "_minWeekdaysParse"},
			"":     {"_minWeekdaysParse", "_weekdaysParse", "_shortWeekdaysParse"},
		}
		key := f
		if key != "dddd" && key != "ddd" {
			key = ""
		}
		if strict {
			return l.strictParse(name, order[key][:1])
		}
		return l.strictParse(name, order[key])
	}
	if _, ok := l.cacheList("_weekdaysParse"); !ok {
		for _, k := range []string{"_weekdaysParse", "_minWeekdaysParse", "_shortWeekdaysParse", "_fullWeekdaysParse"} {
			l.own.Set(k, []any{})
		}
	}
	for i := 0; i < 7; i++ {
		mom := weekdayMoment(i)
		full, _ := l.cacheList("_fullWeekdaysParse")
		if strict && !js.Truthy(cacheItem(full, i)) {
			l.setCacheItem("_fullWeekdaysParse", i, compileOrNil(
				"^"+js.ReplaceFirst(js.ToString(l.Weekdays(mom, "")), ".", `\.?`)+"$", "i"))
			l.setCacheItem("_shortWeekdaysParse", i, compileOrNil(
				"^"+js.ReplaceFirst(js.ToString(l.WeekdaysShort(mom)), ".", `\.?`)+"$", "i"))
			l.setCacheItem("_minWeekdaysParse", i, compileOrNil(
				"^"+js.ReplaceFirst(js.ToString(l.WeekdaysMin(mom)), ".", `\.?`)+"$", "i"))
		}
		parse, _ := l.cacheList("_weekdaysParse")
		if !js.Truthy(cacheItem(parse, i)) {
			source := "^" + js.ToString(l.Weekdays(mom, "")) + "|^" + js.ToString(l.WeekdaysShort(mom)) + "|^" +
				js.ToString(l.WeekdaysMin(mom))
			l.setCacheItem("_weekdaysParse", i, compileOrNil(js.ReplaceFirst(source, ".", ""), "i"))
		}
		full, _ = l.cacheList("_fullWeekdaysParse")
		short, _ := l.cacheList("_shortWeekdaysParse")
		minimal, _ := l.cacheList("_minWeekdaysParse")
		parse, _ = l.cacheList("_weekdaysParse")
		switch {
		case strict && f == "dddd" && testRegex(cacheItem(full, i), name):
			return float64(i)
		case strict && f == "ddd" && testRegex(cacheItem(short, i), name):
			return float64(i)
		case strict && f == "dd" && testRegex(cacheItem(minimal, i), name):
			return float64(i)
		case !strict && testRegex(cacheItem(parse, i), name):
			return float64(i)
		}
	}
	return js.Undef
}

func (l *Locale) computeWeekdaysParse() {
	var minimal, short, long, mixed []string
	for i := 0; i < 7; i++ {
		mom := weekdayMoment(i)
		mn := regexEscape(js.ToString(l.WeekdaysMin(mom)))
		sh := regexEscape(js.ToString(l.WeekdaysShort(mom)))
		lg := regexEscape(js.ToString(l.Weekdays(mom, "")))
		minimal = append(minimal, mn)
		short = append(short, sh)
		long = append(long, lg)
		mixed = append(mixed, mn, sh, lg)
	}
	for _, p := range [][]string{minimal, short, long, mixed} {
		sortByLengthDesc(p)
	}
	mixedRe := compileOrNil("^("+strings.Join(mixed, "|")+")", "i")
	l.own.Set("_weekdaysRegex", mixedRe)
	l.own.Set("_weekdaysShortRegex", mixedRe)
	l.own.Set("_weekdaysMinRegex", mixedRe)
	l.own.Set("_weekdaysStrictRegex", compileOrNil("^("+strings.Join(long, "|")+")", "i"))
	l.own.Set("_weekdaysShortStrictRegex", compileOrNil("^("+strings.Join(short, "|")+")", "i"))
	l.own.Set("_weekdaysMinStrictRegex", compileOrNil("^("+strings.Join(minimal, "|")+")", "i"))
}

// weekdaysRegexFor implements the three weekday regex methods; prefix is "_weekdays", "_weekdaysShort" or
// "_weekdaysMin".
func (l *Locale) weekdaysRegexFor(prefix string, strict bool) *jsre.Regexp {
	if js.Truthy(l.own.Get("_weekdaysParseExact")) {
		if !l.own.Has("_weekdaysRegex") {
			l.computeWeekdaysParse()
		}
		if strict {
			return regexProp(l.own.Get(prefix + "StrictRegex"))
		}
		return regexProp(l.own.Get(prefix + "Regex"))
	}
	if !l.own.Has(prefix + "Regex") {
		l.own.Set(prefix+"Regex", matchWord)
	}
	if js.Truthy(l.own.Get(prefix+"StrictRegex")) && strict {
		return regexProp(l.own.Get(prefix + "StrictRegex"))
	}
	return regexProp(l.own.Get(prefix + "Regex"))
}

// WeekdaysRegex is Locale#weekdaysRegex(strict).
func (l *Locale) WeekdaysRegex(strict bool) *jsre.Regexp {
	return l.weekdaysRegexFor("_weekdays", strict)
}

// WeekdaysShortRegex is Locale#weekdaysShortRegex(strict).
func (l *Locale) WeekdaysShortRegex(strict bool) *jsre.Regexp {
	return l.weekdaysRegexFor("_weekdaysShort", strict)
}

// WeekdaysMinRegex is Locale#weekdaysMinRegex(strict).
func (l *Locale) WeekdaysMinRegex(strict bool) *jsre.Regexp {
	return l.weekdaysRegexFor("_weekdaysMin", strict)
}

// IsPM is Locale#isPM(input).
func (l *Locale) IsPM(input Value) bool {
	if f, ok := l.own.Get("isPM").(func(*Locale, string) bool); ok {
		return f(l, js.ToString(input))
	}
	return js.CharAt16(js.ToLower(js.ToString(input)), 0) == "p"
}

// Meridiem is Locale#meridiem(hours, minutes, isLower).
func (l *Locale) Meridiem(hours, minutes float64, isLower bool) string {
	if f, ok := l.own.Get("meridiem").(func(*Locale, float64, float64, bool) string); ok {
		return f(l, hours, minutes, isLower)
	}
	if hours > 11 {
		if isLower {
			return "pm"
		}
		return "PM"
	}
	if isLower {
		return "am"
	}
	return "AM"
}
