// moment.js locale configuration
// locale : Tamil [ta]
// author : Arjunkumar Krishnamoorthy : https://github.com/tk120404

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ta", func() {
		moment.DefineLocale("ta", js.Obj(
			"months", []string{"ஜனவரி", "பிப்ரவரி", "மார்ச்", "ஏப்ரல்", "மே", "ஜூன்", "ஜூலை", "ஆகஸ்ட்", "செப்டெம்பர்", "அக்டோபர்", "நவம்பர்", "டிசம்பர்"},
			"monthsShort", []string{"ஜனவரி", "பிப்ரவரி", "மார்ச்", "ஏப்ரல்", "மே", "ஜூன்", "ஜூலை", "ஆகஸ்ட்", "செப்டெம்பர்", "அக்டோபர்", "நவம்பர்", "டிசம்பர்"},
			"weekdays", []string{"ஞாயிற்றுக்கிழமை", "திங்கட்கிழமை", "செவ்வாய்கிழமை", "புதன்கிழமை", "வியாழக்கிழமை", "வெள்ளிக்கிழமை", "சனிக்கிழமை"},
			"weekdaysShort", []string{"ஞாயிறு", "திங்கள்", "செவ்வாய்", "புதன்", "வியாழன்", "வெள்ளி", "சனி"},
			"weekdaysMin", []string{"ஞா", "தி", "செ", "பு", "வி", "வெ", "ச"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, HH:mm",
				"LLLL", "dddd, D MMMM YYYY, HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[இன்று] LT",
				"nextDay", "[நாளை] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[நேற்று] LT",
				"lastWeek", "[கடந்த வாரம்] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s இல்",
				"past", "%s முன்",
				"s", "ஒரு சில விநாடிகள்",
				"ss", "%d விநாடிகள்",
				"m", "ஒரு நிமிடம்",
				"mm", "%d நிமிடங்கள்",
				"h", "ஒரு மணி நேரம்",
				"hh", "%d மணி நேரம்",
				"d", "ஒரு நாள்",
				"dd", "%d நாட்கள்",
				"M", "ஒரு மாதம்",
				"MM", "%d மாதங்கள்",
				"y", "ஒரு வருடம்",
				"yy", "%d ஆண்டுகள்",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}வது", ""),
			"ordinal", taOrdinal,
			"preparse", taPreparse,
			"postformat", taPostformat,
			"meridiemParse", jsre.MustCompile("யாமம்|வைகறை|காலை|நண்பகல்|எற்பாடு|மாலை", ""),
			"meridiem", taMeridiem,
			"meridiemHour", taMeridiemHour,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 6.0,
			),
		))
	})
}

var (
	taSymbols, taNumbers = digitTables("௦௧௨௩௪௫௬௭௮௯")
	taDigit              = jsre.MustCompile(`[௧௨௩௪௫௬௭௮௯௦]`, "g")
)

// taOrdinal is the ordinal callback of ta.
func taOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	return str(number) + "வது"
}

// taPreparse reads Tamil digits.
func taPreparse(_ *moment.Locale, s string) string { return mapMatches(taDigit, s, taNumbers) }

// taPostformat writes Tamil digits.
func taPostformat(_ *moment.Locale, s string) string { return mapMatches(asciiDigit, s, taSymbols) }

// taMeridiem is the meridiem callback of ta (its words start with a space).
func taMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 2:
		return " யாமம்"
	case hour < 6:
		return " வைகறை"
	case hour < 10:
		return " காலை"
	case hour < 14:
		return " நண்பகல்"
	case hour < 18:
		return " எற்பாடு"
	case hour < 22:
		return " மாலை"
	}
	return " யாமம்"
}

// taMeridiemHour is the meridiemHour callback of ta.
func taMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "யாமம்":
		if hourNumber(hour) < 2 {
			return hour
		}
		return plus12(hour)
	case "வைகறை", "காலை":
		return hour
	case "நண்பகல்":
		if hourNumber(hour) >= 10 {
			return hour
		}
		return plus12(hour)
	}
	return plus12(hour)
}
