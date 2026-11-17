// moment.js locale configuration
// locale : Catalan [ca]
// author : Juan G. Hurtado : https://github.com/juanghurtado

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ca", func() {
		moment.DefineLocale("ca", js.Obj(
			"months", js.Obj(
				"standalone", []string{"gener", "febrer", "març", "abril", "maig", "juny", "juliol", "agost", "setembre", "octubre", "novembre", "desembre"},
				"format", []string{"de gener", "de febrer", "de març", "d’abril", "de maig", "de juny", "de juliol", "d’agost", "de setembre", "d’octubre", "de novembre", "de desembre"},
				"isFormat", jsre.MustCompile("D[oD]?(\\s)+MMMM", ""),
			),
			"monthsShort", []string{"gen.", "febr.", "març", "abr.", "maig", "juny", "jul.", "ag.", "set.", "oct.", "nov.", "des."},
			"monthsParseExact", true,
			"weekdays", []string{"diumenge", "dilluns", "dimarts", "dimecres", "dijous", "divendres", "dissabte"},
			"weekdaysShort", []string{"dg.", "dl.", "dt.", "dc.", "dj.", "dv.", "ds."},
			"weekdaysMin", []string{"dg", "dl", "dt", "dc", "dj", "dv", "ds"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM [de] YYYY",
				"ll", "D MMM YYYY",
				"LLL", "D MMMM [de] YYYY [a les] H:mm",
				"lll", "D MMM YYYY, H:mm",
				"LLLL", "dddd D MMMM [de] YYYY [a les] H:mm",
				"llll", "ddd D MMM YYYY, H:mm",
			),
			"calendar", js.Obj(
				"sameDay", caCalendarSameDay,
				"nextDay", caCalendarNextDay,
				"nextWeek", caCalendarNextWeek,
				"lastDay", caCalendarLastDay,
				"lastWeek", caCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "d'aquí %s",
				"past", "fa %s",
				"s", "uns segons",
				"ss", "%d segons",
				"m", "un minut",
				"mm", "%d minuts",
				"h", "una hora",
				"hh", "%d hores",
				"d", "un dia",
				"dd", "%d dies",
				"M", "un mes",
				"MM", "%d mesos",
				"y", "un any",
				"yy", "%d anys",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(r|n|t|è|a)", ""),
			"ordinal", caOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// caArticle is "les", or "la" at one o'clock.
func caArticle(m *moment.Moment) string {
	if m.Hours() != 1 {
		return "les"
	}
	return "la"
}

// caCalendarSameDay is calendar.sameDay of ca.
func caCalendarSameDay(m, _ *moment.Moment) string { return "[avui a " + caArticle(m) + "] LT" }

// caCalendarNextDay is calendar.nextDay of ca.
func caCalendarNextDay(m, _ *moment.Moment) string { return "[demà a " + caArticle(m) + "] LT" }

// caCalendarNextWeek is calendar.nextWeek of ca.
func caCalendarNextWeek(m, _ *moment.Moment) string { return "dddd [a " + caArticle(m) + "] LT" }

// caCalendarLastDay is calendar.lastDay of ca.
func caCalendarLastDay(m, _ *moment.Moment) string { return "[ahir a " + caArticle(m) + "] LT" }

// caCalendarLastWeek is calendar.lastWeek of ca.
func caCalendarLastWeek(m, _ *moment.Moment) string {
	return "[el] dddd [passat a " + caArticle(m) + "] LT"
}

// caOrdinal is the ordinal callback of ca.
func caOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	output := "è"
	switch {
	case sameNumber(number, 1), sameNumber(number, 3):
		output = "r"
	case sameNumber(number, 2):
		output = "n"
	case sameNumber(number, 4):
		output = "t"
	}
	if period == "w" || period == "W" {
		output = "a"
	}
	return str(number) + output
}
