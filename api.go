package moment

import (
	"fmt"

	"github.com/moment/moment/internal/js"
)

// HTML5Formats is moment.HTML5_FMT.
var HTML5Formats = js.Obj(
	"DATETIME_LOCAL", "YYYY-MM-DDTHH:mm",
	"DATETIME_LOCAL_SECONDS", "YYYY-MM-DDTHH:mm:ss",
	"DATETIME_LOCAL_MS", "YYYY-MM-DDTHH:mm:ss.SSS",
	"DATE", "YYYY-MM-DD",
	"TIME", "HH:mm",
	"TIME_SECONDS", "HH:mm:ss",
	"TIME_MS", "HH:mm:ss.SSS",
	"WEEK", "GGGG-[W]WW",
	"MONTH", "YYYY-MM",
)

// ISO8601 and RFC2822 are the moment.ISO_8601 and moment.RFC_2822 format constants.
var (
	ISO8601 Value = isoConstant
	RFC2822 Value = rfcConstant
)

func init() {
	js.PrimitiveHook = func(v any, hintString bool) (any, bool) {
		switch x := v.(type) {
		case *Moment:
			if hintString {
				return x.String(), true
			}
			return x.ValueOf(), true
		case *Duration:
			if hintString {
				return x.ToISOString(), true
			}
			return x.AsMilliseconds(), true
		case *Locale:
			return "[object Object]", true
		}
		return nil, false
	}
}

// ErrUnknownName reports a function, method or property name the API does not have.
type ErrUnknownName struct{ Name string }

func (e *ErrUnknownName) Error() string { return "unknown name " + e.Name }

func stringsToValues(list []string) []any {
	out := make([]any, len(list))
	for i, s := range list {
		out[i] = s
	}
	return out
}

// CallStatic calls a function of the moment namespace by its JavaScript path ("moment", "moment.utc",
// "moment.duration.invalid", ...) with JavaScript argument values.
func CallStatic(path string, args []Value) (v Value, err error) {
	defer catchThrow(&err)
	a := func(i int) Value { return arg(args, i) }
	switch path {
	case "moment":
		return localMoment(args...), nil
	case "moment.utc":
		return utcMoment(args...), nil
	case "moment.unix":
		return unixMoment(a(0)), nil
	case "moment.parseZone":
		return parseZoneMoment(args...), nil
	case "moment.invalid":
		return CreateInvalid(a(0)), nil
	case "moment.duration":
		return CreateDuration(a(0), a(1)), nil
	case "moment.duration.invalid":
		return InvalidDuration(), nil
	case "moment.locale":
		return GetSetGlobalLocale(a(0), a(1)), nil
	case "moment.lang":
		langDeprecation.fire()
		return GetSetGlobalLocale(a(0), a(1)), nil
	case "moment.localeData":
		return LocaleData(a(0)), nil
	case "moment.langData":
		langDataDeprecation.fire()
		return LocaleData(a(0)), nil
	case "moment.defineLocale":
		return DefineLocaleValue(js.ToString(a(0)), a(1)), nil
	case "moment.updateLocale":
		return UpdateLocale(js.ToString(a(0)), a(1)), nil
	case "moment.locales":
		return stringsToValues(Locales()), nil
	case "moment.months":
		return ListMonths("months", a(0), a(1)), nil
	case "moment.monthsShort":
		return ListMonths("monthsShort", a(0), a(1)), nil
	case "moment.weekdays", "moment.weekdaysShort", "moment.weekdaysMin":
		return ListWeekdays(path[len("moment."):], a(0), a(1), a(2)), nil
	case "moment.min":
		return Min(args...), nil
	case "moment.max":
		return Max(args...), nil
	case "moment.now":
		return Now(), nil
	case "moment.isMoment":
		return IsMoment(a(0)), nil
	case "moment.isDate":
		return isDate(a(0)), nil
	case "moment.isDuration":
		return IsDuration(a(0)), nil
	case "moment.normalizeUnits":
		return NormalizeUnits(a(0)), nil
	case "moment.relativeTimeRounding":
		return RelativeTimeRounding(a(0)), nil
	case "moment.relativeTimeThreshold":
		return RelativeTimeThreshold(a(0), a(1)), nil
	case "moment.calendarFormat":
		m, ok := a(0).(*Moment)
		if !ok {
			return nil, typeError("myMoment.diff is not a function")
		}
		return CalendarFormat(m, a(1)), nil
	case "moment.parseTwoDigitYear":
		return ParseTwoDigitYear(a(0)), nil
	case "moment.ISO_8601", "moment.RFC_2822":
		return js.Undef, nil
	}
	return nil, &ErrUnknownName{path}
}

