package js

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

// NaN is the JavaScript NaN value.
var NaN = math.NaN()

// PrimitiveHook converts the objects of the embedding library (moments, durations, locales) to a primitive the
// way their valueOf and toString methods do; it reports false for values it does not own.
var PrimitiveHook func(v any, hintString bool) (any, bool)

// ToPrimitive is the abstract operation ToPrimitive with hint number (or string when hintString is set).
func ToPrimitive(v any, hintString bool) any {
	switch x := v.(type) {
	case nil, Undefined, bool, float64, string:
		return v
	case *Date:
		if hintString {
			return x.String()
		}
		return x.t
	case []any, []string:
		return ToString(v)
	case *Object:
		return "[object Object]"
	case *Func:
		return "function " + x.Name + "() { [native code] }"
	}
	if PrimitiveHook != nil {
		if p, ok := PrimitiveHook(v, hintString); ok {
			return p
		}
	}
	return "[object Object]"
}

// ToNumber is the unary plus operator.
func ToNumber(v any) float64 {
	switch x := ToPrimitive(v, false).(type) {
	case nil:
		return 0
	case Undefined:
		return NaN
	case bool:
		if x {
			return 1
		}
		return 0
	case float64:
		return x
	case string:
		return StringToNumber(x)
	}
	return NaN
}

// IsNaNValue is the global isNaN function: ToNumber followed by the NaN test.
func IsNaNValue(v any) bool {
	return math.IsNaN(ToNumber(v))
}

// IsFinite is the global isFinite function.
func IsFinite(v any) bool {
	n := ToNumber(v)
	return !math.IsNaN(n) && !math.IsInf(n, 0)
}

// StringToNumber converts a string as Number(string) does.
func StringToNumber(s string) float64 {
	s = TrimSpace(s)
	if s == "" {
		return 0
	}
	if len(s) > 2 && s[0] == '0' {
		base := 0
		switch s[1] {
		case 'x', 'X':
			base = 16
		case 'o', 'O':
			base = 8
		case 'b', 'B':
			base = 2
		}
		if base != 0 {
			n, ok := new(big.Int).SetString(s[2:], base)
			if !ok || strings.ContainsAny(s[2:], "_+-") {
				return NaN
			}
			f, _ := new(big.Float).SetInt(n).Float64()
			return f
		}
	}
	n := decimalPrefix(s)
	if n != len(s) {
		return NaN
	}
	return parseDecimal(s)
}

// decimalPrefix returns the length of the longest prefix of s that is a StrDecimalLiteral, or 0.
func decimalPrefix(s string) int {
	i := 0
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		i++
	}
	if strings.HasPrefix(s[i:], "Infinity") {
		return i + len("Infinity")
	}
	digits := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
		digits++
	}
	if i < len(s) && s[i] == '.' {
		j := i + 1
		frac := 0
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			j++
			frac++
		}
		if digits+frac > 0 {
			i = j
			digits += frac
		}
	}
	if digits == 0 {
		return 0
	}
	if i < len(s) && (s[i] == 'e' || s[i] == 'E') {
		j := i + 1
		if j < len(s) && (s[j] == '+' || s[j] == '-') {
			j++
		}
		k := j
		for k < len(s) && s[k] >= '0' && s[k] <= '9' {
			k++
		}
		if k > j {
			i = k
		}
	}
	return i
}

func parseDecimal(s string) float64 {
	switch s {
	case "Infinity", "+Infinity":
		return math.Inf(1)
	case "-Infinity":
		return math.Inf(-1)
	}
	// s is a validated decimal literal, so a range error only reports overflow to ±Inf or underflow to ±0.
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// ParseFloat is the global parseFloat function.
func ParseFloat(v any) float64 {
	s := TrimLeftSpace(ToString(v))
	n := decimalPrefix(s)
	if n == 0 {
		return NaN
	}
	return parseDecimal(s[:n])
}

// ParseInt is the global parseInt function with radix 10.
func ParseInt(v any) float64 {
	s := TrimLeftSpace(ToString(v))
	neg := false
	if s != "" && (s[0] == '+' || s[0] == '-') {
		neg = s[0] == '-'
		s = s[1:]
	}
	i := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		i++
	}
	if i == 0 {
		return NaN
	}
	f, _ := strconv.ParseFloat(s[:i], 64)
	if neg {
		return -f
	}
	return f
}

