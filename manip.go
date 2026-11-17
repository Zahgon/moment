package moment

import (
	"math"

	"github.com/moment/moment/internal/js"
)

// Clone is moment#clone.
func (m *Moment) Clone() *Moment { return newMoment(m) }

// ValueOf is moment#valueOf: milliseconds since the epoch, NaN when invalid.
func (m *Moment) ValueOf() float64 {
	offset := 0.0
	if js.Truthy(m.offset) {
		offset = js.ToNumber(m.offset)
	}
	return m.d.Time() - offset*60000
}

// Unix is moment#unix: whole seconds since the epoch.
func (m *Moment) Unix() float64 { return math.Floor(m.ValueOf() / 1000) }

// ToDate is moment#toDate.
func (m *Moment) ToDate() *js.Date { return js.NewDate(m.ValueOf()) }

// ToArray is moment#toArray.
func (m *Moment) ToArray() []Value {
	return []Value{m.Year(), m.Month(), m.Date(), m.Hours(), m.Minutes(), m.Seconds(), m.Milliseconds()}
}

// ToObject is moment#toObject.
func (m *Moment) ToObject() *js.Object {
	return js.Obj("years", m.Year(), "months", m.Month(), "date", m.Date(), "hours", m.Hours(),
		"minutes", m.Minutes(), "seconds", m.Seconds(), "milliseconds", m.Milliseconds())
}

// CreationData is moment#creationData.
func (m *Moment) CreationData() *js.Object {
	var locale Value = js.Undef
	if m.locale != nil {
		locale = m.locale
	}
	return js.Obj("input", m.i, "format", m.f, "locale", locale, "isUTC", m.isUTC, "strict", m.strict)
}

// ParsingFlags is moment#parsingFlags: a copy of the flags the parse recorded.
func (m *Moment) ParsingFlags() *js.Object { return getParsingFlags(m).Clone() }

// InvalidAt is moment#invalidAt: the overflowing unit index, -1 when none.
func (m *Moment) InvalidAt() Value { return getParsingFlags(m).Get("overflow") }

// Locale is moment#locale() as a getter: the locale's name.
func (m *Moment) Locale() Value {
	if m.locale == nil {
		return js.Undef
	}
	return m.locale.Abbr()
}

// SetLocale is moment#locale(key): the locale is changed when key names a known locale.
func (m *Moment) SetLocale(key Value) *Moment {
	if l := getLocale(key); l != nil {
		m.locale = l
	}
	return m
}

// LocaleData is moment#localeData.
func (m *Moment) LocaleData() *Locale { return m.locale }

// Add is moment#add(value, period); a period given first is accepted with a deprecation.
func (m *Moment) Add(val, period Value) *Moment { return m.adder(1, "add", val, period) }

// Subtract is moment#subtract(value, period).
func (m *Moment) Subtract(val, period Value) *Moment { return m.adder(-1, "subtract", val, period) }

func (m *Moment) adder(direction float64, name string, val, period Value) *Moment {
	if period != nil && !js.IsNaNValue(period) {
		deprecateSimple(name, "moment()."+name+"(period, number) is deprecated. Please use moment()."+name+
			"(number, period). See http://momentjs.com/guides/#/warnings/add-inverted-param/ for more info.")
		val, period = period, val
	}
	addSubtract(m, CreateDuration(val, period), direction)
	return m
}

func addSubtract(m *Moment, d *Duration, direction float64) {
	milliseconds := d.milliseconds
	days := absRound(d.days)
	months := absRound(d.months)
	if !m.IsValid() {
		return
	}
	if months != 0 && !math.IsNaN(months) {
		m.setMonth(m.get("Month") + float64(months*direction))
	}
	if days != 0 && !math.IsNaN(days) {
		m.set("Date", m.get("Date")+float64(days*direction))
	}
	if milliseconds != 0 && !math.IsNaN(milliseconds) {
		m.d.SetTime(m.d.Time() + float64(milliseconds*direction))
	}
}