// GetProperty reads a property of the moment namespace by path.
func GetProperty(path string) (Value, error) {
	switch path {
	case "moment.version":
		return Version, nil
	case "moment.defaultFormat":
		return DefaultFormat, nil
	case "moment.defaultFormatUtc":
		return DefaultFormatUtc, nil
	case "moment.HTML5_FMT":
		return HTML5Formats, nil
	case "moment.ISO_8601":
		return ISO8601, nil
	case "moment.RFC_2822":
		return RFC2822, nil
	}
	const prefix = "moment.HTML5_FMT."
	if len(path) > len(prefix) && path[:len(prefix)] == prefix {
		return HTML5Formats.Get(path[len(prefix):]), nil
	}
	return nil, &ErrUnknownName{path}
}

// SetProperty assigns moment.defaultFormat or moment.defaultFormatUtc.
func SetProperty(path string, v Value) error {
	switch path {
	case "moment.defaultFormat":
		DefaultFormat = v
	case "moment.defaultFormatUtc":
		DefaultFormatUtc = v
	default:
		return &ErrUnknownName{path}
	}
	return nil
}

// CallMethod calls a method of a moment, duration or locale by its JavaScript name.
func CallMethod(recv Value, name string, args []Value) (v Value, err error) {
	defer catchThrow(&err)
	switch r := recv.(type) {
	case *Moment:
		return r.call(name, args)
	case *Duration:
		return r.call(name, args)
	case *Locale:
		return r.call(name, args)
	}
	return nil, typeError(fmt.Sprintf("%s is not a function", name))
}

func errorOr(v Value, err error) (Value, error) {
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *Moment) call(name string, args []Value) (Value, error) {
	a := func(i int) Value { return arg(args, i) }
	switch name {
	case "add":
		return m.Add(a(0), a(1)), nil
	case "subtract":
		return m.Subtract(a(0), a(1)), nil
	case "calendar":
		return m.Calendar(args...), nil
	case "clone":
		return m.Clone(), nil
	case "diff":
		return m.Diff(a(0), a(1), a(2)), nil
	case "startOf":
		return m.StartOf(a(0)), nil
	case "endOf":
		return m.EndOf(a(0)), nil
	case "format":
		return m.Format(a(0)), nil
	case "from":
		return m.From(a(0), a(1)), nil
	case "fromNow":
		return m.FromNow(a(0)), nil
	case "to":
		return m.To(a(0), a(1)), nil
	case "toNow":
		return m.ToNow(a(0)), nil
	case "get":
		return m.Get(a(0))
	case "set":
		return m.Set(a(0), a(1)), nil
	case "invalidAt":
		return m.InvalidAt(), nil
	case "isAfter":
		return m.IsAfter(a(0), a(1)), nil
	case "isBefore":
		return m.IsBefore(a(0), a(1)), nil
	case "isBetween":
		return m.IsBetween(a(0), a(1), a(2), a(3)), nil
	case "isSame":
		return m.IsSame(a(0), a(1)), nil
	case "isSameOrAfter":
		return m.IsSameOrAfter(a(0), a(1)), nil
	case "isSameOrBefore":
		return m.IsSameOrBefore(a(0), a(1)), nil
	case "isValid":
		return m.IsValid(), nil
	case "lang":
		momentLangDeprecated.fire()
		if js.IsUndefined(a(0)) {
			return m.LocaleData(), nil
		}
		return m.SetLocale(a(0)), nil
	case "locale":
		if js.IsUndefined(a(0)) {
			return m.Locale(), nil
		}
		return m.SetLocale(a(0)), nil
	case "localeData":
		return m.LocaleData(), nil
	case "min":
		return m.PrototypeMin(args...), nil
	case "max":
		return m.PrototypeMax(args...), nil
	case "parsingFlags":
		return m.ParsingFlags(), nil
	case "toArray":
		return m.ToArray(), nil
	case "toObject":
		return m.ToObject(), nil
	case "toDate":
		return m.ToDate(), nil
	case "toISOString":
		return m.ToISOString(a(0))
	case "toJSON":
		return m.ToJSON()
	case "inspect":
		return m.Inspect(), nil
	case "toString":
		return m.String(), nil
	case "unix":
		return m.Unix(), nil
	case "valueOf":
		return m.ValueOf(), nil
	case "creationData":
		return m.CreationData(), nil
	case "eraName":
		return m.EraName(), nil
	case "eraNarrow":
		return m.EraNarrow(), nil
	case "eraAbbr":
		return m.EraAbbr(), nil
	case "eraYear":
		return m.EraYear(), nil
	case "isLeapYear":
		return m.IsLeapYear(), nil
	case "daysInMonth":
		return m.DaysInMonth(), nil
	case "weeksInYear":
		return m.WeeksInYear(), nil
	case "weeksInWeekYear":
		return m.WeeksInWeekYear(), nil
	case "isoWeeksInYear":
		return m.ISOWeeksInYear(), nil
	case "isoWeeksInISOWeekYear":
		return m.ISOWeeksInISOWeekYear(), nil
	case "utcOffset":
		return m.GetSetOffset(a(0), a(1), a(2)), nil
	case "utc":
		return m.UTC(a(0)), nil
	case "local":
		return m.Local(a(0)), nil
	case "parseZone":
		return m.ParseZone(), nil
	case "hasAlignedHourOffset":
		return m.HasAlignedHourOffset(a(0)), nil
	case "isDST":
		return m.IsDST(), nil
	case "isLocal":
		return m.IsLocal(), nil
	case "isUtcOffset":
		return m.IsUtcOffset(), nil
	case "isUtc", "isUTC":
		return m.IsUTC(), nil
	case "zoneAbbr":
		return m.ZoneAbbr(), nil
	case "zoneName":
		return m.ZoneName(), nil
	case "zone":
		return m.GetSetZone(a(0), a(1)), nil
	case "isDSTShifted":
		return m.IsDSTShifted(), nil
	case "dates":
		datesDeprecation.fire()
		return m.unitAccessor("date", a(0)), nil
	case "months":
		monthsDeprecation.fire()
		return m.unitAccessor("month", a(0)), nil
	case "years":
		yearsDeprecation.fire()
		return m.unitAccessor("year", a(0)), nil
	}
	if unit, ok := momentUnitMethods[name]; ok {
		return m.unitAccessor(unit, a(0)), nil
	}
	return nil, typeError(fmt.Sprintf("m.%s is not a function", name))
}

