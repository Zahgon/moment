package moment

import (
	"math"

	"github.com/moment/moment/internal/js"
)

// Moment is a moment object. The same structure carries the creation config while a moment is parsed, as the
// plain object of moment.js does; fields that JavaScript can leave undefined hold js.Undef.
type Moment struct {
	isAMomentObject Value
	i, f, l, strict Value
	tzm, isUTC      Value
	offset          Value
	pf              *js.Object
	locale          *Locale
	d               *js.Date
	isValid         Value

	useUTC              Value
	a                   []Value
	w                   *js.Object
	dayOfYear           Value
	meridiem            Value
	isPm                Value
	nextDay             Value
	defaultDatePartsNow *nowCache
	isDSTShifted        Value
}

type nowCache struct {
	has   bool
	value float64
}

func newConfig() *Moment {
	u := js.Undef
	return &Moment{isAMomentObject: u, i: u, f: u, l: u, strict: u, tzm: u, isUTC: u, offset: u, isValid: u,
		useUTC: u, dayOfYear: u, meridiem: u, isPm: u, nextDay: u, isDSTShifted: u}
}

// copyConfig copies the moment properties of from onto to.
func copyConfig(to, from *Moment) *Moment {
	if !js.IsUndefined(from.isAMomentObject) {
		to.isAMomentObject = from.isAMomentObject
	}
	if !js.IsUndefined(from.i) {
		to.i = from.i
	}
	if !js.IsUndefined(from.f) {
		to.f = from.f
	}
	if !js.IsUndefined(from.l) {
		to.l = from.l
	}
	if !js.IsUndefined(from.strict) {
		to.strict = from.strict
	}
	if !js.IsUndefined(from.tzm) {
		to.tzm = from.tzm
	}
	if !js.IsUndefined(from.isUTC) {
		to.isUTC = from.isUTC
	}
	if !js.IsUndefined(from.offset) {
		to.offset = from.offset
	}
	if from.pf != nil {
		to.pf = getParsingFlags(from)
	}
	if from.locale != nil {
		to.locale = from.locale
	}
	return to
}

// newMoment is the Moment constructor: it copies config and owns a copy of its date.
func newMoment(config *Moment) *Moment {
	m := copyConfig(newConfig(), config)
	t := js.NaN
	if config.d != nil {
		t = config.d.Time()
	}
	m.d = js.NewDate(t)
	if !m.IsValid() {
		m.d = js.NewDate(js.NaN)
	}
	return m
}

// IsMoment is moment.isMoment.
func IsMoment(v Value) bool {
	_, ok := v.(*Moment)
	return ok
}

func defaultParsingFlags() *js.Object {
	return js.Obj("empty", false, "unusedTokens", []any{}, "unusedInput", []any{}, "overflow", -2.0,
		"charsLeftOver", 0.0, "nullInput", false, "invalidEra", nil, "invalidMonth", nil, "invalidOffset", nil,
		"invalidFormat", false, "userInvalidated", false, "iso", false, "parsedDateParts", []any{}, "era", nil,
		"meridiem", nil, "rfc2822", false, "weekdayMismatch", false)
}

func getParsingFlags(m *Moment) *js.Object {
	if m.pf == nil {
		m.pf = defaultParsingFlags()
	}
	return m.pf
}

func pushFlag(pf *js.Object, key string, v Value) {
	list, _ := pf.Get(key).([]any)
	pf.Set(key, append(append([]any(nil), list...), v))
}

func isValidConfig(m *Moment) bool {
	valid := m.d != nil && m.d.Valid()
	if valid {
		flags := getParsingFlags(m)
		parsedParts := false
		if parts, ok := flags.Get("parsedDateParts").([]any); ok {
			for _, p := range parts {
				if !js.IsNullish(p) {
					parsedParts = true
					break
				}
			}
		}
		meridiem := js.Truthy(flags.Get("meridiem"))
		valid = js.ToNumber(flags.Get("overflow")) < 0 &&
			!js.Truthy(flags.Get("empty")) &&
			!js.Truthy(flags.Get("invalidEra")) &&
			!js.Truthy(flags.Get("invalidMonth")) &&
			!js.Truthy(flags.Get("invalidOffset")) &&
			!js.Truthy(flags.Get("invalidWeekday")) &&
			!js.Truthy(flags.Get("weekdayMismatch")) &&
			!js.Truthy(flags.Get("nullInput")) &&
			!js.Truthy(flags.Get("invalidFormat")) &&
			!js.Truthy(flags.Get("userInvalidated")) &&
			(!meridiem || parsedParts)
		if js.Truthy(m.strict) {
			unused, _ := flags.Get("unusedTokens").([]any)
			valid = valid && js.StrictEquals(flags.Get("charsLeftOver"), 0.0) && len(unused) == 0 &&
				js.IsUndefined(flags.Get("bigHour"))
		}
	}
	m.isValid = valid
	return valid
}

// IsValid is moment#isValid.
func (m *Moment) IsValid() bool {
	return isValidConfig(m)
}

// CreateInvalid is moment.invalid(flags): an invalid moment whose parsing flags are extended with flags.
func CreateInvalid(flags Value) *Moment {
	m := utcMoment(js.NaN)
	if o, ok := flags.(*js.Object); ok {
		getParsingFlags(m).Extend(o)
	} else if !js.IsNullish(flags) {
		extendWith(getParsingFlags(m), flags)
	} else {
		getParsingFlags(m).Set("userInvalidated", true)
	}
	return m
}

// extendWith copies the own enumerable properties of a non-plain value, as extend() does for arrays and strings.
func extendWith(o *js.Object, v Value) {
	switch x := v.(type) {
	case []any:
		for i, e := range x {
			o.Set(js.NumberToString(float64(i)), e)
		}
	case string:
		for i, c := range js.Split(x, "") {
			o.Set(js.NumberToString(float64(i)), c)
		}
	}
}

