package js

import (
	"fmt"
	"math"
	"time"
)

const (
	msPerSecond = 1000
	msPerMinute = 60000
	msPerHour   = 3600000
	msPerDay    = 86400000
	maxTime     = 8.64e15
	// maxLocalTime bounds the local time values that are converted to UTC, as V8 does (ten days of slack).
	maxLocalTime = maxTime + 864000000
	// maxOffsetSeconds is the window after a zone transition in which a local time is resolved against it.
	maxOffsetSeconds = 86400
)

// Date is a JavaScript Date: a time value in milliseconds since the epoch, NaN when invalid. Local time follows
// time.Local (the TZ environment variable), with the rules Node.js applies: a local time skipped or repeated by a
// transition resolves with the offset in force before the transition.
type Date struct {
	t float64
}

// NewDate is new Date(number).
func NewDate(t float64) *Date {
	return &Date{t: TimeClip(t)}
}

// Time is Date.prototype.getTime and valueOf.
func (d *Date) Time() float64 { return d.t }

// SetTime is Date.prototype.setTime.
func (d *Date) SetTime(t float64) float64 {
	d.t = TimeClip(t)
	return d.t
}

// Valid reports whether the time value is not NaN.
func (d *Date) Valid() bool { return !math.IsNaN(d.t) }

// TimeClip limits a time value to the Date range and truncates it to an integer.
func TimeClip(t float64) float64 {
	if math.IsNaN(t) || math.Abs(t) > maxTime {
		return NaN
	}
	return math.Trunc(t) + 0
}

// MakeTime combines hours, minutes, seconds and milliseconds into a time within a day.
func MakeTime(h, m, s, ms float64) float64 {
	for _, v := range []float64{h, m, s, ms} {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return NaN
		}
	}
	return math.Trunc(h)*msPerHour + math.Trunc(m)*msPerMinute + math.Trunc(s)*msPerSecond + math.Trunc(ms)
}

// MakeDay returns the day number of the given year, month (any integer, carried into the year) and date.
func MakeDay(year, month, date float64) float64 {
	if !(year >= -1000000 && year <= 1000000) || !(month >= -10000000 && month <= 10000000) ||
		math.IsNaN(date) || math.IsInf(date, 0) {
		return NaN
	}
	y := int64(year)
	m := int64(month)
	y += m / 12
	m %= 12
	if m < 0 {
		m += 12
		y--
	}
	return float64(daysFromCivil(y, m+1, 1)) + math.Trunc(date) - 1
}

// MakeDate combines a day number and a time within the day.
func MakeDate(day, t float64) float64 {
	if math.IsNaN(day) || math.IsInf(day, 0) || math.IsNaN(t) || math.IsInf(t, 0) {
		return NaN
	}
	return day*msPerDay + t
}

// daysFromCivil returns the days since 1970-01-01 of a proleptic Gregorian date.
func daysFromCivil(y, m, d int64) int64 {
	if m <= 2 {
		y--
	}
	era := floorDiv(y, 400)
	yoe := y - era*400
	mp := (m + 9) % 12
	doy := (153*mp+2)/5 + d - 1
	doe := yoe*365 + yoe/4 - yoe/100 + doy
	return era*146097 + doe - 719468
}

// civilFromDays is the inverse of daysFromCivil; month is 1-based.
func civilFromDays(z int64) (y, m, d int64) {
	z += 719468
	era := floorDiv(z, 146097)
	doe := z - era*146097
	yoe := (doe - doe/1460 + doe/36524 - doe/146096) / 365
	y = yoe + era*400
	doy := doe - (365*yoe + yoe/4 - yoe/100)
	mp := (5*doy + 2) / 153
	d = doy - (153*mp+2)/5 + 1
	m = mp + 3
	if m > 12 {
		m -= 12
	}
	if m <= 2 {
		y++
	}
	return y, m, d
}