var momentUnitMethods = map[string]string{
	"year": "year", "weekYear": "weekYear", "isoWeekYear": "isoWeekYear", "quarter": "quarter",
	"quarters": "quarter", "month": "month", "week": "week", "weeks": "week", "isoWeek": "isoWeek",
	"isoWeeks": "isoWeek", "date": "date", "day": "day", "days": "day", "weekday": "weekday",
	"isoWeekday": "isoWeekday", "dayOfYear": "dayOfYear", "hour": "hour", "hours": "hour", "minute": "minute",
	"minutes": "minute", "second": "second", "seconds": "second", "millisecond": "millisecond",
	"milliseconds": "millisecond",
}

func (d *Duration) call(name string, args []Value) (Value, error) {
	a := func(i int) Value { return arg(args, i) }
	switch name {
	case "isValid":
		return d.IsValid(), nil
	case "abs":
		return d.Abs(), nil
	case "add":
		return d.Add(a(0), a(1)), nil
	case "subtract":
		return d.Subtract(a(0), a(1)), nil
	case "as":
		return errorOr(d.As(a(0)))
	case "asMilliseconds", "valueOf":
		return d.as("ms"), nil
	case "asSeconds":
		return d.as("s"), nil
	case "asMinutes":
		return d.as("m"), nil
	case "asHours":
		return d.as("h"), nil
	case "asDays":
		return d.as("d"), nil
	case "asWeeks":
		return d.as("w"), nil
	case "asMonths":
		return d.as("M"), nil
	case "asQuarters":
		return d.as("Q"), nil
	case "asYears":
		return d.as("y"), nil
	case "clone":
		return d.Clone(), nil
	case "get":
		return d.Get(a(0))
	case "milliseconds":
		return d.Milliseconds(), nil
	case "seconds":
		return d.Seconds(), nil
	case "minutes":
		return d.Minutes(), nil
	case "hours":
		return d.Hours(), nil
	case "days":
		return d.Days(), nil
	case "weeks":
		return d.Weeks(), nil
	case "months":
		return d.Months(), nil
	case "years":
		return d.Years(), nil
	case "humanize":
		return d.Humanize(a(0), a(1)), nil
	case "toISOString", "toString", "toJSON":
		return d.ToISOString(), nil
	case "toIsoString":
		toIsoStringDeprecation.fire()
		return d.ToISOString(), nil
	case "locale":
		if js.IsUndefined(a(0)) {
			return d.Locale(), nil
		}
		return d.SetLocale(a(0)), nil
	case "localeData":
		return d.LocaleData(), nil
	case "lang":
		momentLangDeprecated.fire()
		if js.IsUndefined(a(0)) {
			return d.LocaleData(), nil
		}
		return d.SetLocale(a(0)), nil
	}
	return nil, typeError(fmt.Sprintf("d.%s is not a function", name))
}

func momentArg(v Value) *Moment {
	m, _ := v.(*Moment)
	return m
}

