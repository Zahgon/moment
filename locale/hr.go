// moment.js locale configuration
// locale : Croatian [hr]
// author : Bojan Marković : https://github.com/bmarkovic

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("hr", func() {
		moment.DefineLocale("hr", js.Obj(
			"months", js.Obj(
				"format", []string{"siječnja", "veljače", "ožujka", "travnja", "svibnja", "lipnja", "srpnja", "kolovoza", "rujna", "listopada", "studenoga", "prosinca"},
				"standalone", []string{"siječanj", "veljača", "ožujak", "travanj", "svibanj", "lipanj", "srpanj", "kolovoz", "rujan", "listopad", "studeni", "prosinac"},
			),
			"monthsShort", []string{"sij.", "velj.", "ožu.", "tra.", "svi.", "lip.", "srp.", "kol.", "ruj.", "lis.", "stu.", "pro."},
			"monthsParseExact", true,
			"weekdays", []string{"nedjelja", "ponedjeljak", "utorak", "srijeda", "četvrtak", "petak", "subota"},
			"weekdaysShort", []string{"ned.", "pon.", "uto.", "sri.", "čet.", "pet.", "sub."},
			"weekdaysMin", []string{"ne", "po", "ut", "sr", "če", "pe", "su"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "Do MMMM YYYY",
				"LLL", "Do MMMM YYYY H:mm",
				"LLLL", "dddd, Do MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[danas u] LT",
				"nextDay", "[sutra u] LT",
				"nextWeek", hrCalendarNextWeek,
				"lastDay", "[jučer u] LT",
				"lastWeek", hrCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "za %s",
				"past", "prije %s",
				"s", "par sekundi",
				"ss", hrRelativeTimeSecondsN,
				"m", hrRelativeTimeSecondsN,
				"mm", hrRelativeTimeSecondsN,
				"h", hrRelativeTimeSecondsN,
				"hh", hrRelativeTimeSecondsN,
				"d", "dan",
				"dd", hrRelativeTimeSecondsN,
				"M", "mjesec",
				"MM", hrRelativeTimeSecondsN,
				"y", "godinu",
				"yy", hrRelativeTimeSecondsN,
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

// hrCalendarNextWeek is calendar.nextWeek of hr.
func hrCalendarNextWeek(m, _ *moment.Moment) string { return bsCalendarNextWeekText(m) }

// hrCalendarLastWeek is calendar.lastWeek of hr.
func hrCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[prošlu] [nedjelju] [u] LT"
	case 3:
		return "[prošlu] [srijedu] [u] LT"
	case 6:
		return "[prošle] [subote] [u] LT"
	case 1, 2, 4, 5:
		return "[prošli] dddd [u] LT"
	}
	return ""
}

// bsCalendarNextWeekText is the next-week calendar text shared by hr and bs.
func bsCalendarNextWeekText(m *moment.Moment) string { return bsCalendarNextWeek(m, nil) }

// hrRelativeTimeSecondsN is translate of hr.
func hrRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, _ bool) string {
	switch key {
	case "m":
		if withoutSuffix {
			return "jedna minuta"
		}
		return "jedne minute"
	case "h":
		if withoutSuffix {
			return "jedan sat"
		}
		return "jednog sata"
	}
	return bsRelativeTimeSecondsN(number, withoutSuffix, key, false)
}
