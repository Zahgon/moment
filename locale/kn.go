// moment.js locale configuration
// locale : Kannada [kn]
// author : Rajeev Naik : https://github.com/rajeevnaikte

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("kn", func() {
		moment.DefineLocale("kn", js.Obj(
			"months", []string{"ಜನವರಿ", "ಫೆಬ್ರವರಿ", "ಮಾರ್ಚ್", "ಏಪ್ರಿಲ್", "ಮೇ", "ಜೂನ್", "ಜುಲೈ", "ಆಗಸ್ಟ್", "ಸೆಪ್ಟೆಂಬರ್", "ಅಕ್ಟೋಬರ್", "ನವೆಂಬರ್", "ಡಿಸೆಂಬರ್"},
			"monthsShort", []string{"ಜನ", "ಫೆಬ್ರ", "ಮಾರ್ಚ್", "ಏಪ್ರಿಲ್", "ಮೇ", "ಜೂನ್", "ಜುಲೈ", "ಆಗಸ್ಟ್", "ಸೆಪ್ಟೆಂ", "ಅಕ್ಟೋ", "ನವೆಂ", "ಡಿಸೆಂ"},
			"monthsParseExact", true,
			"weekdays", []string{"ಭಾನುವಾರ", "ಸೋಮವಾರ", "ಮಂಗಳವಾರ", "ಬುಧವಾರ", "ಗುರುವಾರ", "ಶುಕ್ರವಾರ", "ಶನಿವಾರ"},
			"weekdaysShort", []string{"ಭಾನು", "ಸೋಮ", "ಮಂಗಳ", "ಬುಧ", "ಗುರು", "ಶುಕ್ರ", "ಶನಿ"},
			"weekdaysMin", []string{"ಭಾ", "ಸೋ", "ಮಂ", "ಬು", "ಗು", "ಶು", "ಶ"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm",
				"LTS", "A h:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm",
				"LLLL", "dddd, D MMMM YYYY, A h:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[ಇಂದು] LT",
				"nextDay", "[ನಾಳೆ] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[ನಿನ್ನೆ] LT",
				"lastWeek", "[ಕೊನೆಯ] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s ನಂತರ",
				"past", "%s ಹಿಂದೆ",
				"s", "ಕೆಲವು ಕ್ಷಣಗಳು",
				"ss", "%d ಸೆಕೆಂಡುಗಳು",
				"m", "ಒಂದು ನಿಮಿಷ",
				"mm", "%d ನಿಮಿಷ",
				"h", "ಒಂದು ಗಂಟೆ",
				"hh", "%d ಗಂಟೆ",
				"d", "ಒಂದು ದಿನ",
				"dd", "%d ದಿನ",
				"M", "ಒಂದು ತಿಂಗಳು",
				"MM", "%d ತಿಂಗಳು",
				"y", "ಒಂದು ವರ್ಷ",
				"yy", "%d ವರ್ಷ",
			),
			"preparse", knPreparse,
			"postformat", knPostformat,
			"meridiemParse", jsre.MustCompile("ರಾತ್ರಿ|ಬೆಳಿಗ್ಗೆ|ಮಧ್ಯಾಹ್ನ|ಸಂಜೆ", ""),
			"meridiemHour", knMeridiemHour,
			"meridiem", knMeridiem,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(ನೇ)", ""),
			"ordinal", knOrdinal,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	knSymbols, knNumbers = digitTables("೦೧೨೩೪೫೬೭೮೯")
	knDigit              = jsre.MustCompile(`[೧೨೩೪೫೬೭೮೯೦]`, "g")
)

// knPreparse reads Kannada digits.
func knPreparse(_ *moment.Locale, s string) string { return mapMatches(knDigit, s, knNumbers) }

// knPostformat writes Kannada digits.
func knPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, knSymbols) }

// knMeridiemHour is the meridiemHour callback of kn.
func knMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return dayPartHour(hour, meridiem, "ರಾತ್ರಿ", "ಬೆಳಿಗ್ಗೆ", "ಮಧ್ಯಾಹ್ನ", "ಸಂಜೆ", 10)
}

// knMeridiem is the meridiem callback of kn.
func knMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "ರಾತ್ರಿ"
	case hour < 10:
		return "ಬೆಳಿಗ್ಗೆ"
	case hour < 17:
		return "ಮಧ್ಯಾಹ್ನ"
	case hour < 20:
		return "ಸಂಜೆ"
	}
	return "ರಾತ್ರಿ"
}

// knOrdinal is the ordinal callback of kn.
func knOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return str(number) + "ನೇ"
}
