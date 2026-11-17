// Package locale holds the 139 locales of moment. Each file defines one locale the way a file under locale/ of
// moment.js does: init registers it, and moment loads it the first time its name is used.
package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

// mapMatches replaces every match of re in s with its entry in table; a missing entry is the text "undefined", as
// a JavaScript lookup of a missing key would be.
func mapMatches(re *jsre.Regexp, s string, table map[string]string) string {
	return re.ReplaceFunc(s, func(m *jsre.Match) string {
		if v, ok := table[m.Groups[0]]; ok {
			return v
		}
		return "undefined"
	})
}

// digitTables returns the table from ASCII digits to a script's digits and its inverse.
func digitTables(digits string) (symbols, numbers map[string]string) {
	symbols, numbers = map[string]string{}, map[string]string{}
	for i, d := range []rune(digits) {
		ascii := string(rune('0' + i))
		symbols[ascii] = string(d)
		numbers[string(d)] = ascii
	}
	return symbols, numbers
}

var (
	asciiDigit  = jsre.MustCompile(`\d`, "g")
	comma       = jsre.MustCompile(`,`, "g")
	arabicComma = jsre.MustCompile(`،`, "g")
	percentD    = jsre.MustCompile(`%d`, "i")
)

// num is a number as JavaScript concatenates it into a string.
func num(n float64) string { return js.NumberToString(n) }

// str is any value as JavaScript concatenates it into a string; ordinal callbacks receive their argument as it was
// passed, a number or a string.
func str(v moment.Value) string { return js.ToString(v) }

// sameNumber reports whether v === n: v is a number (not a string holding one) equal to n.
func sameNumber(v moment.Value, n float64) bool {
	f, ok := v.(float64)
	return ok && f == n
}

// digitSuffix is the suffix of n's last digit, its tens or its hundreds in a Turkic suffix table, or undefined when
// none applies (a JavaScript lookup of a missing key).
func digitSuffix(suffixes map[string]string, n float64) moment.Value {
	a := js.Mod(n, 10)
	c := "null"
	if n >= 100 {
		c = "100"
	}
	for _, key := range []string{num(a), num(js.Mod(n, 100) - a), c} {
		if s := suffixes[key]; s != "" {
			return s
		}
	}
	return js.Undef
}

// replaceCount substitutes the first %d (in any case) of format with the number.
func replaceCount(format string, n float64) string { return percentD.Replace(format, num(n)) }

// slavicPlural picks forms[0], forms[1] or forms[2] of a word list by the East Slavic plural rule.
func slavicPlural(forms []string, n float64) string {
	switch {
	case js.Mod(n, 10) == 1 && js.Mod(n, 100) != 11:
		return forms[0]
	case js.Mod(n, 10) >= 2 && js.Mod(n, 10) <= 4 && (js.Mod(n, 100) < 10 || js.Mod(n, 100) >= 20):
		return forms[1]
	}
	return forms[2]
}

// arabicPluralForm is the six-way Arabic plural category of n.
func arabicPluralForm(n float64) int {
	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	case n == 2:
		return 2
	case js.Mod(n, 100) >= 3 && js.Mod(n, 100) <= 10:
		return 3
	case js.Mod(n, 100) >= 11:
		return 4
	}
	return 5
}

type arabicForms struct {
	forms [6]string
	dual  [2]string
}

var arabicPlurals = map[string]arabicForms{
	"s": {[6]string{"أقل من ثانية", "ثانية واحدة", "", "%d ثوان", "%d ثانية", "%d ثانية"}, [2]string{"ثانيتان", "ثانيتين"}},
	"m": {[6]string{"أقل من دقيقة", "دقيقة واحدة", "", "%d دقائق", "%d دقيقة", "%d دقيقة"}, [2]string{"دقيقتان", "دقيقتين"}},
	"h": {[6]string{"أقل من ساعة", "ساعة واحدة", "", "%d ساعات", "%d ساعة", "%d ساعة"}, [2]string{"ساعتان", "ساعتين"}},
	"d": {[6]string{"أقل من يوم", "يوم واحد", "", "%d أيام", "%d يومًا", "%d يوم"}, [2]string{"يومان", "يومين"}},
	"M": {[6]string{"أقل من شهر", "شهر واحد", "", "%d أشهر", "%d شهرا", "%d شهر"}, [2]string{"شهران", "شهرين"}},
	"y": {[6]string{"أقل من عام", "عام واحد", "", "%d أعوام", "%d عامًا", "%d عام"}, [2]string{"عامان", "عامين"}},
}

