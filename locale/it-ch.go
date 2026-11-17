// moment.js locale configuration
// locale : Italian (Switzerland) [it-ch]
// author : xfh : https://github.com/xfh

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("it-ch", func() {
		moment.DefineLocale("it-ch", js.Obj(
			"months", []string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"},
			"monthsShort", []string{"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
			"weekdays", []string{"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"},
			"weekdaysShort", []string{"dom", "lun", "mar", "mer", "gio", "ven", "sab"},
			"weekdaysMin", []string{"do", "lu", "ma", "me", "gi", "ve", "sa"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Oggi alle] LT",
				"nextDay", "[Domani alle] LT",
				"nextWeek", "dddd [alle] LT",
				"lastDay", "[Ieri alle] LT",
				"lastWeek", itChCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", itChRelativeTimeFuture,
				"past", "%s fa",
				"s", "alcuni secondi",
				"ss", "%d secondi",
				"m", "un minuto",
				"mm", "%d minuti",
				"h", "un'ora",
				"hh", "%d ore",
				"d", "un giorno",
				"dd", "%d giorni",
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

var itChNumberFirst = jsre.MustCompile(`^[0-9].+$`, "")

// itChCalendarLastWeek is calendar.lastWeek of it-ch.
func itChCalendarLastWeek(m, _ *moment.Moment) string {
	if m.Day() == 0 {
		return "[la scorsa] dddd [alle] LT"
	}
	return "[lo scorso] dddd [alle] LT"
}

// itChRelativeTimeFuture is relativeTime.future of it-ch.
func itChRelativeTimeFuture(s string) string {
	if itChNumberFirst.Test(s) {
		return "tra " + s
	}
	return "in " + s
}