func (l *Locale) call(name string, args []Value) (Value, error) {
	a := func(i int) Value { return arg(args, i) }
	switch name {
	case "calendar":
		return l.Calendar(a(0), momentArg(a(1)), momentArg(a(2))), nil
	case "longDateFormat":
		return l.LongDateFormat(a(0)), nil
	case "invalidDate":
		return l.InvalidDate(), nil
	case "ordinal":
		return l.Ordinal(a(0), a(1)), nil
	case "preparse":
		return l.Preparse(js.ToString(a(0))), nil
	case "postformat":
		return l.Postformat(js.ToString(a(0))), nil
	case "relativeTime":
		return l.RelativeTime(a(0), a(1), a(2), a(3)), nil
	case "pastFuture":
		return l.PastFuture(a(0), a(1)), nil
	case "months":
		return l.Months(momentArg(a(0)), a(1)), nil
	case "monthsShort":
		return l.MonthsShort(momentArg(a(0)), a(1)), nil
	case "monthsParse":
		return l.MonthsParse(js.ToString(a(0)), a(1), js.Truthy(a(2))), nil
	case "monthsRegex":
		return regexValue(l.MonthsRegex(js.Truthy(a(0)))), nil
	case "monthsShortRegex":
		return regexValue(l.MonthsShortRegex(js.Truthy(a(0)))), nil
	case "week":
		m := momentArg(a(0))
		if m == nil {
			return nil, typeError("mom.year is not a function")
		}
		return l.Week(m), nil
	case "firstDayOfYear":
		return l.FirstDayOfYear(), nil
	case "firstDayOfWeek":
		return l.FirstDayOfWeek(), nil
	case "weekdays":
		return l.Weekdays(weekdayArg(a(0)), a(1)), nil
	case "weekdaysShort":
		return l.WeekdaysShort(weekdayArg(a(0))), nil
	case "weekdaysMin":
		return l.WeekdaysMin(weekdayArg(a(0))), nil
	case "weekdaysParse":
		return l.WeekdaysParse(js.ToString(a(0)), a(1), js.Truthy(a(2))), nil
	case "weekdaysRegex":
		return regexValue(l.WeekdaysRegex(js.Truthy(a(0)))), nil
	case "weekdaysShortRegex":
		return regexValue(l.WeekdaysShortRegex(js.Truthy(a(0)))), nil
	case "weekdaysMinRegex":
		return regexValue(l.WeekdaysMinRegex(js.Truthy(a(0)))), nil
	case "isPM":
		return l.IsPM(a(0)), nil
	case "meridiem":
		return l.Meridiem(js.ToNumber(a(0)), js.ToNumber(a(1)), js.Truthy(a(2))), nil
	case "eras":
		return l.Eras(), nil
	case "erasParse":
		return l.ErasParse(js.ToString(a(0)), a(1), js.Truthy(a(2))), nil
	case "erasConvertYear":
		return l.ErasConvertYear(a(0), a(1)), nil
	case "erasAbbrRegex":
		return regexValue(l.ErasAbbrRegex(js.Truthy(a(0)))), nil
	case "erasNameRegex":
		return regexValue(l.ErasNameRegex(js.Truthy(a(0)))), nil
	case "erasNarrowRegex":
		return regexValue(l.ErasNarrowRegex(js.Truthy(a(0)))), nil
	}
	return nil, typeError(fmt.Sprintf("l.%s is not a function", name))
}

func weekdayArg(v Value) WeekdayArg {
	switch x := v.(type) {
	case bool:
		return x
	case *Moment:
		return x
	}
	return nil
}

func regexValue[T any](re *T) Value {
	if re == nil {
		return js.Undef
	}
	return re
}

// catchThrow turns a JavaScript exception raised deep inside parsing (a *ThrowError panic, the library's model of
// JavaScript's throw) into the returned error of a public entry point. Any other panic is a bug and is re-raised.
func catchThrow(err *error) {
	if r := recover(); r != nil {
		t, ok := r.(*ThrowError)
		if !ok {
			panic(r)
		}
		*err = t
	}
}

// New is moment(input, format, locale, strict): a moment in local time. An error the original throws while parsing
// (such as the TypeError of a "Do" token that matches nothing) is returned as a *ThrowError.
func New(args ...Value) (m *Moment, err error) {
	defer catchThrow(&err)
	return localMoment(args...), nil
}

// CreateUTC is moment.utc(input, format, locale, strict), returning what the original throws as a *ThrowError.
func CreateUTC(args ...Value) (m *Moment, err error) {
	defer catchThrow(&err)
	return utcMoment(args...), nil
}

// Unix is moment.unix(seconds).
func Unix(seconds Value) (m *Moment, err error) {
	defer catchThrow(&err)
	return unixMoment(seconds), nil
}

// ParseZone is moment.parseZone(input, format, locale, strict), returning what the original throws as a *ThrowError.
func ParseZone(args ...Value) (m *Moment, err error) {
	defer catchThrow(&err)
	return parseZoneMoment(args...), nil
}
