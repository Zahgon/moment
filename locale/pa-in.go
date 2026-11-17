// moment.js locale configuration
// locale : Punjabi (India) [pa-in]
// author : Harpreet Singh : https://github.com/harpreetkhalsagtbit

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("pa-in", func() {
		moment.DefineLocale("pa-in", js.Obj(
			"months", []string{"ਜਨਵਰੀ", "ਫ਼ਰਵਰੀ", "ਮਾਰਚ", "ਅਪ੍ਰੈਲ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾਈ", "ਅਗਸਤ", "ਸਤੰਬਰ", "ਅਕਤੂਬਰ", "ਨਵੰਬਰ", "ਦਸੰਬਰ"},
			"monthsShort", []string{"ਜਨਵਰੀ", "ਫ਼ਰਵਰੀ", "ਮਾਰਚ", "ਅਪ੍ਰੈਲ", "ਮਈ", "ਜੂਨ", "ਜੁਲਾਈ", "ਅਗਸਤ", "ਸਤੰਬਰ", "ਅਕਤੂਬਰ", "ਨਵੰਬਰ", "ਦਸੰਬਰ"},
			"weekdays", []string{"ਐਤਵਾਰ", "ਸੋਮਵਾਰ", "ਮੰਗਲਵਾਰ", "ਬੁਧਵਾਰ", "ਵੀਰਵਾਰ", "ਸ਼ੁੱਕਰਵਾਰ", "ਸ਼ਨੀਚਰਵਾਰ"},
			"weekdaysShort", []string{"ਐਤ", "ਸੋਮ", "ਮੰਗਲ", "ਬੁਧ", "ਵੀਰ", "ਸ਼ੁਕਰ", "ਸ਼ਨੀ"},
			"weekdaysMin", []string{"ਐਤ", "ਸੋਮ", "ਮੰਗਲ", "ਬੁਧ", "ਵੀਰ", "ਸ਼ੁਕਰ", "ਸ਼ਨੀ"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm ਵਜੇ",
				"LTS", "A h:mm:ss ਵਜੇ",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm ਵਜੇ",
				"LLLL", "dddd, D MMMM YYYY, A h:mm ਵਜੇ",
			),
			"calendar", js.Obj(
				"sameDay", "[ਅਜ] LT",
				"nextDay", "[ਕਲ] LT",
				"nextWeek", "[ਅਗਲਾ] dddd, LT",
				"lastDay", "[ਕਲ] LT",
				"lastWeek", "[ਪਿਛਲੇ] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s ਵਿੱਚ",
				"past", "%s ਪਿਛਲੇ",
				"s", "ਕੁਝ ਸਕਿੰਟ",
				"ss", "%d ਸਕਿੰਟ",
				"m", "ਇਕ ਮਿੰਟ",
				"mm", "%d ਮਿੰਟ",
				"h", "ਇੱਕ ਘੰਟਾ",
				"hh", "%d ਘੰਟੇ",
				"d", "ਇੱਕ ਦਿਨ",
				"dd", "%d ਦਿਨ",
				"M", "ਇੱਕ ਮਹੀਨਾ",
				"MM", "%d ਮਹੀਨੇ",
				"y", "ਇੱਕ ਸਾਲ",
				"yy", "%d ਸਾਲ",
			),
			"preparse", paInPreparse,
			"postformat", paInPostformat,
			"meridiemParse", jsre.MustCompile("ਰਾਤ|ਸਵੇਰ|ਦੁਪਹਿਰ|ਸ਼ਾਮ", ""),
			"meridiemHour", paInMeridiemHour,
			"meridiem", paInMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	paInSymbols, paInNumbers = digitTables("੦੧੨੩੪੫੬੭੮੯")
	paInDigit                = jsre.MustCompile(`[੧੨੩੪੫੬੭੮੯੦]`, "g")
)

// paInPreparse reads Gurmukhi digits.
func paInPreparse(_ *moment.Locale, s string) string { return mapMatches(paInDigit, s, paInNumbers) }

// paInPostformat writes Gurmukhi digits.
func paInPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, paInSymbols) }

// paInMeridiemHour is the meridiemHour callback of pa-in.
func paInMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return dayPartHour(hour, meridiem, "ਰਾਤ", "ਸਵੇਰ", "ਦੁਪਹਿਰ", "ਸ਼ਾਮ", 10)
}

// paInMeridiem is the meridiem callback of pa-in.
func paInMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "ਰਾਤ"
	case hour < 10:
		return "ਸਵੇਰ"
	case hour < 17:
		return "ਦੁਪਹਿਰ"
	case hour < 20:
		return "ਸ਼ਾਮ"
	}
	return "ਰਾਤ"
}
