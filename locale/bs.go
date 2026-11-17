// moment.js locale configuration
// locale : Bosnian [bs]
// author : Nedim Cholich : https://github.com/frontyard
// author : Rasid Redzic : https://github.com/rasidre
// based on (hr) translation by Bojan Marković

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("bs", func() {
		moment.DefineLocale("bs", js.Obj(
			"months", []string{"januar", "februar", "mart", "april", "maj", "juni", "juli", "august", "septembar", "oktobar", "novembar", "decembar"},
			"monthsShort", []string{"jan.", "feb.", "mar.", "apr.", "maj.", "jun.", "jul.", "aug.", "sep.", "okt.", "nov.", "dec."},
			"monthsParseExact", true,
			"weekdays", []string{"nedjelja", "ponedjeljak", "utorak", "srijeda", "četvrtak", "petak", "subota"},
			"weekdaysShort", []string{"ned.", "pon.", "uto.", "sri.", "čet.", "pet.", "sub."},
			"weekdaysMin", []string{"ne", "po", "ut", "sr", "če", "pe", "su"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY H:mm",
				"LLLL", "dddd, D. MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[danas u] LT",
				"nextDay", "[sutra u] LT",
				"nextWeek", bsCalendarNextWeek,
				"lastDay", "[jučer u] LT",
				"lastWeek", bsCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "za %s",
				"past", "prije %s",
				"s", "par sekundi",
				"ss", bsRelativeTimeSecondsN,
				"m", bsRelativeTimeMinute,
				"mm", bsRelativeTimeSecondsN,
				"h", bsRelativeTimeSecondsN,
				"hh", bsRelativeTimeSecondsN,
				"d", "dan",
				"dd", bsRelativeTimeSecondsN,
				"M", "mjesec",
				"MM", bsRelativeTimeSecondsN,
				"y", "godinu",
				"yy", bsRelativeTimeSecondsN,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}\\.", ""),
			"ordinal", "%d.",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

// bsCalendarNextWeek is calendar.nextWeek of bs.
func bsCalendarNextWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[u] [nedjelju] [u] LT"
	case 3:
		return "[u] [srijedu] [u] LT"
	case 6:
		return "[u] [subotu] [u] LT"
	case 1, 2, 4, 5:
		return "[u] dddd [u] LT"
	}
	return ""
}

// bsCalendarLastWeek is calendar.lastWeek of bs.
func bsCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0, 3:
		return "[prošlu] dddd [u] LT"
	case 6:
		return "[prošle] [subote] [u] LT"
	case 1, 2, 4, 5:
		return "[prošli] dddd [u] LT"
	}
	return ""
}

// bsForms appends the singular, paucal (2 to 4) or plural form to the number.
func bsForms(number float64, one, few, many string) string {
	switch {
	case number == 1:
		return num(number) + " " + one
	case number == 2 || number == 3 || number == 4:
		return num(number) + " " + few
	}
	return num(number) + " " + many
}

// bsRelativeTimeSecondsN is translate of bs.
func bsRelativeTimeSecondsN(number float64, _ bool, key string, _ bool) string {
	switch key {
	case "ss":
		return bsForms(number, "sekunda", "sekunde", "sekundi")
	case "mm":
		return bsForms(number, "minuta", "minute", "minuta")
	case "h":
		return "jedan sat"
	case "hh":
		return bsForms(number, "sat", "sata", "sati")
	case "dd":
		return bsForms(number, "dan", "dana", "dana")
	case "MM":
		return bsForms(number, "mjesec", "mjeseca", "mjeseci")
	case "yy":
		return bsForms(number, "godina", "godine", "godina")
	}
	return ""
}

// bsRelativeTimeMinute is processRelativeTime of bs.
func bsRelativeTimeMinute(_ float64, withoutSuffix bool, key string, isFuture bool) string {
	if key != "m" {
		return ""
	}
	switch {
	case withoutSuffix:
		return "jedna minuta"
	case isFuture:
		return "jednu minutu"
	}
	return "jedne minute"
}
