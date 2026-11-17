package moment

import (
	"math"

	"github.com/moment/moment/internal/js"
)

// Indexes of the date-parts array and the overflow units reported by invalidAt.
const (
	yearIndex = iota
	monthIndex
	dateIndex
	hourIndex
	minuteIndex
	secondIndex
	millisecondIndex
	weekIndex
	weekdayIndex
)

// createDate is new Date(y, m, d, h, M, s, ms) without the remapping of years 0 to 99 to the 1900s.
func createDate(parts ...Value) *js.Date {
	n := make([]float64, 7)
	for i := range n {
		n[i] = js.ToNumber(arg(parts, i))
	}
	y := js.ToNumber(arg(parts, 0))
	if y < 100 && y >= 0 {
		date := js.NewLocalDate(js.ToNumber(js.Add(arg(parts, 0), 400.0)), n[1], n[2], n[3], n[4], n[5], n[6])
		if fy := date.FullYear(false); !math.IsNaN(fy) && !math.IsInf(fy, 0) {
			date.SetFullYear(false, y)
		}
		return date
	}
	return js.NewLocalDate(n[0], n[1], n[2:]...)
}

// createUTCDate is new Date(Date.UTC(parts...)) without the remapping of years 0 to 99 to the 1900s.
func createUTCDate(parts ...float64) *js.Date {
	return createUTCDateFrom(parts[0], parts[1:]...)
}

// createUTCDateFrom is createUTCDate with the year as any JavaScript value: a string year below 100 is
// concatenated with 400 (as `y + 400` does in JavaScript) before Date.UTC converts it.
func createUTCDateFrom(year Value, rest ...float64) *js.Date {
	y := js.ToNumber(year)
	first := y
	small := y < 100 && y >= 0
	if small {
		first = js.ToNumber(js.Add(year, 400.0))
	}
	month := 0.0
	if len(rest) > 0 {
		month = rest[0]
	}
	var others []float64
	if len(rest) > 1 {
		others = rest[1:]
	}
	date := js.NewDate(js.DateUTC(first, month, others...))
	if small {
		if fy := date.FullYear(true); !math.IsNaN(fy) && !math.IsInf(fy, 0) {
			date.SetFullYear(true, y)
		}
	}
	return date
}

func currentDateNow(config *Moment) float64 {
	now := config.defaultDatePartsNow
	if now == nil {
		return Now()
	}
	if !now.has {
		now.value = Now()
		now.has = true
	}
	return now.value
}

// defaultDateParts returns today's year, month and date, as the local moment of now when forWeek, otherwise in
// the config's UTC or local time.
func defaultDateParts(config *Moment, now float64, forWeek bool) []float64 {
	if forWeek {
		m := localMoment(now)
		return []float64{m.Year(), m.Month(), m.Date()}
	}
	d := js.NewDate(now)
	utc := js.Truthy(config.useUTC)
	return []float64{d.FullYear(utc), d.Month(utc), d.DateOfMonth(utc)}
}

