// moment.js locale configuration
// locale : Bulgarian [bg]
// author : Krasen Borisov : https://github.com/kraz

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("bg", func() {
		moment.DefineLocale("bg", js.Obj(
			"months", []string{"януари", "февруари", "март", "април", "май", "юни", "юли", "август", "септември", "октомври", "ноември", "декември"},
			"monthsShort", []string{"яну", "фев", "мар", "апр", "май", "юни", "юли", "авг", "сеп", "окт", "ное", "дек"},
			"weekdays", []string{"неделя", "понеделник", "вторник", "сряда", "четвъртък", "петък", "събота"},
			"weekdaysShort", []string{"нед", "пон", "вто", "сря", "чет", "пет", "съб"},
			"weekdaysMin", []string{"нд", "пн", "вт", "ср", "чт", "пт", "сб"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "D.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY H:mm",
				"LLLL", "dddd, D MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Днес в] LT",
				"nextDay", "[Утре в] LT",
				"nextWeek", "dddd [в] LT",
				"lastDay", "[Вчера в] LT",
				"lastWeek", bgCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "след %s",
				"past", "преди %s",
				"s", "няколко секунди",
				"ss", "%d секунди",
				"m", "минута",
				"mm", "%d минути",
				"h", "час",
				"hh", "%d часа",
				"d", "ден",
				"dd", "%d дена",
				"w", "седмица",
				"ww", "%d седмици",
				"M", "месец",
				"MM", "%d месеца",
				"y", "година",
				"yy", "%d години",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(ев|ен|ти|ви|ри|ми)", ""),
			"ordinal", bgOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

// bgCalendarLastWeek is calendar.lastWeek of bg.
func bgCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0, 3, 6:
		return "[Миналата] dddd [в] LT"
	case 1, 2, 4, 5:
		return "[Миналия] dddd [в] LT"
	}
	return ""
}

// bgOrdinal is the ordinal callback of bg.
func bgOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	n := js.ToNumber(number)
	lastDigit, last2Digits := js.Mod(n, 10), js.Mod(n, 100)
	switch {
	case sameNumber(number, 0):
		return str(number) + "-ев"
	case last2Digits == 0:
		return str(number) + "-ен"
	case last2Digits > 10 && last2Digits < 20:
		return str(number) + "-ти"
	case lastDigit == 1:
		return str(number) + "-ви"
	case lastDigit == 2:
		return str(number) + "-ри"
	case lastDigit == 7 || lastDigit == 8:
		return str(number) + "-ми"
	}
	return str(number) + "-ти"
}
