package moment

import (
	"math"
	"strings"

	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

// Duration is a moment.duration: milliseconds, days and months kept apart, and their bubbled-up units.
type Duration struct {
	milliseconds, days, months float64
	data                       durationData
	locale                     *Locale
	valid                      bool
}

type durationData struct {
	milliseconds, seconds, minutes, hours, days, months, years float64
}

var (
	aspNetDurationRegex = jsre.MustCompile(`^(-|\+)?(?:(\d*)[. ])?(\d+):(\d+)(?::(\d+)(\.\d*)?)?$`, "")
	isoDurationRegex    = jsre.MustCompile(`^(-|\+)?P(?:([-+]?[0-9,.]*)Y)?(?:([-+]?[0-9,.]*)M)?(?:([-+]?[0-9,.]*)W)?`+
		`(?:([-+]?[0-9,.]*)D)?(?:T(?:([-+]?[0-9,.]*)H)?(?:([-+]?[0-9,.]*)M)?(?:([-+]?[0-9,.]*)S)?)?$`, "")
	isoTrailingZeros = jsre.MustCompile(`\.?0+$`, "")
	durationOrdering = []string{"year", "quarter", "month", "week", "day", "hour", "minute", "second",
		"millisecond"}
)

// IsDuration is moment.isDuration.
func IsDuration(v Value) bool {
	_, ok := v.(*Duration)
	return ok
}

// CreateDuration is moment.duration(input, key).
func CreateDuration(input, key Value) *Duration {
	var spec Value = input
	switch {
	case IsDuration(input):
		d := input.(*Duration)
		spec = js.Obj("ms", d.milliseconds, "d", d.days, "M", d.months)
	case isNumberValue(input) || !js.IsNaNValue(input):
		o := js.NewObject()
		if js.Truthy(key) {
			o.Set(js.ToString(key), js.ToNumber(input))
		} else {
			o.Set("milliseconds", js.ToNumber(input))
		}
		spec = o
	default:
		s := js.ToString(input)
		if match := aspNetDurationRegex.Exec(s); match != nil {
			sign := 1.0
			if match.Groups[1] == "-" && match.Defined[1] {
				sign = -1
			}
			g := func(i int) Value {
				if match.Defined[i] {
					return match.Groups[i]
				}
				return js.Undef
			}
			ms := toInt(absRound(js.ToNumber(g(6)) * 1000))
			spec = js.Obj("y", 0.0, "d", toInt(g(2))*sign, "h", toInt(g(3))*sign, "m", toInt(g(4))*sign,
				"s", toInt(g(5))*sign, "ms", ms*sign)
		} else if match := isoDurationRegex.Exec(s); match != nil {
			sign := 1.0
			if match.Groups[1] == "-" && match.Defined[1] {
				sign = -1
			}
			p := func(i int) float64 {
				res := 0.0
				if match.Defined[i] && match.Groups[i] != "" {
					res = js.ParseFloat(js.ReplaceFirst(match.Groups[i], ",", "."))
					if math.IsNaN(res) {
						res = 0
					}
				}
				return res * sign
			}
			spec = js.Obj("y", p(2), "M", p(3), "w", p(4), "d", p(5), "h", p(6), "m", p(7), "s", p(8))
		} else if js.IsNullish(input) {
			spec = js.NewObject()
		} else if o, ok := input.(*js.Object); ok && (o.Has("from") || o.Has("to")) {
			months, ms := momentsDifference(localMoment(o.Get("from")), localMoment(o.Get("to")))
			spec = js.Obj("ms", ms, "M", months)
		}
	}
	ret := newDuration(spec)
	if d, ok := input.(*Duration); ok {
		ret.locale = d.locale
		ret.valid = d.valid
	}
	return ret
}

func isNumberValue(v Value) bool {
	_, ok := v.(float64)
	return ok
}

// InvalidDuration is moment.duration.invalid().
func InvalidDuration() *Duration { return CreateDuration(js.NaN, js.Undef) }

func newDuration(spec Value) *Duration {
	n := normalizeObjectUnits(spec)
	unit := func(names ...string) Value {
		for _, name := range names {
			if v := n.Get(name); js.Truthy(v) {
				return v
			}
		}
		return 0.0
	}
	years, quarters, months := js.ToNumber(unit("year")), js.ToNumber(unit("quarter")), js.ToNumber(unit("month"))
	weeks, days, hours := js.ToNumber(unit("week", "isoWeek")), unit("day"), js.ToNumber(unit("hour"))
	minutes, seconds, ms := js.ToNumber(unit("minute")), js.ToNumber(unit("second")), unit("millisecond")
	d := &Duration{valid: isDurationValid(n), locale: getLocale(js.Undef)}
	total := js.ToNumber(ms) + float64(seconds*1e3)
	total += float64(minutes * 6e4)
	d.milliseconds = total + float64(float64(float64(hours*1000)*60)*60)
	d.days = js.ToNumber(days) + float64(weeks*7)
	d.months = js.ToNumber(months) + float64(quarters*3) + float64(years*12)
	d.bubble()
	return d
}

func isDurationValid(n *js.Object) bool {
	for _, k := range n.Keys() {
		known := false
		for _, u := range durationOrdering {
			if u == k {
				known = true
				break
			}
		}
		if !known || !(js.IsNullish(n.Get(k)) || !js.IsNaNValue(n.Get(k))) {
			return false
		}
	}
	unitHasDecimal := false
	for _, u := range durationOrdering {
		v := n.Get(u)
		if !js.Truthy(v) {
			continue
		}
		if unitHasDecimal {
			return false
		}
		if js.ParseFloat(v) != toInt(v) {
			unitHasDecimal = true
		}
	}
	return true
}

func daysToMonths(days float64) float64 { return float64(days*4800) / 146097 }

func monthsToDays(months float64) float64 { return float64(months*146097) / 4800 }

func (d *Duration) bubble() {
	ms, days, months := d.milliseconds, d.days, d.months
	if !((ms >= 0 && days >= 0 && months >= 0) || (ms <= 0 && days <= 0 && months <= 0)) {
		ms += float64(absCeil(monthsToDays(months)+days) * 864e5)
		days, months = 0, 0
	}
	d.data.milliseconds = math.Mod(ms, 1000)
	seconds := absFloor(ms / 1000)
	d.data.seconds = math.Mod(seconds, 60)
	minutes := absFloor(seconds / 60)
	d.data.minutes = math.Mod(minutes, 60)
	hours := absFloor(minutes / 60)
	d.data.hours = math.Mod(hours, 24)
	days += absFloor(hours / 24)
	monthsFromDays := absFloor(daysToMonths(days))
	months += monthsFromDays
	days -= absCeil(monthsToDays(monthsFromDays))
	years := absFloor(months / 12)
	months = math.Mod(months, 12)
	d.data.days, d.data.months, d.data.years = days, months, years
}

func positiveMomentsDifference(base, other *Moment) (months, ms float64) {
	months = other.Month() - base.Month() + float64((other.Year()-base.Year())*12)
	if base.Clone().Add(months, "M").IsAfter(other, js.Undef) {
		months--
	}
	ms = other.ValueOf() - base.Clone().Add(months, "M").ValueOf()
	return months, ms
}

func momentsDifference(base, other *Moment) (months, ms float64) {
	if !(base.IsValid() && other.IsValid()) {
		return 0, 0
	}
	other = cloneWithOffset(other, base)
	if base.IsBefore(other, js.Undef) {
		return positiveMomentsDifference(base, other)
	}
	months, ms = positiveMomentsDifference(other, base)
	return -months, -ms
}

// IsValid is duration#isValid.
func (d *Duration) IsValid() bool { return d.valid }

// Abs is duration#abs; note that it leaves the bubbled days as they are, as the original does.
func (d *Duration) Abs() *Duration {
	d.milliseconds, d.days, d.months = math.Abs(d.milliseconds), math.Abs(d.days), math.Abs(d.months)
	d.data.milliseconds = math.Abs(d.data.milliseconds)
	d.data.seconds = math.Abs(d.data.seconds)
	d.data.minutes = math.Abs(d.data.minutes)
	d.data.hours = math.Abs(d.data.hours)
	d.data.months = math.Abs(d.data.months)
	d.data.years = math.Abs(d.data.years)
	return d
}

// Add is duration#add(input, value).
func (d *Duration) Add(input, value Value) *Duration { return d.addSubtract(input, value, 1) }

// Subtract is duration#subtract(input, value).
func (d *Duration) Subtract(input, value Value) *Duration { return d.addSubtract(input, value, -1) }

func (d *Duration) addSubtract(input, value Value, direction float64) *Duration {
	other := CreateDuration(input, value)
	d.milliseconds += float64(direction * other.milliseconds)
	d.days += float64(direction * other.days)
	d.months += float64(direction * other.months)
	d.bubble()
	return d
}

// As is duration#as(units); an unknown unit is an Error "Unknown unit <units>".
func (d *Duration) As(units Value) (float64, error) {
	if !d.valid {
		return js.NaN, nil
	}
	unit := normalizeUnit(units)
	ms := d.milliseconds
	if unit == "month" || unit == "quarter" || unit == "year" {
		days := d.days + ms/864e5
		months := d.months + daysToMonths(days)
		switch unit {
		case "month":
			return months, nil
		case "quarter":
			return months / 3, nil
		}
		return months / 12, nil
	}
	days := d.days + js.Round(monthsToDays(d.months))
	switch unit {
	case "week":
		return days/7 + ms/6048e5, nil
	case "day":
		return days + ms/864e5, nil
	case "hour":
		return float64(days*24) + ms/36e5, nil
	case "minute":
		return float64(days*1440) + ms/6e4, nil
	case "second":
		return float64(days*86400) + ms/1000, nil
	case "millisecond":
		return math.Floor(days*864e5) + ms, nil
	}
	return 0, &ThrowError{Name: "Error", Message: "Unknown unit " + js.ToString(NormalizeUnits(units))}
}

func (d *Duration) as(unit string) float64 {
	v, _ := d.As(unit)
	return v
}

// AsMilliseconds is duration#asMilliseconds (also its valueOf).
func (d *Duration) AsMilliseconds() float64 { return d.as("ms") }

// Get is duration#get(units).
func (d *Duration) Get(units Value) (Value, error) {
	unit := js.ToString(NormalizeUnits(units))
	if !d.valid {
		return js.NaN, nil
	}
	switch unit + "s" {
	case "milliseconds":
		return d.data.milliseconds, nil
	case "seconds":
		return d.data.seconds, nil
	case "minutes":
		return d.data.minutes, nil
	case "hours":
		return d.data.hours, nil
	case "days":
		return d.data.days, nil
	case "weeks":
		return d.Weeks(), nil
	case "months":
		return d.data.months, nil
	case "years":
		return d.data.years, nil
	}
	return nil, typeError("this[(units + \"s\")] is not a function")
}

func (d *Duration) field(v float64) float64 {
	if !d.valid {
		return js.NaN
	}
	return v
}

// Milliseconds is duration#milliseconds; the unit getters below read the bubbled values.
func (d *Duration) Milliseconds() float64 { return d.field(d.data.milliseconds) }

// Seconds is duration#seconds.
func (d *Duration) Seconds() float64 { return d.field(d.data.seconds) }

// Minutes is duration#minutes.
func (d *Duration) Minutes() float64 { return d.field(d.data.minutes) }

// Hours is duration#hours.
func (d *Duration) Hours() float64 { return d.field(d.data.hours) }

// Days is duration#days.
func (d *Duration) Days() float64 { return d.field(d.data.days) }

// Weeks is duration#weeks.
func (d *Duration) Weeks() float64 { return absFloor(d.Days() / 7) }

// Months is duration#months.
func (d *Duration) Months() float64 { return d.field(d.data.months) }

// Years is duration#years.
func (d *Duration) Years() float64 { return d.field(d.data.years) }

// Clone is duration#clone.
func (d *Duration) Clone() *Duration { return CreateDuration(d, js.Undef) }

// Locale is duration#locale() as a getter.
func (d *Duration) Locale() Value { return d.locale.Abbr() }

// SetLocale is duration#locale(key).
func (d *Duration) SetLocale(key Value) *Duration {
	if l := getLocale(key); l != nil {
		d.locale = l
	}
	return d
}

// LocaleData is duration#localeData.
func (d *Duration) LocaleData() *Locale { return d.locale }

func sign(x float64) float64 {
	s := 0.0
	if x > 0 {
		s = 1
	} else if x < 0 {
		s = -1
	}
	if s != 0 {
		return s
	}
	return x
}

// ToISOString is duration#toISOString: an ISO 8601 duration that keeps days and months apart.
func (d *Duration) ToISOString() string {
	if !d.valid {
		return d.LocaleData().InvalidDate()
	}
	seconds := math.Abs(d.milliseconds) / 1000
	days := math.Abs(d.days)
	months := math.Abs(d.months)
	total := d.as("s")
	if total == 0 || math.IsNaN(total) {
		return "P0D"
	}
	minutes := absFloor(seconds / 60)
	hours := absFloor(minutes / 60)
	seconds = math.Mod(seconds, 60)
	minutes = math.Mod(minutes, 60)
	years := absFloor(months / 12)
	months = math.Mod(months, 12)
	s := ""
	if seconds != 0 && !math.IsNaN(seconds) {
		s = isoTrailingZeros.Replace(js.ToFixed(seconds, 3), "")
	}
	pick := func(cond bool, v string) string {
		if cond {
			return v
		}
		return ""
	}
	totalSign := pick(total < 0, "-")
	ymSign := pick(sign(d.months) != sign(total), "-")
	daysSign := pick(sign(d.days) != sign(total), "-")
	hmsSign := pick(sign(d.milliseconds) != sign(total), "-")
	nonzero := func(v float64) bool { return v != 0 && !math.IsNaN(v) }
	var b strings.Builder
	b.WriteString(totalSign + "P")
	b.WriteString(pick(nonzero(years), ymSign+js.NumberToString(years)+"Y"))
	b.WriteString(pick(nonzero(months), ymSign+js.NumberToString(months)+"M"))
	b.WriteString(pick(nonzero(days), daysSign+js.NumberToString(days)+"D"))
	b.WriteString(pick(nonzero(hours) || nonzero(minutes) || nonzero(seconds), "T"))
	b.WriteString(pick(nonzero(hours), hmsSign+js.NumberToString(hours)+"H"))
	b.WriteString(pick(nonzero(minutes), hmsSign+js.NumberToString(minutes)+"M"))
	b.WriteString(pick(nonzero(seconds), hmsSign+s+"S"))
	return b.String()
}

var (
	thresholds = js.Obj("ss", 44.0, "s", 45.0, "m", 45.0, "h", 22.0, "d", 26.0, "w", nil, "M", 11.0)
	// roundingFunction is moment.relativeTimeRounding's current function; nil means Math.round.
	roundingFunction *js.Func
	mathRound        = &js.Func{Name: "round", Call: func(_ any, args []any) any {
		return js.Round(js.ToNumber(arg(args, 0)))
	}}
)

// RelativeTimeRounding is moment.relativeTimeRounding(fn): the current function without an argument, true when
// fn is a function and was installed, false otherwise.
func RelativeTimeRounding(fn Value) Value {
	if js.IsUndefined(fn) {
		if roundingFunction == nil {
			return mathRound
		}
		return roundingFunction
	}
	if f, ok := fn.(*js.Func); ok {
		roundingFunction = f
		return true
	}
	return false
}

// RelativeTimeThreshold is moment.relativeTimeThreshold(threshold, limit).
func RelativeTimeThreshold(threshold, limit Value) Value {
	key := js.ToString(threshold)
	if !thresholds.Has(key) {
		return false
	}
	if js.IsUndefined(limit) {
		return thresholds.Get(key)
	}
	thresholds.Set(key, limit)
	if key == "s" {
		thresholds.Set("ss", js.ToNumber(limit)-1)
	}
	return true
}

func round(x float64) float64 {
	if roundingFunction == nil {
		return js.Round(x)
	}
	return js.ToNumber(roundingFunction.Call(js.Undef, []any{x}))
}

func relativeTime(posNeg *Duration, withoutSuffix bool, th *js.Object, locale *Locale) string {
	d := CreateDuration(posNeg, js.Undef).Abs()
	seconds, minutes, hours := round(d.as("s")), round(d.as("m")), round(d.as("h"))
	days, months, weeks, years := round(d.as("d")), round(d.as("M")), round(d.as("w")), round(d.as("y"))
	t := func(k string) float64 { return js.ToNumber(th.Get(k)) }
	key, number := "", js.NaN
	switch {
	case seconds <= t("ss"):
		key, number = "s", seconds
	case seconds < t("s"):
		key, number = "ss", seconds
	case minutes <= 1:
		key = "m"
	case minutes < t("m"):
		key, number = "mm", minutes
	case hours <= 1:
		key = "h"
	case hours < t("h"):
		key, number = "hh", hours
	case days <= 1:
		key = "d"
	case days < t("d"):
		key, number = "dd", days
	}
	if key == "" && !js.IsNullish(th.Get("w")) {
		switch {
		case weeks <= 1:
			key = "w"
		case weeks < t("w"):
			key, number = "ww", weeks
		}
	}
	if key == "" {
		switch {
		case months <= 1:
			key = "M"
		case months < t("M"):
			key, number = "MM", months
		case years <= 1:
			key = "y"
		default:
			key, number = "yy", years
		}
	}
	if number == 0 || math.IsNaN(number) {
		number = 1
	}
	return locale.relativeTimeWithoutPostformat(number, withoutSuffix, key, posNeg.AsMilliseconds() > 0)
}

// Humanize is duration#humanize(withSuffix, thresholds).
func (d *Duration) Humanize(argWithSuffix, argThresholds Value) string {
	if !d.valid {
		return d.LocaleData().InvalidDate()
	}
	withSuffix := false
	if isObjectArg(argWithSuffix) {
		argThresholds = argWithSuffix
		argWithSuffix = false
	}
	if b, ok := argWithSuffix.(bool); ok {
		withSuffix = b
	}
	th := thresholds
	if isObjectArg(argThresholds) {
		th = thresholds.Clone()
		if o, ok := argThresholds.(*js.Object); ok {
			th.Extend(o)
			if !js.IsNullish(o.Get("s")) && js.IsNullish(o.Get("ss")) {
				th.Set("ss", js.ToNumber(o.Get("s"))-1)
			}
		}
	}
	locale := d.LocaleData()
	output := relativeTime(d, !withSuffix, th, locale)
	if withSuffix {
		output = locale.pastFutureWithoutPostformat(d.AsMilliseconds(), output)
	}
	return locale.Postformat(output)
}

func isObjectArg(v Value) bool {
	switch v.(type) {
	case *js.Object, []any, []string, *Moment, *Duration, *Locale, *js.Date:
		return true
	}
	return false
}