// arabicPluralize returns the relative-time callback of the Arabic locales for unit u.
func arabicPluralize(u string) func(float64, bool, string, bool) string {
	return func(number float64, withoutSuffix bool, _ string, _ bool) string {
		forms := arabicPlurals[u]
		f := arabicPluralForm(number)
		str := forms.forms[f]
		if f == 2 {
			str = forms.dual[1]
			if withoutSuffix {
				str = forms.dual[0]
			}
		}
		return replaceCount(str, number)
	}
}

// arabicMeridiem is the meridiem of the Arabic locales.
func arabicMeridiem(_ *moment.Locale, hours, _ float64, _ bool) string {
	if hours < 12 {
		return "ص"
	}
	return "م"
}

// arabicIsPM is the isPM of the Arabic locales.
func arabicIsPM(_ *moment.Locale, input string) bool { return input == "م" }

var arabicSymbols, arabicNumbers = digitTables("٠١٢٣٤٥٦٧٨٩")

// arabicDigitsPostformat writes Arabic-Indic digits and the Arabic comma.
func arabicDigitsPostformat(_ *moment.Locale, s string) string {
	return comma.Replace(mapMatches(asciiDigit, s, arabicSymbols), "،")
}

var arabicDigit = jsre.MustCompile(`[١٢٣٤٥٦٧٨٩٠]`, "g")

// arabicDigitsPreparse reads Arabic-Indic digits and the Arabic comma.
func arabicDigitsPreparse(_ *moment.Locale, s string) string {
	return arabicComma.Replace(mapMatches(arabicDigit, s, arabicNumbers), ",")
}

// reverseUnits reverses a string by UTF-16 code units, as split(”).reverse().join(”) does.
func reverseUnits(s string) string {
	u := js.UTF16(s)
	for i, j := 0, len(u)-1; i < j; i, j = i+1, j-1 {
		u[i], u[j] = u[j], u[i]
	}
	return js.FromUTF16(u)
}

// noonToZero is the `if (hour === 12) hour = 0` that starts most meridiemHour callbacks.
func noonToZero(hour any) any {
	if js.StrictEquals(hour, 12.0) {
		return 0.0
	}
	return hour
}

// hourNumber is the hour as a number in comparisons; undefined compares as NaN.
func hourNumber(hour any) float64 { return js.ToNumber(hour) }

// plus12 is `hour + 12`, NaN when the hour is undefined.
func plus12(hour any) any { return js.ToNumber(hour) + 12 }

var dashMonthFormat = jsre.MustCompile(`-MMM-`, "")

// pickDottedMonth is the monthsShort of the Dutch and Frisian locales: the dotless abbreviation in a -MMM-
// format, the dotted one otherwise, and the dotted list without a moment.
func pickDottedMonth(m *moment.Moment, format any, withDots, withoutDots []string) any {
	if m == nil {
		return withDots
	}
	list := withDots
	if dashMonthFormat.Test(js.ToString(format)) {
		list = withoutDots
	}
	if i := m.Month(); i >= 0 && int(i) < len(list) {
		return list[int(i)]
	}
	return js.Undef
}

// dayPartHour is the meridiemHour of the Indic locales: night hours before 4 stay, morning hours stay, midday
// hours from the given hour stay, evening hours move by twelve.
func dayPartHour(hour any, meridiem, night, morning, midday, evening string, middayFrom float64) any {
	hour = noonToZero(hour)
	switch meridiem {
	case night:
		if hourNumber(hour) < 4 {
			return hour
		}
		return plus12(hour)
	case morning:
		return hour
	case midday:
		if hourNumber(hour) >= middayFrom {
			return hour
		}
		return plus12(hour)
	case evening:
		return plus12(hour)
	}
	return js.Undef
}
