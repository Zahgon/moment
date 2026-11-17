// moment.js locale configuration
// locale : Galician [gl]
// author : Juan G. Hurtado : https://github.com/juanghurtado

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("gl", func() {
		moment.DefineLocale("gl", js.Obj(
			"months", []string{"xaneiro", "febreiro", "marzo", "abril", "maio", "xuño", "xullo", "agosto", "setembro", "outubro", "novembro", "decembro"},
			"monthsShort", []string{"xan.", "feb.", "mar.", "abr.", "mai.", "xuñ.", "xul.", "ago.", "set.", "out.", "nov.", "dec."},
			"monthsParseExact", true,
			"weekdays", []string{"domingo", "luns", "martes", "mércores", "xoves", "venres", "sábado"},
			"weekdaysShort", []string{"dom.", "lun.", "mar.", "mér.", "xov.", "ven.", "sáb."},
			"weekdaysMin", []string{"do", "lu", "ma", "mé", "xo", "ve", "sá"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D [de] MMMM [de] YYYY",
				"LLL", "D [de] MMMM [de] YYYY H:mm",
				"LLLL", "dddd, D [de] MMMM [de] YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", glCalendarSameDay,
				"nextDay", glCalendarNextDay,
				"nextWeek", glCalendarNextWeek,
				"lastDay", glCalendarLastDay,
				"lastWeek", glCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", glRelativeTimeFuture,
				"past", "hai %s",
				"s", "uns segundos",
				"ss", "%d segundos",
				"m", "un minuto",
				"mm", "%d minutos",
				"h", "unha hora",
				"hh", "%d horas",
				"d", "un día",
				"dd", "%d días",
				"M", "un mes",
				"MM", "%d meses",
				"y", "un ano",
				"yy", "%d anos",
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

// glPlural is "ás" (or the given singular at one o'clock).
func glPlural(m *moment.Moment, singular string) string {
	if m.Hours() != 1 {
		return "ás"
	}
	return singular
}

// glCalendarSameDay is calendar.sameDay of gl.
func glCalendarSameDay(m, _ *moment.Moment) string { return "[hoxe " + glPlural(m, "á") + "] LT" }

// glCalendarNextDay is calendar.nextDay of gl.
func glCalendarNextDay(m, _ *moment.Moment) string { return "[mañá " + glPlural(m, "á") + "] LT" }

// glCalendarNextWeek is calendar.nextWeek of gl.
func glCalendarNextWeek(m, _ *moment.Moment) string { return "dddd [" + glPlural(m, "a") + "] LT" }

// glCalendarLastDay is calendar.lastDay of gl.
func glCalendarLastDay(m, _ *moment.Moment) string {
	if m.Hours() != 1 {
		return "[onte á] LT"
	}
	return "[onte a] LT"
}

// glCalendarLastWeek is calendar.lastWeek of gl.
func glCalendarLastWeek(m, _ *moment.Moment) string {
	return "[o] dddd [pasado " + glPlural(m, "a") + "] LT"
}

// glRelativeTimeFuture is relativeTime.future of gl.
func glRelativeTimeFuture(str string) string {
	if js.IndexOf16(str, "un") == 0 {
		return "n" + str
	}
	return "en " + str
}