const (
	msPerSecond   = 1000
	msPerMinute   = 60 * msPerSecond
	msPerHour     = 60 * msPerMinute
	msPer400Years = (365*400 + 97) * 24 * msPerHour
)

func localStartOfDate(y, m, d float64) float64 {
	if y < 100 && y >= 0 {
		return js.NewLocalDate(y+400, m, d).Time() - msPer400Years
	}
	return js.NewLocalDate(y, m, d).Time()
}

func utcStartOfDate(y, m, d float64) float64 {
	if y < 100 && y >= 0 {
		return js.DateUTC(y+400, m, d) - msPer400Years
	}
	return js.DateUTC(y, m, d)
}

// boundary computes the start of units, or the end when end is set; ok is false for units it ignores. A unit that
// normalizes to a name with no boundary of its own (dayOfYear, weekday, isoWeekday, weekYear, isoWeekYear) leaves the
// time unset, which makes the moment invalid, as the original's setTime(undefined) does.
func (m *Moment) boundary(units Value, end bool) (float64, bool) {
	unit := normalizeUnit(units)
	if unit == "" || unit == "millisecond" || !m.IsValid() {
		return 0, false
	}
	startOfDate := localStartOfDate
	if m.utcMode() {
		startOfDate = utcStartOfDate
	}
	extra, last := 0.0, 0.0
	if end {
		extra, last = 1, 1
	}
	t := m.d.Time()
	localShift := 0.0
	if !m.utcMode() {
		localShift = js.ToNumber(m.UTCOffset()) * msPerMinute
	}
	switch unit {
	case "year":
		return startOfDate(m.Year()+extra, 0, 1) - last, true
	case "quarter":
		return startOfDate(m.Year(), m.Month()-math.Mod(m.Month(), 3)+3*extra, 1) - last, true
	case "month":
		return startOfDate(m.Year(), m.Month()+extra, 1) - last, true
	case "week":
		return startOfDate(m.Year(), m.Month(), m.Date()-m.Weekday()+7*extra) - last, true
	case "isoWeek":
		return startOfDate(m.Year(), m.Month(), m.Date()-(m.ISOWeekday()-1)+7*extra) - last, true
	case "day", "date":
		return startOfDate(m.Year(), m.Month(), m.Date()+extra) - last, true
	case "hour":
		if end {
			return t + msPerHour - mod(t+localShift, msPerHour) - 1, true
		}
		return t - mod(t+localShift, msPerHour), true
	case "minute":
		if end {
			return t + msPerMinute - mod(t, msPerMinute) - 1, true
		}
		return t - mod(t, msPerMinute), true
	case "second":
		if end {
			return t + msPerSecond - mod(t, msPerSecond) - 1, true
		}
		return t - mod(t, msPerSecond), true
	}
	return math.NaN(), true
}

// StartOf is moment#startOf(units).
func (m *Moment) StartOf(units Value) *Moment {
	if t, ok := m.boundary(units, false); ok {
		m.d.SetTime(t)
	}
	return m
}

// EndOf is moment#endOf(units).
func (m *Moment) EndOf(units Value) *Moment {
	if t, ok := m.boundary(units, true); ok {
		m.d.SetTime(t)
	}
	return m
}

func localInput(input Value) *Moment {
	if mm, ok := input.(*Moment); ok {
		return mm
	}
	return localMoment(input)
}

func compareUnit(units Value) string {
	if u := normalizeUnit(units); u != "" {
		return u
	}
	return "millisecond"
}

// IsAfter is moment#isAfter(input, units).
func (m *Moment) IsAfter(input, units Value) bool {
	other := localInput(input)
	if !(m.IsValid() && other.IsValid()) {
		return false
	}
	unit := compareUnit(units)
	if unit == "millisecond" {
		return m.ValueOf() > other.ValueOf()
	}
	return other.ValueOf() < m.Clone().StartOf(unit).ValueOf()
}

