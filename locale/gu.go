// moment.js locale configuration
// locale : Gujarati [gu]
// author : Kaushik Thanki : https://github.com/Kaushik1987

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("gu", func() {
		moment.DefineLocale("gu", js.Obj(
			"months", []string{"જાન્યુઆરી", "ફેબ્રુઆરી", "માર્ચ", "એપ્રિલ", "મે", "જૂન", "જુલાઈ", "ઑગસ્ટ", "સપ્ટેમ્બર", "ઑક્ટ્બર", "નવેમ્બર", "ડિસેમ્બર"},
			"monthsShort", []string{"જાન્યુ.", "ફેબ્રુ.", "માર્ચ", "એપ્રિ.", "મે", "જૂન", "જુલા.", "ઑગ.", "સપ્ટે.", "ઑક્ટ્.", "નવે.", "ડિસે."},
			"monthsParseExact", true,
			"weekdays", []string{"રવિવાર", "સોમવાર", "મંગળવાર", "બુધ્વાર", "ગુરુવાર", "શુક્રવાર", "શનિવાર"},
			"weekdaysShort", []string{"રવિ", "સોમ", "મંગળ", "બુધ્", "ગુરુ", "શુક્ર", "શનિ"},
			"weekdaysMin", []string{"ર", "સો", "મં", "બુ", "ગુ", "શુ", "શ"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm વાગ્યે",
				"LTS", "A h:mm:ss વાગ્યે",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm વાગ્યે",
				"LLLL", "dddd, D MMMM YYYY, A h:mm વાગ્યે",
			),
			"calendar", js.Obj(
				"sameDay", "[આજ] LT",
				"nextDay", "[કાલે] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[ગઇકાલે] LT",
				"lastWeek", "[પાછલા] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s મા",
				"past", "%s પહેલા",
				"s", "અમુક પળો",
				"ss", "%d સેકંડ",
				"m", "એક મિનિટ",
				"mm", "%d મિનિટ",
				"h", "એક કલાક",
				"hh", "%d કલાક",
				"d", "એક દિવસ",
				"dd", "%d દિવસ",
				"M", "એક મહિનો",
				"MM", "%d મહિનો",
				"y", "એક વર્ષ",
				"yy", "%d વર્ષ",
			),
			"preparse", guPreparse,
			"postformat", guPostformat,
			"meridiemParse", jsre.MustCompile("રાત|બપોર|સવાર|સાંજ", ""),
			"meridiemHour", guMeridiemHour,
			"meridiem", guMeridiem,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	guSymbols, guNumbers = digitTables("૦૧૨૩૪૫૬૭૮૯")
	guDigit              = jsre.MustCompile(`[૧૨૩૪૫૬૭૮૯૦]`, "g")
)

// guPreparse reads Gujarati digits.
func guPreparse(_ *moment.Locale, s string) string { return mapMatches(guDigit, s, guNumbers) }

// guPostformat writes Gujarati digits.
func guPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, guSymbols) }

// guMeridiemHour is the meridiemHour callback of gu.
func guMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "રાત":
		if hourNumber(hour) < 4 {
			return hour
		}
		return plus12(hour)
	case "સવાર":
		return hour
	case "બપોર":
		if hourNumber(hour) >= 10 {
			return hour
		}
		return plus12(hour)
	case "સાંજ":
		return plus12(hour)
	}
	return js.Undef
}

// guMeridiem is the meridiem callback of gu.
func guMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "રાત"
	case hour < 10:
		return "સવાર"
	case hour < 17:
		return "બપોર"
	case hour < 20:
		return "સાંજ"
	}
	return "રાત"
}