// configFromArray builds the date from config.a, defaulting the missing parts.
func configFromArray(config *Moment) {
	if config.d != nil {
		return
	}
	var now float64
	var currentDate []float64
	if js.IsNullish(config.aGet(yearIndex)) || js.IsNullish(config.aGet(monthIndex)) ||
		js.IsNullish(config.aGet(dateIndex)) {
		now = currentDateNow(config)
		currentDate = defaultDateParts(config, now, false)
	}
	if config.w != nil && js.IsNullish(config.aGet(dateIndex)) && js.IsNullish(config.aGet(monthIndex)) {
		dayOfYearFromWeekInfo(config, defaultDateParts(config, now, true))
	}
	if !js.IsNullish(config.dayOfYear) {
		yearToUse := config.aGet(yearIndex)
		if js.IsNullish(yearToUse) {
			yearToUse = currentDate[yearIndex]
		}
		doy := js.ToNumber(config.dayOfYear)
		if doy > daysInYear(js.ToNumber(yearToUse)) || doy == 0 {
			getParsingFlags(config).Set("_overflowDayOfYear", true)
		}
		date := createUTCDate(js.ToNumber(yearToUse), 0, doy)
		config.aSet(monthIndex, date.Month(true))
		config.aSet(dateIndex, date.DateOfMonth(true))
	}
	dateIsDefaulted := js.IsNullish(config.aGet(yearIndex)) || js.IsNullish(config.aGet(monthIndex)) ||
		js.IsNullish(config.aGet(dateIndex))
	input := make([]Value, 7)
	i := 0
	for ; i < 3 && js.IsNullish(config.aGet(i)); i++ {
		config.aSet(i, currentDate[i])
		input[i] = currentDate[i]
	}
	for ; i < 7; i++ {
		v := config.aGet(i)
		if js.IsNullish(v) {
			v = 0.0
			if i == 2 {
				v = 1.0
			}
		}
		config.aSet(i, v)
		input[i] = v
	}
	if js.StrictEquals(config.aGet(hourIndex), 24.0) && js.StrictEquals(config.aGet(minuteIndex), 0.0) &&
		js.StrictEquals(config.aGet(secondIndex), 0.0) && js.StrictEquals(config.aGet(millisecondIndex), 0.0) {
		config.nextDay = true
		config.aSet(hourIndex, 0.0)
	}
	utc := js.Truthy(config.useUTC)
	if utc {
		nums := make([]float64, 7)
		for j, v := range input {
			nums[j] = js.ToNumber(v)
		}
		config.d = createUTCDate(nums...)
	} else {
		config.d = createDate(input...)
	}
	expectedWeekday := config.d.Day(utc)
	if !js.IsNullish(config.tzm) {
		config.d.SetMinutes(true, config.d.Minutes(true)-js.ToNumber(config.tzm))
	}
	if js.Truthy(config.nextDay) {
		config.aSet(hourIndex, 24.0)
	}
	if config.w != nil && !js.IsUndefined(config.w.Get("d")) && !dateIsDefaulted &&
		!js.StrictEquals(config.w.Get("d"), expectedWeekday) {
		getParsingFlags(config).Set("weekdayMismatch", true)
	}
}

func defaults(values ...Value) Value {
	for _, v := range values[:len(values)-1] {
		if !js.IsNullish(v) {
			return v
		}
	}
	return values[len(values)-1]
}

func dayOfYearFromWeekInfo(config *Moment, currentDate []float64) {
	w := config.w
	var weekYear, week, weekday Value
	var dow, doy float64
	weekdayOverflow := false
	if !js.IsNullish(w.Get("GG")) || !js.IsNullish(w.Get("W")) || !js.IsNullish(w.Get("E")) {
		dow, doy = 1, 4
		weekYear = defaults(w.Get("GG"), config.aGet(yearIndex),
			weekOfYearFromDate(currentDate[0], currentDate[1], currentDate[2], 1, 4).year)
		week = defaults(w.Get("W"), 1.0)
		weekday = defaults(w.Get("E"), 1.0)
		if n := js.ToNumber(weekday); n < 1 || n > 7 {
			weekdayOverflow = true
		}
	} else {
		dow = js.ToNumber(config.locale.weekSetting("dow"))
		doy = js.ToNumber(config.locale.weekSetting("doy"))
		cur := weekOfYearFromDate(currentDate[0], currentDate[1], currentDate[2], dow, doy)
		weekYear = defaults(w.Get("gg"), config.aGet(yearIndex), cur.year)
		week = defaults(w.Get("w"), cur.week)
		switch {
		case !js.IsNullish(w.Get("d")):
			weekday = w.Get("d")
			if n := js.ToNumber(weekday); n < 0 || n > 6 {
				weekdayOverflow = true
			}
		case !js.IsNullish(w.Get("e")):
			weekday = js.ToNumber(w.Get("e")) + dow
			if n := js.ToNumber(w.Get("e")); n < 0 || n > 6 {
				weekdayOverflow = true
			}
		default:
			weekday = dow
		}
	}
	wy, wk := js.ToNumber(weekYear), js.ToNumber(week)
	if wk < 1 || wk > weeksInYear(wy, dow, doy) {
		getParsingFlags(config).Set("_overflowWeeks", true)
	} else if weekdayOverflow {
		getParsingFlags(config).Set("_overflowWeekday", true)
	} else {
		res := dayOfYearFromWeeks(wy, wk, js.ToNumber(weekday), dow, doy)
		config.aSet(yearIndex, js.ToNumber(res.year))
		config.dayOfYear = res.dayOfYear
	}
}

// configFromObject builds the date from an object of units.
func configFromObject(config *Moment) {
	if config.d != nil {
		return
	}
	i := normalizeObjectUnits(config.i)
	dayOrDate := i.Get("day")
	if js.IsUndefined(dayOrDate) {
		dayOrDate = i.Get("date")
	}
	config.a = nil
	for _, v := range []Value{i.Get("year"), i.Get("month"), dayOrDate, i.Get("hour"), i.Get("minute"),
		i.Get("second"), i.Get("millisecond")} {
		if js.Truthy(v) {
			v = js.ParseInt(v)
		}
		config.a = append(config.a, v)
	}
	configFromArray(config)
}