// IsBefore is moment#isBefore(input, units).
func (m *Moment) IsBefore(input, units Value) bool {
	other := localInput(input)
	if !(m.IsValid() && other.IsValid()) {
		return false
	}
	unit := compareUnit(units)
	if unit == "millisecond" {
		return m.ValueOf() < other.ValueOf()
	}
	return m.Clone().EndOf(unit).ValueOf() < other.ValueOf()
}

// IsBetween is moment#isBetween(from, to, units, inclusivity).
func (m *Moment) IsBetween(from, to, units, inclusivity Value) bool {
	lf, lt := localInput(from), localInput(to)
	if !(m.IsValid() && lf.IsValid() && lt.IsValid()) {
		return false
	}
	inc := "()"
	if js.Truthy(inclusivity) {
		inc = js.ToString(inclusivity)
	}
	first, second := js.CharAt16(inc, 0), js.CharAt16(inc, 1)
	var left, right bool
	if first == "(" {
		left = m.IsAfter(lf, units)
	} else {
		left = !m.IsBefore(lf, units)
	}
	if !left {
		return false
	}
	if second == ")" {
		right = m.IsBefore(lt, units)
	} else {
		right = !m.IsAfter(lt, units)
	}
	return right
}

// IsSame is moment#isSame(input, units).
func (m *Moment) IsSame(input, units Value) bool {
	other := localInput(input)
	if !(m.IsValid() && other.IsValid()) {
		return false
	}
	unit := compareUnit(units)
	if unit == "millisecond" {
		return m.ValueOf() == other.ValueOf()
	}
	ms := other.ValueOf()
	return m.Clone().StartOf(unit).ValueOf() <= ms && ms <= m.Clone().EndOf(unit).ValueOf()
}

// IsSameOrAfter is moment#isSameOrAfter(input, units).
func (m *Moment) IsSameOrAfter(input, units Value) bool {
	return m.IsSame(input, units) || m.IsAfter(input, units)
}

// IsSameOrBefore is moment#isSameOrBefore(input, units).
func (m *Moment) IsSameOrBefore(input, units Value) bool {
	return m.IsSame(input, units) || m.IsBefore(input, units)
}

// Diff is moment#diff(input, units, asFloat).
func (m *Moment) Diff(input, units, asFloat Value) float64 {
	if !m.IsValid() {
		return js.NaN
	}
	that := cloneWithOffset(input, m)
	if !that.IsValid() {
		return js.NaN
	}
	zoneDelta := (js.ToNumber(that.UTCOffset()) - js.ToNumber(m.UTCOffset())) * 6e4
	delta := m.ValueOf() - that.ValueOf()
	var output float64
	switch normalizeUnit(units) {
	case "year":
		output = monthDiff(m, that) / 12
	case "month":
		output = monthDiff(m, that)
	case "quarter":
		output = monthDiff(m, that) / 3
	case "second":
		output = delta / 1e3
	case "minute":
		output = delta / 6e4
	case "hour":
		output = delta / 36e5
	case "day":
		output = (delta - zoneDelta) / 864e5
	case "week":
		output = (delta - zoneDelta) / 6048e5
	default:
		output = delta
	}
	if js.Truthy(asFloat) {
		return output
	}
	return absFloor(output)
}

func monthDiff(a, b *Moment) float64 {
	if a.Date() < b.Date() {
		return -monthDiff(b, a)
	}
	whole := float64((b.Year()-a.Year())*12) + (b.Month() - a.Month())
	anchor := a.Clone().Add(whole, "months")
	var adjust float64
	if b.ValueOf()-anchor.ValueOf() < 0 {
		anchor2 := a.Clone().Add(whole-1, "months")
		adjust = (b.ValueOf() - anchor.ValueOf()) / (anchor.ValueOf() - anchor2.ValueOf())
	} else {
		anchor2 := a.Clone().Add(whole+1, "months")
		adjust = (b.ValueOf() - anchor.ValueOf()) / (anchor2.ValueOf() - anchor.ValueOf())
	}
	r := -(whole + adjust)
	if r == 0 || math.IsNaN(r) {
		return 0
	}
	return r
}

