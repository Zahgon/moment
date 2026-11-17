// moment.js locale configuration
// locale : Macedonian [mk]
// author : Borislav Mickov : https://github.com/B0k0
// author : Sashko Todorov : https://github.com/bkyceh

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("mk", func() {
		moment.DefineLocale("mk", js.Obj(
			"months", []string{"јануари", "февруари", "март", "април", "мај", "јуни", "јули", "август", "септември", "октомври", "ноември", "декември"},
			"monthsShort", []string{"јан", "фев", "мар", "апр", "мај", "јун", "јул", "авг", "сеп", "окт", "ное", "дек"},
			"weekdays", []string{"недела", "понеделник", "вторник", "среда", "четврток", "петок", "сабота"},
			"weekdaysShort", []string{"нед", "пон", "вто", "сре", "чет", "пет", "саб"},
			"weekdaysMin", []string{"нe", "пo", "вт", "ср", "че", "пе", "сa"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "D.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY H:mm",
				"LLLL", "dddd, D MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Денес во] LT",
				"nextDay", "[Утре во] LT",
				"nextWeek", "[Во] dddd [во] LT",
				"lastDay", "[Вчера во] LT",
				"lastWeek", mkCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "за %s",
				"past", "пред %s",
				"s", "неколку секунди",
				"ss", "%d секунди",
				"m", "една минута",
				"mm", "%d минути",
				"h", "еден час",
				"hh", "%d часа",
				"d", "еден ден",
				"dd", "%d дена",
				"M", "еден месец",
				"MM", "%d месеци",
				"y", "една година",
				"yy", "%d години",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}-(ев|ен|ти|ви|ри|ми)", ""),
			"ordinal", mkOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

// mkCalendarLastWeek is calendar.lastWeek of mk.
func mkCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0, 3, 6:
		return "[Изминатата] dddd [во] LT"
	case 1, 2, 4, 5:
		return "[Изминатиот] dddd [во] LT"
	}
	return ""
}

// mkOrdinal is the ordinal callback of mk (the rules of bg).
func mkOrdinal(l *moment.Locale, number moment.Value, period string) moment.Value {
	return bgOrdinal(l, number, period)
}