// checkOverflow records in the parsing flags which date part is out of range.
func checkOverflow(m *Moment) *Moment {
	if m.a == nil || !js.StrictEquals(getParsingFlags(m).Get("overflow"), -2.0) {
		return m
	}
	n := func(i int) float64 { return js.ToNumber(m.aGet(i)) }
	is := func(i int, v float64) bool { return js.StrictEquals(m.aGet(i), v) }
	var overflow float64
	switch {
	case n(monthIndex) < 0 || n(monthIndex) > 11:
		overflow = monthIndex
	case n(dateIndex) < 1 || n(dateIndex) > daysInMonth(n(yearIndex), n(monthIndex)):
		overflow = dateIndex
	case n(hourIndex) < 0 || n(hourIndex) > 24 ||
		(is(hourIndex, 24) && (!is(minuteIndex, 0) || !is(secondIndex, 0) || !is(millisecondIndex, 0))):
		overflow = hourIndex
	case n(minuteIndex) < 0 || n(minuteIndex) > 59:
		overflow = minuteIndex
	case n(secondIndex) < 0 || n(secondIndex) > 59:
		overflow = secondIndex
	case n(millisecondIndex) < 0 || n(millisecondIndex) > 999:
		overflow = millisecondIndex
	default:
		overflow = -1
	}
	pf := getParsingFlags(m)
	if js.Truthy(pf.Get("_overflowDayOfYear")) && (overflow < yearIndex || overflow > dateIndex) {
		overflow = dateIndex
	}
	if js.Truthy(pf.Get("_overflowWeeks")) && overflow == -1 {
		overflow = weekIndex
	}
	if js.Truthy(pf.Get("_overflowWeekday")) && overflow == -1 {
		overflow = weekdayIndex
	}
	pf.Set("overflow", overflow)
	return m
}

type weekYear struct {
	week, year float64
}

type yearDay struct {
	year      Value
	dayOfYear float64
}

// firstWeekOffset is the start of the first week minus the start of the year, in days.
func firstWeekOffset(year Value, dow, doy float64) float64 {
	fwd := 7 + dow - doy
	fwdlw := math.Mod(7+createUTCDateFrom(year, 0, fwd).Day(true)-dow, 7)
	return -fwdlw + fwd - 1
}

// dayOfYearFromWeeks converts a week date to a year and day of the year; the year keeps its JavaScript type
// (a string year stays a string unless arithmetic changes it).
func dayOfYearFromWeeks(year Value, week, weekday, dow, doy float64) yearDay {
	localWeekday := math.Mod(7+weekday-dow, 7)
	weekOffset := firstWeekOffset(year, dow, doy)
	dayOfYear := 1 + float64(7*(week-1)) + localWeekday + weekOffset
	y := js.ToNumber(year)
	switch {
	case dayOfYear <= 0:
		return yearDay{y - 1, daysInYear(y-1) + dayOfYear}
	case dayOfYear > daysInYear(y):
		return yearDay{js.Add(year, 1.0), dayOfYear - daysInYear(y)}
	}
	return yearDay{year, dayOfYear}
}

func weekOfYearFromDayOfYear(year, dayOfYear, dow, doy float64) weekYear {
	weekOffset := firstWeekOffset(year, dow, doy)
	week := math.Floor((dayOfYear-weekOffset-1)/7) + 1
	switch {
	case week < 1:
		return weekYear{week + weeksInYear(year-1, dow, doy), year - 1}
	case week > weeksInYear(year, dow, doy):
		return weekYear{week - weeksInYear(year, dow, doy), year + 1}
	}
	return weekYear{week, year}
}

func weekOfYear(m *Moment, dow, doy float64) weekYear {
	return weekOfYearFromDayOfYear(m.Year(), m.DayOfYear(), dow, doy)
}

func weekOfYearFromDate(year, month, date, dow, doy float64) weekYear {
	diff := createUTCDate(year, month, date).Time() - createUTCDate(year, 0, 1).Time()
	return weekOfYearFromDayOfYear(year, js.Round(diff/864e5)+1, dow, doy)
}

func weeksInYear(year Value, dow, doy float64) float64 {
	weekOffset := firstWeekOffset(year, dow, doy)
	next := firstWeekOffset(js.Add(year, 1.0), dow, doy)
	return (daysInYear(js.ToNumber(year)) - weekOffset + next) / 7
}