// CalendarFormat is moment.calendarFormat(moment, now): the calendar key for the distance in days.
func CalendarFormat(m *Moment, now Value) string {
	diff := m.Diff(now, "days", true)
	switch {
	case diff < -6:
		return "sameElse"
	case diff < -1:
		return "lastWeek"
	case diff < 0:
		return "lastDay"
	case diff < 1:
		return "sameDay"
	case diff < 2:
		return "nextDay"
	case diff < 7:
		return "nextWeek"
	}
	return "sameElse"
}

var momentInputProperties = []string{"years", "year", "y", "months", "month", "M", "days", "day", "d", "dates",
	"date", "D", "hours", "hour", "h", "minutes", "minute", "m", "seconds", "second", "s", "milliseconds",
	"millisecond", "ms"}

var calendarSpecProperties = []string{"sameDay", "nextDay", "lastDay", "nextWeek", "lastWeek", "sameElse"}

func hasAnyKey(v Value, keys []string) bool {
	o, ok := v.(*js.Object)
	if !ok || o.Len() == 0 {
		return false
	}
	for _, k := range keys {
		if o.Has(k) {
			return true
		}
	}
	return false
}

func isMomentInput(v Value) bool {
	switch x := v.(type) {
	case *Moment, *js.Date, string, float64, nil, js.Undefined:
		return true
	case []any:
		return len(x) >= 0
	case []string:
		return true
	}
	return hasAnyKey(v, momentInputProperties)
}

// Calendar is moment#calendar(time, formats): this moment formatted with the calendar format for its distance
// from time (now by default).
func (m *Moment) Calendar(args ...Value) string {
	time, formats := arg(args, 0), arg(args, 1)
	if len(args) == 1 {
		switch {
		case !js.Truthy(args[0]):
			time, formats = js.Undef, js.Undef
		case isMomentInput(args[0]):
			time, formats = args[0], js.Undef
		case hasAnyKey(args[0], calendarSpecProperties):
			time, formats = js.Undef, args[0]
		}
	}
	var now Value = time
	if !js.Truthy(time) {
		now = localMoment()
	}
	sod := cloneWithOffset(now, m).StartOf("day")
	format := CalendarFormat(m, sod)
	var output Value
	if f, ok := formats.(*js.Object); ok {
		output = f.Get(format)
	}
	if js.Truthy(output) {
		return m.Format(output)
	}
	return m.Format(m.LocaleData().Calendar(format, m, localMoment(now)))
}

// From is moment#from(time, withoutSuffix).
func (m *Moment) From(time, withoutSuffix Value) string {
	return m.relative(js.Obj("to", m, "from", time), time, withoutSuffix)
}

// To is moment#to(time, withoutSuffix).
func (m *Moment) To(time, withoutSuffix Value) string {
	return m.relative(js.Obj("from", m, "to", time), time, withoutSuffix)
}

func (m *Moment) relative(span *js.Object, time, withoutSuffix Value) string {
	other, isM := time.(*Moment)
	if m.IsValid() && ((isM && other.IsValid()) || localMoment(time).IsValid()) {
		d := CreateDuration(span, js.Undef)
		d.SetLocale(m.Locale())
		return d.Humanize(!js.Truthy(withoutSuffix), js.Undef)
	}
	return m.LocaleData().InvalidDate()
}

// FromNow is moment#fromNow(withoutSuffix).
func (m *Moment) FromNow(withoutSuffix Value) string { return m.From(localMoment(), withoutSuffix) }

