package moment

import (
	"math"
	"sort"

	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

var unitAliases = map[string]string{
	"D": "date", "dates": "date", "date": "date",
	"d": "day", "days": "day", "day": "day",
	"e": "weekday", "weekdays": "weekday", "weekday": "weekday",
	"E": "isoWeekday", "isoweekdays": "isoWeekday", "isoweekday": "isoWeekday",
	"DDD": "dayOfYear", "dayofyears": "dayOfYear", "dayofyear": "dayOfYear",
	"h": "hour", "hours": "hour", "hour": "hour",
	"ms": "millisecond", "milliseconds": "millisecond", "millisecond": "millisecond",
	"m": "minute", "minutes": "minute", "minute": "minute",
	"M": "month", "months": "month", "month": "month",
	"Q": "quarter", "quarters": "quarter", "quarter": "quarter",
	"s": "second", "seconds": "second", "second": "second",
	"gg": "weekYear", "weekyears": "weekYear", "weekyear": "weekYear",
	"GG": "isoWeekYear", "isoweekyears": "isoWeekYear", "isoweekyear": "isoWeekYear",
	"w": "week", "weeks": "week", "week": "week",
	"W": "isoWeek", "isoweeks": "isoWeek", "isoweek": "isoWeek",
	"y": "year", "years": "year", "year": "year",
}

var unitPriorities = map[string]float64{
	"date": 9, "day": 11, "weekday": 11, "isoWeekday": 11, "dayOfYear": 4, "hour": 13, "millisecond": 16,
	"minute": 14, "month": 8, "quarter": 7, "second": 15, "weekYear": 1, "isoWeekYear": 1, "week": 5,
	"isoWeek": 5, "year": 1,
}

// NormalizeUnits is moment.normalizeUnits: the canonical unit name of an alias, or undefined.
func NormalizeUnits(units Value) Value {
	s, ok := units.(string)
	if !ok {
		return js.Undef
	}
	if u, ok := unitAliases[s]; ok {
		return u
	}
	if u, ok := unitAliases[js.ToLower(s)]; ok {
		return u
	}
	return js.Undef
}

func normalizeUnit(units Value) string {
	s, _ := NormalizeUnits(units).(string)
	return s
}

// normalizeObjectUnits maps the keys of an object to unit names; later keys win.
func normalizeObjectUnits(input Value) *js.Object {
	out := js.NewObject()
	o, ok := input.(*js.Object)
	if !ok {
		return out
	}
	for _, k := range o.Keys() {
		if u := normalizeUnit(k); u != "" {
			out.Set(u, o.Get(k))
		}
	}
	return out
}

func prioritizedUnits(units *js.Object) []string {
	keys := units.Keys()
	sort.SliceStable(keys, func(i, j int) bool { return unitPriorities[keys[i]] < unitPriorities[keys[j]] })
	return keys
}

func (m *Moment) utcMode() bool { return js.Truthy(m.isUTC) }

// get reads a Date field in the moment's mode (UTC or local); NaN when invalid.
func (m *Moment) get(unit string) float64 {
	if !m.IsValid() {
		return js.NaN
	}
	utc := m.utcMode()
	switch unit {
	case "Milliseconds":
		return m.d.Milliseconds(utc)
	case "Seconds":
		return m.d.Seconds(utc)
	case "Minutes":
		return m.d.Minutes(utc)
	case "Hours":
		return m.d.Hours(utc)
	case "Date":
		return m.d.DateOfMonth(utc)
	case "Day":
		return m.d.Day(utc)
	case "Month":
		return m.d.Month(utc)
	case "FullYear":
		return m.d.FullYear(utc)
	}
	return js.NaN
}

// set writes a Date field; February 29 becomes the 28th when the new year is not a leap year.
func (m *Moment) set(unit string, value Value) {
	if !m.IsValid() || js.IsNaNValue(value) {
		return
	}
	utc := m.utcMode()
	v := js.ToNumber(value)
	switch unit {
	case "Milliseconds":
		m.d.SetMilliseconds(utc, v)
	case "Seconds":
		m.d.SetSeconds(utc, v)
	case "Minutes":
		m.d.SetMinutes(utc, v)
	case "Hours":
		m.d.SetHours(utc, v)
	case "Date":
		m.d.SetDate(utc, v)
	case "FullYear":
		month, date := m.Month(), m.Date()
		if date == 29 && month == 1 && !isLeapYear(v) {
			date = 28
		}
		m.d.SetFullYear(utc, v, month, date)
	}
}

func (m *Moment) getSet(unit string, value Value) Value {
	if !js.IsNullish(value) {
		m.set(unit, value)
		return m
	}
	return m.get(unit)
}

// Year is moment#year() as a getter; the other unit getters below follow the same pattern.
func (m *Moment) Year() float64 { return m.get("FullYear") }

// Month is moment#month() (0 to 11).
func (m *Moment) Month() float64 { return m.get("Month") }

// Date is moment#date() (day of the month).
func (m *Moment) Date() float64 { return m.get("Date") }

// Hours is moment#hours().
func (m *Moment) Hours() float64 { return m.get("Hours") }

// Minutes is moment#minutes().
func (m *Moment) Minutes() float64 { return m.get("Minutes") }

// Seconds is moment#seconds().
func (m *Moment) Seconds() float64 { return m.get("Seconds") }

// Milliseconds is moment#milliseconds().
func (m *Moment) Milliseconds() float64 { return m.get("Milliseconds") }

// Day is moment#day() (0 is Sunday).
func (m *Moment) Day() float64 {
	if !m.IsValid() {
		return js.NaN
	}
	return m.get("Day")
}

// Weekday is moment#weekday(): the day of the week counted from the locale's first day.
func (m *Moment) Weekday() float64 {
	if !m.IsValid() {
		return js.NaN
	}
	return math.Mod(m.Day()+7-js.ToNumber(m.LocaleData().weekSetting("dow")), 7)
}

// ISOWeekday is moment#isoWeekday() (1 is Monday, 7 is Sunday).
func (m *Moment) ISOWeekday() float64 {
	if !m.IsValid() {
		return js.NaN
	}
	if d := m.Day(); d != 0 && !math.IsNaN(d) {
		return d
	}
	return 7
}

// DayOfYear is moment#dayOfYear().
func (m *Moment) DayOfYear() float64 {
	return js.Round((m.Clone().StartOf("day").ValueOf()-m.Clone().StartOf("year").ValueOf())/864e5) + 1
}

// Quarter is moment#quarter().
func (m *Moment) Quarter() float64 { return math.Ceil((m.Month() + 1) / 3) }

// Week is moment#week(): the week of the year in the locale's week rules.
func (m *Moment) Week() float64 { return m.LocaleData().Week(m) }

// ISOWeek is moment#isoWeek().
func (m *Moment) ISOWeek() float64 { return weekOfYear(m, 1, 4).week }

// WeekYear is moment#weekYear().
func (m *Moment) WeekYear() float64 {
	l := m.LocaleData()
	return weekOfYear(m, js.ToNumber(l.weekSetting("dow")), js.ToNumber(l.weekSetting("doy"))).year
}

// ISOWeekYear is moment#isoWeekYear().
func (m *Moment) ISOWeekYear() float64 { return weekOfYear(m, 1, 4).year }

func (m *Moment) setMonth(value Value) {
	if !m.IsValid() {
		return
	}
	if s, ok := value.(string); ok {
		if digitsOnly.Test(s) {
			value = toInt(s)
		} else {
			value = m.LocaleData().MonthsParse(s, js.Undef, false)
			if _, ok := value.(float64); !ok {
				return
			}
		}
	}
	month := js.ToNumber(value)
	date := m.Date()
	if !(date < 29) {
		date = math.Min(date, daysInMonth(m.Year(), month))
	}
	m.d.SetMonth(m.utcMode(), month, date)
}

var digitsOnly = jsre.MustCompile(`^\d+$`, "")

// GetSetMonth is moment#month(value): a getter without a value.
func (m *Moment) GetSetMonth(value Value) Value {
	if !js.IsNullish(value) {
		m.setMonth(value)
		return m
	}
	return m.Month()
}

func parseWeekday(input Value, l *Locale) Value {
	s, ok := input.(string)
	if !ok {
		return input
	}
	if !js.IsNaNValue(s) {
		return js.ParseInt(s)
	}
	if v, ok := l.WeekdaysParse(s, js.Undef, false).(float64); ok {
		return v
	}
	return nil
}

// GetSetDay is moment#day(value).
func (m *Moment) GetSetDay(input Value) Value {
	if !m.IsValid() {
		if !js.IsNullish(input) {
			return m
		}
		return js.NaN
	}
	day := m.get("Day")
	if !js.IsNullish(input) {
		return m.Add(js.ToNumber(parseWeekday(input, m.LocaleData()))-day, "d")
	}
	return day
}

// GetSetWeekday is moment#weekday(value).
func (m *Moment) GetSetWeekday(input Value) Value {
	if !m.IsValid() {
		if !js.IsNullish(input) {
			return m
		}
		return js.NaN
	}
	weekday := m.Weekday()
	if js.IsNullish(input) {
		return weekday
	}
	return m.Add(js.ToNumber(input)-weekday, "d")
}

// GetSetISOWeekday is moment#isoWeekday(value).
func (m *Moment) GetSetISOWeekday(input Value) Value {
	if !m.IsValid() {
		if !js.IsNullish(input) {
			return m
		}
		return js.NaN
	}
	if js.IsNullish(input) {
		return m.ISOWeekday()
	}
	var weekday Value
	if s, ok := input.(string); ok {
		n := math.Mod(js.ToNumber(m.LocaleData().WeekdaysParse(s, js.Undef, false)), 7)
		if n == 0 || math.IsNaN(n) {
			n = 7
		}
		weekday = n
	} else if js.IsNaNValue(input) {
		weekday = nil
	} else {
		weekday = input
	}
	if d := math.Mod(m.Day(), 7); d != 0 && !math.IsNaN(d) {
		return m.GetSetDay(weekday)
	}
	return m.GetSetDay(js.ToNumber(weekday) - 7)
}

// GetSetDayOfYear is moment#dayOfYear(value).
func (m *Moment) GetSetDayOfYear(input Value) Value {
	doy := m.DayOfYear()
	if js.IsNullish(input) {
		return doy
	}
	return m.Add(js.ToNumber(input)-doy, "d")
}

// GetSetQuarter is moment#quarter(value).
func (m *Moment) GetSetQuarter(input Value) Value {
	if js.IsNullish(input) {
		return m.Quarter()
	}
	return m.GetSetMonth((js.ToNumber(input)-1)*3 + math.Mod(m.Month(), 3))
}

// GetSetWeek is moment#week(value).
func (m *Moment) GetSetWeek(input Value) Value {
	week := m.Week()
	if js.IsNullish(input) {
		return week
	}
	return m.Add((js.ToNumber(input)-week)*7, "d")
}

// GetSetISOWeek is moment#isoWeek(value).
func (m *Moment) GetSetISOWeek(input Value) Value {
	week := m.ISOWeek()
	if js.IsNullish(input) {
		return week
	}
	return m.Add((js.ToNumber(input)-week)*7, "d")
}

// GetSetWeekYear is moment#weekYear(value).
func (m *Moment) GetSetWeekYear(input Value) Value {
	l := m.LocaleData()
	dow, doy := js.ToNumber(l.weekSetting("dow")), js.ToNumber(l.weekSetting("doy"))
	return m.weekYearHelper(input, m.Week(), m.Weekday()+dow, dow, doy)
}

// GetSetISOWeekYear is moment#isoWeekYear(value).
func (m *Moment) GetSetISOWeekYear(input Value) Value {
	return m.weekYearHelper(input, m.ISOWeek(), m.ISOWeekday(), 1, 4)
}

func (m *Moment) weekYearHelper(input Value, week, weekday, dow, doy float64) Value {
	if js.IsNullish(input) {
		return weekOfYear(m, dow, doy).year
	}
	target := weeksInYear(input, dow, doy)
	if week > target {
		week = target
	}
	res := dayOfYearFromWeeks(input, week, weekday, dow, doy)
	date := createUTCDateFrom(res.year, 0, res.dayOfYear)
	m.getSet("FullYear", date.FullYear(true))
	m.GetSetMonth(date.Month(true))
	m.getSet("Date", date.DateOfMonth(true))
	return m
}

// WeeksInYear is moment#weeksInYear() in the locale's week rules; the other counters follow.
func (m *Moment) WeeksInYear() float64 {
	l := m.LocaleData()
	return weeksInYear(m.Year(), js.ToNumber(l.weekSetting("dow")), js.ToNumber(l.weekSetting("doy")))
}

// WeeksInWeekYear is moment#weeksInWeekYear().
func (m *Moment) WeeksInWeekYear() float64 {
	l := m.LocaleData()
	return weeksInYear(m.WeekYear(), js.ToNumber(l.weekSetting("dow")), js.ToNumber(l.weekSetting("doy")))
}

// ISOWeeksInYear is moment#isoWeeksInYear().
func (m *Moment) ISOWeeksInYear() float64 { return weeksInYear(m.Year(), 1, 4) }

// ISOWeeksInISOWeekYear is moment#isoWeeksInISOWeekYear().
func (m *Moment) ISOWeeksInISOWeekYear() float64 { return weeksInYear(m.ISOWeekYear(), 1, 4) }

// IsLeapYear is moment#isLeapYear().
func (m *Moment) IsLeapYear() bool { return isLeapYear(m.Year()) }

// DaysInMonth is moment#daysInMonth().
func (m *Moment) DaysInMonth() float64 { return daysInMonth(m.Year(), m.Month()) }

var chunkOffset = jsre.MustCompile(`([\+\-]|\d\d)`, "gi")

// offsetFromString reads the last offset matcher finds in s as minutes; nil when there is none or it is out of
// range.
func offsetFromString(matcher *jsre.Regexp, s string) Value {
	matches := matcher.Match(s)
	if matches == nil {
		return nil
	}
	chunk := matches.Groups[len(matches.Groups)-1]
	parts := []Value{"-", 0.0, 0.0}
	if p := chunkOffset.Match(chunk); p != nil {
		parts = make([]Value, 3)
		for i := range parts {
			if i < len(p.Groups) {
				parts[i] = p.Groups[i]
			} else {
				parts[i] = js.Undef
			}
		}
	}
	minutes := js.ToNumber(parts[1])*60 + toInt(parts[2])
	plus := js.StrictEquals(parts[0], "+")
	if toInt(parts[2]) > 59 || (plus && minutes > 14*60) || (!plus && minutes > 12*60) {
		return nil
	}
	if minutes == 0 {
		return 0.0
	}
	if plus {
		return minutes
	}
	return -minutes
}

// cloneWithOffset returns input as a moment in the same mode (UTC offset or local) as model.
func cloneWithOffset(input Value, model *Moment) *Moment {
	if model.utcMode() {
		res := model.Clone()
		var other float64
		switch x := input.(type) {
		case *Moment:
			other = x.ValueOf()
		case *js.Date:
			other = x.Time()
		default:
			other = localMoment(input).ValueOf()
		}
		diff := other - res.ValueOf()
		res.d.SetTime(res.d.Time() + diff)
		return res
	}
	return localMoment(input).Local(js.Undef)
}

func (m *Moment) dateOffset() float64 {
	return -js.Round(m.d.TimezoneOffset())
}

// UTCOffset is moment#utcOffset() as a getter: minutes east of UTC, NaN when invalid.
func (m *Moment) UTCOffset() Value {
	return m.GetSetOffset(js.Undef, js.Undef, js.Undef)
}

// GetSetOffset is moment#utcOffset(input, keepLocalTime, keepMinutes); numbers under 16 in absolute value are
// hours unless keepMinutes.
func (m *Moment) GetSetOffset(input, keepLocalTime, keepMinutes Value) Value {
	offset := 0.0
	if js.Truthy(m.offset) {
		offset = js.ToNumber(m.offset)
	}
	if !m.IsValid() {
		if !js.IsNullish(input) {
			return m
		}
		return js.NaN
	}
	if js.IsNullish(input) {
		if m.utcMode() {
			return offset
		}
		return m.dateOffset()
	}
	var value float64
	if s, ok := input.(string); ok {
		parsed := offsetFromString(matchShortOffset, s)
		if parsed == nil {
			return m
		}
		value = parsed.(float64)
	} else {
		value = js.ToNumber(input)
		if math.Abs(value) < 16 && !js.Truthy(keepMinutes) {
			value *= 60
		}
	}
	var localAdjust Value = js.Undef
	if !m.utcMode() && js.Truthy(keepLocalTime) {
		localAdjust = m.dateOffset()
	}
	m.offset = value
	m.isUTC = true
	if !js.IsUndefined(localAdjust) {
		m.Add(localAdjust, "m")
	}
	if offset != value && !js.Truthy(keepLocalTime) {
		addSubtract(m, CreateDuration(value-offset, "m"), 1)
	}
	return m
}

// GetSetZone is the deprecated moment#zone(input): minutes west of UTC.
func (m *Moment) GetSetZone(input, keepLocalTime Value) Value {
	zoneDeprecation.fire()
	if !js.IsNullish(input) {
		if _, ok := input.(string); !ok {
			input = -js.ToNumber(input)
		}
		m.GetSetOffset(input, keepLocalTime, js.Undef)
		return m
	}
	return -js.ToNumber(m.UTCOffset())
}

// UTC is moment#utc(keepLocalTime).
func (m *Moment) UTC(keepLocalTime Value) *Moment {
	m.GetSetOffset(0.0, keepLocalTime, js.Undef)
	return m
}

// Local is moment#local(keepLocalTime).
func (m *Moment) Local(keepLocalTime Value) *Moment {
	if m.utcMode() {
		m.GetSetOffset(0.0, keepLocalTime, js.Undef)
		m.isUTC = false
		if js.Truthy(keepLocalTime) {
			m.Subtract(m.dateOffset(), "m")
		}
	}
	return m
}

// ParseZone is moment#parseZone(): the offset of the parsed input.
func (m *Moment) ParseZone() *Moment {
	if !js.IsNullish(m.tzm) {
		m.GetSetOffset(m.tzm, false, true)
	} else if s, ok := m.i.(string); ok {
		if tz := offsetFromString(matchOffset, s); tz != nil {
			m.GetSetOffset(tz, js.Undef, js.Undef)
		} else {
			m.GetSetOffset(0.0, true, js.Undef)
		}
	}
	return m
}

// HasAlignedHourOffset is moment#hasAlignedHourOffset(input).
func (m *Moment) HasAlignedHourOffset(input Value) bool {
	if !m.IsValid() {
		return false
	}
	other := 0.0
	if js.Truthy(input) {
		other = js.ToNumber(localMoment(input).UTCOffset())
	}
	return math.Mod(js.ToNumber(m.UTCOffset())-other, 60) == 0
}

// IsDST is moment#isDST().
func (m *Moment) IsDST() bool {
	offset := js.ToNumber(m.UTCOffset())
	return offset > js.ToNumber(m.Clone().GetSetMonth(0.0).(*Moment).UTCOffset()) ||
		offset > js.ToNumber(m.Clone().GetSetMonth(5.0).(*Moment).UTCOffset())
}

// IsDSTShifted is the deprecated moment#isDSTShifted().
func (m *Moment) IsDSTShifted() bool {
	dstShiftedDeprecation.fire()
	if b, ok := m.isDSTShifted.(bool); ok {
		return b
	}
	c := copyConfig(newConfig(), m)
	c = prepareConfig(c)
	result := false
	if c.a != nil {
		var other *Moment
		if js.Truthy(c.isUTC) {
			other = utcMoment(append([]any(nil), c.a...))
		} else {
			other = localMoment(append([]any(nil), c.a...))
		}
		result = m.IsValid() && compareArrays(c.a, other.ToArray()) > 0
	}
	m.isDSTShifted = result
	return result
}

func compareArrays(a, b []Value) int {
	n := min(len(a), len(b))
	diffs := 0
	for i := 0; i < n; i++ {
		if toInt(a[i]) != toInt(b[i]) {
			diffs++
		}
	}
	if len(a) > len(b) {
		return diffs + len(a) - len(b)
	}
	return diffs + len(b) - len(a)
}

// IsLocal is moment#isLocal().
func (m *Moment) IsLocal() bool { return m.IsValid() && !m.utcMode() }

// IsUtcOffset is moment#isUtcOffset().
func (m *Moment) IsUtcOffset() bool { return m.IsValid() && m.utcMode() }

// IsUTC is moment#isUtc().
func (m *Moment) IsUTC() bool {
	return m.IsValid() && m.utcMode() && js.StrictEquals(m.offset, 0.0)
}

// ZoneAbbr is moment#zoneAbbr().
func (m *Moment) ZoneAbbr() string {
	if m.utcMode() {
		return "UTC"
	}
	return ""
}

// ZoneName is moment#zoneName().
func (m *Moment) ZoneName() string {
	if m.utcMode() {
		return "Coordinated Universal Time"
	}
	return ""
}

func (m *Moment) eraValue(pick func(era *js.Object) Value) Value {
	val := m.Clone().StartOf("day").ValueOf()
	for _, e := range m.LocaleData().Eras() {
		era, ok := e.(*js.Object)
		if !ok {
			continue
		}
		since, until := js.ToNumber(era.Get("since")), js.ToNumber(era.Get("until"))
		if (since <= val && val <= until) || (until <= val && val <= since) {
			return pick(era)
		}
	}
	return nil
}

// EraName is moment#eraName(); the era methods return "" outside every era.
func (m *Moment) EraName() string {
	if v := m.eraValue(func(e *js.Object) Value { return e.Get("name") }); v != nil {
		return js.ToString(v)
	}
	return ""
}

// EraNarrow is moment#eraNarrow().
func (m *Moment) EraNarrow() string {
	if v := m.eraValue(func(e *js.Object) Value { return e.Get("narrow") }); v != nil {
		return js.ToString(v)
	}
	return ""
}

// EraAbbr is moment#eraAbbr().
func (m *Moment) EraAbbr() string {
	if v := m.eraValue(func(e *js.Object) Value { return e.Get("abbr") }); v != nil {
		return js.ToString(v)
	}
	return ""
}

// EraYear is moment#eraYear(): the year counted in the moment's era, or the year outside every era.
func (m *Moment) EraYear() float64 {
	v := m.eraValue(func(era *js.Object) Value {
		since, until := js.ToNumber(era.Get("since")), js.ToNumber(era.Get("until"))
		dir := -1.0
		if since <= until {
			dir = 1
		}
		years := m.Year() - localMoment(era.Get("since")).Year()
		return float64(years*dir) + js.ToNumber(era.Get("offset"))
	})
	if v != nil {
		return v.(float64)
	}
	return m.Year()
}
