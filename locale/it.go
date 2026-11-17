// moment.js locale configuration
// locale : Italian [it]
// author : Lorenzo : https://github.com/aliem
// author: Mattia Larentis: https://github.com/nostalgiaz
// author: Marco : https://github.com/Manfre98

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("it", func() {
		moment.DefineLocale("it", js.Obj(
			"months", []string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"},
			"monthsShort", []string{"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
			"weekdays", []string{"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"},
			"weekdaysShort", []string{"dom", "lun", "mar", "mer", "gio", "ven", "sab"},
			"weekdaysMin", []string{"do", "lu", "ma", "me", "gi", "ve", "sa"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", itCalendarSameDay,
				"nextDay", itCalendarNextDay,
				"nextWeek", itCalendarNextWeek,
				"lastDay", itCalendarLastDay,
				"lastWeek", itCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "tra %s",
				"past", "%s fa",
				"s", "alcuni secondi",
				"ss", "%d secondi",
				"m", "un minuto",
				"mm", "%d minuti",
				"h", "un'ora",
				"hh", "%d ore",
				"d", "un giorno",
				"dd", "%d giorni",
				"w", "una settimana",
				"ww", "%d settimane",
				"M", "un mese",
				"MM", "%d mesi",
				"y", "un anno",
				"yy", "%d anni",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}º", ""),
			"ordinal", "%dº",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// itAt is the article after "a": "lle " after one o'clock, " " at midnight, "ll'" at one.
func itAt(m *moment.Moment) string {
	switch h := m.Hours(); {
	case h > 1:
		return "lle "
	case h == 0:
		return " "
	}
	return "ll'"
}

// itCalendarSameDay is calendar.sameDay of it.
func itCalendarSameDay(m, _ *moment.Moment) string { return "[Oggi a" + itAt(m) + "]LT" }

// itCalendarNextDay is calendar.nextDay of it.
func itCalendarNextDay(m, _ *moment.Moment) string { return "[Domani a" + itAt(m) + "]LT" }

// itCalendarNextWeek is calendar.nextWeek of it.
func itCalendarNextWeek(m, _ *moment.Moment) string { return "dddd [a" + itAt(m) + "]LT" }

// itCalendarLastDay is calendar.lastDay of it.
func itCalendarLastDay(m, _ *moment.Moment) string { return "[Ieri a" + itAt(m) + "]LT" }

// itCalendarLastWeek is calendar.lastWeek of it.
func itCalendarLastWeek(m, _ *moment.Moment) string {
	if m.Day() == 0 {
		return "[La scorsa] dddd [a" + itAt(m) + "]LT"
	}
	return "[Lo scorso] dddd [a" + itAt(m) + "]LT"
}
