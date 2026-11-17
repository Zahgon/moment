// moment.js locale configuration
// locale : Hebrew [he]
// author : Tomer Cohen : https://github.com/tomer
// author : Moshe Simantov : https://github.com/DevelopmentIL
// author : Tal Ater : https://github.com/TalAter

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("he", func() {
		moment.DefineLocale("he", js.Obj(
			"months", []string{"ינואר", "פברואר", "מרץ", "אפריל", "מאי", "יוני", "יולי", "אוגוסט", "ספטמבר", "אוקטובר", "נובמבר", "דצמבר"},
			"monthsShort", []string{"ינו׳", "פבר׳", "מרץ", "אפר׳", "מאי", "יוני", "יולי", "אוג׳", "ספט׳", "אוק׳", "נוב׳", "דצמ׳"},
			"weekdays", []string{"ראשון", "שני", "שלישי", "רביעי", "חמישי", "שישי", "שבת"},
			"weekdaysShort", []string{"א׳", "ב׳", "ג׳", "ד׳", "ה׳", "ו׳", "ש׳"},
			"weekdaysMin", []string{"א", "ב", "ג", "ד", "ה", "ו", "ש"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D [ב]MMMM YYYY",
				"LLL", "D [ב]MMMM YYYY HH:mm",
				"LLLL", "dddd, D [ב]MMMM YYYY HH:mm",
				"l", "D/M/YYYY",
				"ll", "D MMM YYYY",
				"lll", "D MMM YYYY HH:mm",
				"llll", "ddd, D MMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[היום ב־]LT",
				"nextDay", "[מחר ב־]LT",
				"nextWeek", "dddd [בשעה] LT",
				"lastDay", "[אתמול ב־]LT",
				"lastWeek", "[ביום] dddd [האחרון בשעה] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "בעוד %s",
				"past", "לפני %s",
				"s", "מספר שניות",
				"ss", "%d שניות",
				"m", "דקה",
				"mm", "%d דקות",
				"h", "שעה",
				"hh", heRelativeTimeHours,
				"d", "יום",
				"dd", heRelativeTimeDays,
				"M", "חודש",
				"MM", heRelativeTimeMonths,
				"y", "שנה",
				"yy", heRelativeTimeYears,
			),
			"meridiemParse", jsre.MustCompile("אחה\"צ|לפנה\"צ|אחרי הצהריים|לפני הצהריים|לפנות בוקר|בבוקר|בערב", "i"),
			"isPM", heIsPM,
			"meridiem", heMeridiem,
		))
	})
}

var hePM = jsre.MustCompile(`^(אחה"צ|אחרי הצהריים|בערב)$`, "")

// heDual returns the dual word for 2 and the number with the plural word otherwise.
func heDual(number float64, dual, plural string) string {
	if number == 2 {
		return dual
	}
	return num(number) + " " + plural
}

// heRelativeTimeHours is relativeTime.hh of he.
func heRelativeTimeHours(number float64, _ bool, _ string, _ bool) string {
	return heDual(number, "שעתיים", "שעות")
}

// heRelativeTimeDays is relativeTime.dd of he.
func heRelativeTimeDays(number float64, _ bool, _ string, _ bool) string {
	return heDual(number, "יומיים", "ימים")
}

// heRelativeTimeMonths is relativeTime.MM of he.
func heRelativeTimeMonths(number float64, _ bool, _ string, _ bool) string {
	return heDual(number, "חודשיים", "חודשים")
}

// heRelativeTimeYears is relativeTime.yy of he.
func heRelativeTimeYears(number float64, _ bool, _ string, _ bool) string {
	if number != 2 && js.Mod(number, 10) == 0 && number != 10 {
		return num(number) + " שנה"
	}
	return heDual(number, "שנתיים", "שנים")
}

// heIsPM is the isPM callback of he.
func heIsPM(_ *moment.Locale, input string) bool { return hePM.Test(input) }

// heMeridiem is the meridiem callback of he.
func heMeridiem(_ *moment.Locale, hour, _ float64, isLower bool) string {
	switch {
	case hour < 5:
		return "לפנות בוקר"
	case hour < 10:
		return "בבוקר"
	case hour < 12:
		if isLower {
			return `לפנה"צ`
		}
		return "לפני הצהריים"
	case hour < 18:
		if isLower {
			return `אחה"צ`
		}
		return "אחרי הצהריים"
	}
	return "בערב"
}