func floorDiv(a, b int64) int64 {
	q := a / b
	if (a%b != 0) && ((a < 0) != (b < 0)) {
		q--
	}
	return q
}

// fields are the calendar fields of a time value.
type fields struct {
	year, month, date, day, hours, minutes, seconds, ms float64
}

func split(t float64) fields {
	ms := int64(t)
	days := floorDiv(ms, msPerDay)
	within := ms - days*msPerDay
	y, m, d := civilFromDays(days)
	return fields{
		year: float64(y), month: float64(m - 1), date: float64(d), day: float64(((days+4)%7 + 7) % 7),
		hours: float64(within / msPerHour), minutes: float64(within / msPerMinute % 60),
		seconds: float64(within / msPerSecond % 60), ms: float64(within % msPerSecond),
	}
}

// offsetAt returns the zone offset in milliseconds of time.Local at the UTC time value t.
func offsetAt(t float64) float64 {
	_, off := time.UnixMilli(int64(t)).In(time.Local).Zone()
	return float64(off) * msPerSecond
}

// LocalTime converts a UTC time value to local time.
func LocalTime(t float64) float64 {
	return t + offsetAt(t)
}

type transition struct {
	at            int64 // seconds since the epoch
	before, after int64 // offsets in seconds
}

// transitionsNear lists the zone transitions of time.Local within two days of the instant sec.
func transitionsNear(sec int64) (list []transition, first int64) {
	t := time.Unix(sec-2*86400, 0).In(time.Local)
	_, off := t.Zone()
	first = int64(off)
	for range 16 {
		_, end := t.ZoneBounds()
		if end.IsZero() || end.Unix() > sec+2*86400 {
			break
		}
		_, after := end.In(time.Local).Zone()
		_, before := end.Add(-time.Second).In(time.Local).Zone()
		list = append(list, transition{at: end.Unix(), before: int64(before), after: int64(after)})
		t = end.In(time.Local)
	}
	return list, first
}

// UTCFromLocal converts a local time value to UTC. A local time inside a gap or an overlap uses the offset in
// force before the transition.
func UTCFromLocal(local float64) float64 {
	if math.IsNaN(local) || !(local >= -maxLocalTime && local <= maxLocalTime) {
		return NaN
	}
	l := int64(local)
	sec := floorDiv(l, msPerSecond)
	list, first := transitionsNear(sec)
	offset := first
	if len(list) > 0 {
		offset = list[0].before
	}
	for i := len(list) - 1; i >= 0; i-- {
		tr := list[i]
		threshold := tr.at
		if sec >= tr.at-maxOffsetSeconds {
			if tr.after-tr.before >= 0 {
				threshold += tr.after
			} else {
				threshold += tr.before
			}
		}
		if sec >= threshold {
			offset = tr.after
			break
		}
	}
	return float64(l - offset*msPerSecond)
}

func (d *Date) local() (fields, bool) {
	if math.IsNaN(d.t) {
		return fields{}, false
	}
	return split(LocalTime(d.t)), true
}

func (d *Date) utc() (fields, bool) {
	if math.IsNaN(d.t) {
		return fields{}, false
	}
	return split(d.t), true
}

func (d *Date) field(utc bool, pick func(fields) float64) float64 {
	f, ok := d.pick(utc)
	if !ok {
		return NaN
	}
	return pick(f)
}

// FullYear is getFullYear (UTC selects getUTCFullYear); the other getters follow the same pattern.
func (d *Date) FullYear(utc bool) float64 {
	return d.field(utc, func(f fields) float64 { return f.year })
}

// Month is getMonth or getUTCMonth.
func (d *Date) Month(utc bool) float64 {
	return d.field(utc, func(f fields) float64 { return f.month })
}

// DateOfMonth is getDate or getUTCDate.
func (d *Date) DateOfMonth(utc bool) float64 {
	return d.field(utc, func(f fields) float64 { return f.date })
}