// NumberToString is Number.prototype.toString() in radix 10.
func NumberToString(x float64) string {
	switch {
	case math.IsNaN(x):
		return "NaN"
	case x == 0:
		return "0"
	case math.IsInf(x, 1):
		return "Infinity"
	case math.IsInf(x, -1):
		return "-Infinity"
	case x < 0:
		return "-" + NumberToString(-x)
	}
	e := strconv.FormatFloat(x, 'e', -1, 64)
	mant, exp, _ := strings.Cut(e, "e")
	digits := strings.Replace(mant, ".", "", 1)
	k := len(digits)
	ex, _ := strconv.Atoi(exp)
	n := ex + 1
	switch {
	case k <= n && n <= 21:
		return digits + strings.Repeat("0", n-k)
	case 0 < n && n <= 21:
		return digits[:n] + "." + digits[n:]
	case -6 < n && n <= 0:
		return "0." + strings.Repeat("0", -n) + digits
	}
	sign := "+"
	if n-1 < 0 {
		sign = "-"
	}
	exponent := strconv.Itoa(absInt(n - 1))
	if k == 1 {
		return digits + "e" + sign + exponent
	}
	return digits[:1] + "." + digits[1:] + "e" + sign + exponent
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// ToFixed is Number.prototype.toFixed(digits): the exact value is rounded to digits decimals, ties away from
// zero.
func ToFixed(x float64, digits int) string {
	if math.IsNaN(x) {
		return "NaN"
	}
	if math.Abs(x) >= 1e21 {
		return NumberToString(x)
	}
	sign := ""
	if x < 0 {
		sign = "-"
		x = -x
	}
	exact := new(big.Float).SetPrec(2048).SetFloat64(x)
	scaled := new(big.Float).SetPrec(2048).Mul(exact, new(big.Float).SetPrec(2048).SetInt(
		new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(digits)), nil)))
	whole, _ := scaled.Int(nil)
	rest := new(big.Float).SetPrec(2048).Sub(scaled, new(big.Float).SetPrec(2048).SetInt(whole))
	if rest.Cmp(big.NewFloat(0.5)) >= 0 {
		whole.Add(whole, big.NewInt(1))
	}
	s := whole.String()
	if digits > 0 {
		if len(s) <= digits {
			s = strings.Repeat("0", digits-len(s)+1) + s
		}
		s = s[:len(s)-digits] + "." + s[len(s)-digits:]
	}
	if sign != "" && strings.Trim(s, "0.") == "" {
		return "-" + s
	}
	return sign + s
}

// Round is Math.round: halves round towards +Infinity and the sign of zero is kept.
func Round(x float64) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) || x == math.Trunc(x) {
		return x
	}
	r := math.Floor(x)
	if x-r >= 0.5 {
		r++
	}
	if r == 0 && x < 0 {
		return math.Copysign(0, -1)
	}
	return r
}

// ToInt32 is the ToInt32 conversion behind JavaScript's bitwise operators (so ~~x is ToInt32(x)).
func ToInt32(x float64) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return 0
	}
	t := math.Mod(math.Trunc(x), 4294967296)
	if t < 0 {
		t += 4294967296
	}
	if t >= 2147483648 {
		t -= 4294967296
	}
	return t + 0
}

// ToIntegerOrInfinity truncates towards zero; NaN becomes 0.
func ToIntegerOrInfinity(x float64) float64 {
	if math.IsNaN(x) {
		return 0
	}
	return math.Trunc(x) + 0
}

// Mod is the % operator.
func Mod(a, b float64) float64 {
	return math.Mod(a, b)
}

// Add is the binary + operator on two values: string concatenation when either primitive is a string.
func Add(a, b any) any {
	pa, pb := ToPrimitive(a, false), ToPrimitive(b, false)
	_, sa := pa.(string)
	_, sb := pb.(string)
	if sa || sb {
		return ToString(pa) + ToString(pb)
	}
	return ToNumber(pa) + ToNumber(pb)
}