// localMoment is moment(input, format, locale, strict): a moment in local time. Where the original throws while
// parsing it panics with a *ThrowError, which the exported entry points return as an error.
func localMoment(args ...Value) *Moment {
	return createLocalOrUTC(arg(args, 0), arg(args, 1), arg(args, 2), arg(args, 3), false)
}

// utcMoment is moment.utc(input, format, locale, strict); like localMoment it panics with a *ThrowError where the
// original throws.
func utcMoment(args ...Value) *Moment {
	return createLocalOrUTC(arg(args, 0), arg(args, 1), arg(args, 2), arg(args, 3), true).UTC(js.Undef)
}

// unixMoment is moment.unix(seconds).
func unixMoment(input Value) *Moment {
	return localMoment(js.ToNumber(input) * 1000)
}

// parseZoneMoment is moment.parseZone(input, format, locale, strict).
func parseZoneMoment(args ...Value) *Moment {
	return localMoment(args...).ParseZone()
}

func arg(args []Value, i int) Value {
	if i < len(args) {
		return args[i]
	}
	return js.Undef
}

func isObject(v Value) bool {
	switch v.(type) {
	case *js.Object, *Moment, *Duration, *Locale:
		return true
	}
	return false
}

func isArray(v Value) bool {
	switch v.(type) {
	case []any, []string:
		return true
	}
	return false
}

func asArray(v Value) []any {
	switch x := v.(type) {
	case []any:
		return x
	case []string:
		out := make([]any, len(x))
		for i, s := range x {
			out[i] = s
		}
		return out
	}
	return nil
}

func createLocalOrUTC(input, format, locale, strict Value, isUTC bool) *Moment {
	if b, ok := format.(bool); ok {
		strict = b
		format = js.Undef
	}
	if b, ok := locale.(bool); ok {
		strict = b
		locale = js.Undef
	}
	if o, ok := input.(*js.Object); ok && o.Len() == 0 {
		input = js.Undef
	}
	if isArray(input) && len(asArray(input)) == 0 {
		input = js.Undef
	}
	c := newConfig()
	c.isAMomentObject = true
	c.useUTC = isUTC
	c.isUTC = isUTC
	c.l = locale
	c.i = input
	c.f = format
	c.strict = strict
	return createFromConfig(c)
}

func createFromConfig(config *Moment) *Moment {
	res := newMoment(checkOverflow(prepareConfig(config)))
	if js.Truthy(res.nextDay) {
		res.Add(1.0, "d")
		res.nextDay = js.Undef
	}
	return res
}

func prepareConfig(config *Moment) *Moment {
	input := config.i
	format := config.f
	if config.locale == nil {
		config.locale = getLocale(config.l)
	}
	if input == nil || (js.IsUndefined(format) && js.StrictEquals(input, "")) {
		return CreateInvalid(js.Obj("nullInput", true))
	}
	if s, ok := input.(string); ok {
		input = config.locale.Preparse(s)
		config.i = input
	}
	switch {
	case IsMoment(input):
		return newMoment(checkOverflow(input.(*Moment)))
	case isDate(input):
		config.d = input.(*js.Date)
	case isArray(format):
		configFromStringAndArray(config)
	case js.Truthy(format):
		configFromStringAndFormat(config)
	default:
		configFromInput(config)
	}
	if !isValidConfig(config) {
		config.d = nil
	}
	return config
}

func isDate(v Value) bool {
	_, ok := v.(*js.Date)
	return ok
}

func configFromInput(config *Moment) {
	switch input := config.i.(type) {
	case js.Undefined:
		config.d = js.NewDate(Now())
	case *js.Date:
		config.d = js.NewDate(input.Time())
	case string:
		configFromString(config)
	case []any, []string:
		list := asArray(input)
		config.a = make([]Value, len(list))
		for i, v := range list {
			config.a[i] = js.ParseInt(v)
		}
		configFromArray(config)
	case *js.Object, *Moment, *Duration, *Locale:
		configFromObject(config)
	case float64:
		config.d = js.NewDate(input)
	default:
		createFromInputFallback(config)
	}
}

// toInt is moment's toInt: the number truncated towards zero, 0 for NaN, ±Infinity and -0.
func toInt(v Value) float64 {
	n := js.ToNumber(v)
	if n != 0 && !math.IsNaN(n) && !math.IsInf(n, 0) {
		return absFloor(n)
	}
	return 0
}

func absFloor(n float64) float64 {
	if n < 0 {
		return math.Ceil(n) + 0
	}
	return math.Floor(n)
}

func absCeil(n float64) float64 {
	if n < 0 {
		return math.Floor(n)
	}
	return math.Ceil(n)
}

func absRound(n float64) float64 {
	if n < 0 {
		return js.Round(-1*n) * -1
	}
	return js.Round(n)
}

func mod(n, x float64) float64 {
	return math.Mod(math.Mod(n, x)+x, x)
}

func isLeapYear(year float64) bool {
	return (math.Mod(year, 4) == 0 && math.Mod(year, 100) != 0) || math.Mod(year, 400) == 0
}

func daysInYear(year float64) float64 {
	if isLeapYear(year) {
		return 366
	}
	return 365
}

func daysInMonth(year, month float64) float64 {
	if math.IsNaN(year) || math.IsNaN(month) {
		return js.NaN
	}
	modMonth := mod(month, 12)
	year += (month - modMonth) / 12
	if modMonth == 1 {
		if isLeapYear(year) {
			return 29
		}
		return 28
	}
	return 31 - math.Mod(math.Mod(modMonth, 7), 2)
}