// Day is getDay or getUTCDay.
func (d *Date) Day(utc bool) float64 { return d.field(utc, func(f fields) float64 { return f.day }) }

// Hours is getHours or getUTCHours.
func (d *Date) Hours(utc bool) float64 {
	return d.field(utc, func(f fields) float64 { return f.hours })
}

// Minutes is getMinutes or getUTCMinutes.
func (d *Date) Minutes(utc bool) float64 {
	return d.field(utc, func(f fields) float64 { return f.minutes })
}

// Seconds is getSeconds or getUTCSeconds.
func (d *Date) Seconds(utc bool) float64 {
	return d.field(utc, func(f fields) float64 { return f.seconds })
}

// Milliseconds is getMilliseconds or getUTCMilliseconds.
func (d *Date) Milliseconds(utc bool) float64 {
	return d.field(utc, func(f fields) float64 { return f.ms })
}

func (d *Date) pick(utc bool) (fields, bool) {
	if utc {
		return d.utc()
	}
	return d.local()
}

// TimezoneOffset is getTimezoneOffset: whole minutes, truncated towards zero.
func (d *Date) TimezoneOffset() float64 {
	if math.IsNaN(d.t) {
		return NaN
	}
	t := int64(d.t)
	l := int64(LocalTime(d.t))
	return float64((t - l) / msPerMinute)
}

// store finishes a setter: local results are converted to UTC first.
func (d *Date) store(v float64, utc bool) float64 {
	if utc {
		d.t = TimeClip(v)
	} else {
		d.t = TimeClip(UTCFromLocal(v))
	}
	return d.t
}

func optional(args []float64, i int, def float64) float64 {
	if i < len(args) {
		return args[i]
	}
	return def
}

// SetMilliseconds is setMilliseconds or setUTCMilliseconds.
func (d *Date) SetMilliseconds(utc bool, ms float64) float64 {
	f, ok := d.pick(utc)
	if !ok {
		return NaN
	}
	day := math.Floor(d.base(utc) / msPerDay)
	return d.store(MakeDate(day, MakeTime(f.hours, f.minutes, f.seconds, ms)), utc)
}

// SetSeconds is setSeconds(s[, ms]) or its UTC form.
func (d *Date) SetSeconds(utc bool, s float64, rest ...float64) float64 {
	f, ok := d.pick(utc)
	if !ok {
		return NaN
	}
	day := math.Floor(d.base(utc) / msPerDay)
	return d.store(MakeDate(day, MakeTime(f.hours, f.minutes, s, optional(rest, 0, f.ms))), utc)
}

// SetMinutes is setMinutes(m[, s, ms]) or its UTC form.
func (d *Date) SetMinutes(utc bool, m float64, rest ...float64) float64 {
	f, ok := d.pick(utc)
	if !ok {
		return NaN
	}
	day := math.Floor(d.base(utc) / msPerDay)
	return d.store(MakeDate(day, MakeTime(f.hours, m, optional(rest, 0, f.seconds), optional(rest, 1, f.ms))), utc)
}

// SetHours is setHours(h[, m, s, ms]) or its UTC form.
func (d *Date) SetHours(utc bool, h float64, rest ...float64) float64 {
	f, ok := d.pick(utc)
	if !ok {
		return NaN
	}
	day := math.Floor(d.base(utc) / msPerDay)
	t := MakeTime(h, optional(rest, 0, f.minutes), optional(rest, 1, f.seconds), optional(rest, 2, f.ms))
	return d.store(MakeDate(day, t), utc)
}

// SetDate is setDate or setUTCDate.
func (d *Date) SetDate(utc bool, date float64) float64 {
	f, ok := d.pick(utc)
	if !ok {
		return NaN
	}
	return d.store(MakeDate(MakeDay(f.year, f.month, date), d.timeInDay(utc)), utc)
}