// ToNow is moment#toNow(withoutSuffix).
func (m *Moment) ToNow(withoutSuffix Value) string { return m.To(localMoment(), withoutSuffix) }

// Get is moment#get(units).
func (m *Moment) Get(units Value) (Value, error) {
	unit := normalizeUnit(units)
	if unit == "" {
		return m, nil
	}
	return m.unitAccessor(unit, js.Undef), nil
}

// Set is moment#set(units, value); an object sets several units, highest-order units first.
func (m *Moment) Set(units, value Value) Value {
	if o, ok := units.(*js.Object); ok {
		normalized := normalizeObjectUnits(o)
		for _, u := range prioritizedUnits(normalized) {
			m.unitAccessor(u, normalized.Get(u))
		}
		return m
	}
	if units == nil || isArray(units) || isObject(units) {
		return m
	}
	if unit := normalizeUnit(units); unit != "" {
		return m.unitAccessor(unit, value)
	}
	return m
}

// unitAccessor calls the getter-setter method of a normalized unit name.
func (m *Moment) unitAccessor(unit string, value Value) Value {
	switch unit {
	case "year":
		return m.getSet("FullYear", value)
	case "month":
		return m.GetSetMonth(value)
	case "date":
		return m.getSet("Date", value)
	case "day":
		return m.GetSetDay(value)
	case "weekday":
		return m.GetSetWeekday(value)
	case "isoWeekday":
		return m.GetSetISOWeekday(value)
	case "dayOfYear":
		return m.GetSetDayOfYear(value)
	case "hour":
		return m.getSet("Hours", value)
	case "minute":
		return m.getSet("Minutes", value)
	case "second":
		return m.getSet("Seconds", value)
	case "millisecond":
		return m.getSet("Milliseconds", value)
	case "quarter":
		return m.GetSetQuarter(value)
	case "week":
		return m.GetSetWeek(value)
	case "isoWeek":
		return m.GetSetISOWeek(value)
	case "weekYear":
		return m.GetSetWeekYear(value)
	case "isoWeekYear":
		return m.GetSetISOWeekYear(value)
	}
	return m
}

func pickBy(after bool, moments []Value) *Moment {
	if len(moments) == 1 && isArray(moments[0]) {
		moments = asArray(moments[0])
	}
	if len(moments) == 0 {
		return localMoment()
	}
	var res *Moment
	i := 0
	for ; i < len(moments); i++ {
		if mm, ok := moments[i].(*Moment); ok {
			res = mm
			break
		}
	}
	if res == nil {
		return CreateInvalid(js.Undef)
	}
	for i++; i < len(moments); i++ {
		mm, ok := moments[i].(*Moment)
		if !ok {
			continue
		}
		if !mm.IsValid() || (after && mm.IsAfter(res, js.Undef)) || (!after && mm.IsBefore(res, js.Undef)) {
			res = mm
		}
	}
	return res
}

// Min is moment.min(moments...): the earliest moment (an array argument is unwrapped).
func Min(args ...Value) *Moment { return pickBy(false, args) }

// Max is moment.max(moments...): the latest moment.
func Max(args ...Value) *Moment { return pickBy(true, args) }

// PrototypeMin is the deprecated moment#min(input...): the later of this moment and the input.
func (m *Moment) PrototypeMin(args ...Value) *Moment {
	minDeprecation.fire()
	other := localMoment(args...)
	if m.IsValid() && other.IsValid() {
		if other.ValueOf() < m.ValueOf() {
			return m
		}
		return other
	}
	return CreateInvalid(js.Undef)
}

// PrototypeMax is the deprecated moment#max(input...): the earlier of this moment and the input.
func (m *Moment) PrototypeMax(args ...Value) *Moment {
	maxDeprecation.fire()
	other := localMoment(args...)
	if m.IsValid() && other.IsValid() {
		if other.ValueOf() > m.ValueOf() {
			return m
		}
		return other
	}
	return CreateInvalid(js.Undef)
}