// SetMonth is setMonth(m[, date]) or its UTC form.
func (d *Date) SetMonth(utc bool, month float64, rest ...float64) float64 {
	f, ok := d.pick(utc)
	if !ok {
		return NaN
	}
	return d.store(MakeDate(MakeDay(f.year, month, optional(rest, 0, f.date)), d.timeInDay(utc)), utc)
}

// SetFullYear is setFullYear(y[, m, date]) or its UTC form; an invalid date starts from time value +0.
func (d *Date) SetFullYear(utc bool, year float64, rest ...float64) float64 {
	base := 0.0
	if !math.IsNaN(d.t) {
		base = d.base(utc)
	}
	f := split(base)
	b := int64(base)
	within := float64(b - floorDiv(b, msPerDay)*msPerDay)
	day := MakeDay(year, optional(rest, 0, f.month), optional(rest, 1, f.date))
	return d.store(MakeDate(day, within), utc)
}

func (d *Date) base(utc bool) float64 {
	if utc {
		return d.t
	}
	return LocalTime(d.t)
}

func (d *Date) timeInDay(utc bool) float64 {
	b := int64(d.base(utc))
	return float64(b - floorDiv(b, msPerDay)*msPerDay)
}

// NewLocalDate is new Date(year, month[, date, hours, minutes, seconds, ms]) with local fields; years 0 to 99
// mean 1900 to 1999.
func NewLocalDate(year, month float64, rest ...float64) *Date {
	return &Date{t: TimeClip(UTCFromLocal(dateValue(year, month, rest)))}
}

// DateUTC is Date.UTC(year, month[, date, hours, minutes, seconds, ms]).
func DateUTC(year, month float64, rest ...float64) float64 {
	return TimeClip(dateValue(year, month, rest))
}

func dateValue(year, month float64, rest []float64) float64 {
	if !math.IsNaN(year) {
		if y := ToIntegerOrInfinity(year); y >= 0 && y <= 99 {
			year = 1900 + y
		}
	}
	day := MakeDay(year, month, optional(rest, 0, 1))
	t := MakeTime(optional(rest, 1, 0), optional(rest, 2, 0), optional(rest, 3, 0), optional(rest, 4, 0))
	return MakeDate(day, t)
}

// ParseDate is new Date(string): the string grammar of Date.parse, local time when the string has no zone.
func ParseDate(s string) *Date {
	f, ok := ParseDateString(s)
	if !ok {
		return &Date{t: NaN}
	}
	date := MakeDate(MakeDay(f.Year, f.Month, f.Day), MakeTime(f.Hour, f.Minute, f.Second, f.Millisecond))
	if f.HasOffset {
		date -= f.Offset
		if !(date >= -maxTime && date <= maxTime) {
			return &Date{t: NaN}
		}
		return &Date{t: TimeClip(date)}
	}
	if !(date >= -maxTime && date <= maxTime) {
		return &Date{t: NaN}
	}
	return &Date{t: TimeClip(UTCFromLocal(date))}
}

// ISOString is Date.prototype.toISOString; it fails with the RangeError message for an invalid date.
func (d *Date) ISOString() (string, error) {
	if math.IsNaN(d.t) {
		return "", fmt.Errorf("Invalid time value")
	}
	f := split(d.t)
	year := fmt.Sprintf("%04d", int64(f.year))
	if f.year < 0 {
		year = fmt.Sprintf("-%06d", -int64(f.year))
	} else if f.year > 9999 {
		year = fmt.Sprintf("+%06d", int64(f.year))
	}
	return fmt.Sprintf("%s-%02d-%02dT%02d:%02d:%02d.%03dZ", year, int64(f.month)+1, int64(f.date),
		int64(f.hours), int64(f.minutes), int64(f.seconds), int64(f.ms)), nil
}

// String is a compact form of Date.prototype.toString used when a Date is converted to a string.
func (d *Date) String() string {
	s, err := d.ISOString()
	if err != nil {
		return "Invalid Date"
	}
	return s
}
