// moment.js locale configuration
// locale : Slovak [sk]
// author : Martin Minka : https://github.com/k2s
// based on work of petrbela : https://github.com/petrbela

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("sk", func() {
		moment.DefineLocale("sk", js.Obj(
			"months", []string{"január", "február", "marec", "apríl", "máj", "jún", "júl", "august", "september", "október", "november", "december"},
			"monthsShort", []string{"jan", "feb", "mar", "apr", "máj", "jún", "júl", "aug", "sep", "okt", "nov", "dec"},
			"weekdays", []string{"nedeľa", "pondelok", "utorok", "streda", "štvrtok", "piatok", "sobota"},
			"weekdaysShort", []string{"ne", "po", "ut", "st", "št", "pi", "so"},
			"weekdaysMin", []string{"ne", "po", "ut", "st", "št", "pi", "so"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY H:mm",
				"LLLL", "dddd D. MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[dnes o] LT",
				"nextDay", "[zajtra o] LT",
				"nextWeek", skCalendarNextWeek,
				"lastDay", "[včera o] LT",
				"lastWeek", skCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "za %s",
				"past", "pred %s",
				"s", skRelativeTimeSeconds,
				"ss", skRelativeTimeSeconds,
				"m", skRelativeTimeSeconds,
				"mm", skRelativeTimeSeconds,
				"h", skRelativeTimeSeconds,
				"hh", skRelativeTimeSeconds,
				"d", skRelativeTimeSeconds,
				"dd", skRelativeTimeSeconds,
				"M", skRelativeTimeSeconds,
				"MM", skRelativeTimeSeconds,
				"y", skRelativeTimeSeconds,
				"yy", skRelativeTimeSeconds,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}\\.", ""),
			"ordinal", "%d.",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// skCalendarNextWeek is calendar.nextWeek of sk.
func skCalendarNextWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[v nedeľu o] LT"
	case 1, 2:
		return "[v] dddd [o] LT"
	case 3:
		return "[v stredu o] LT"
	case 4:
		return "[vo štvrtok o] LT"
	case 5:
		return "[v piatok o] LT"
	case 6:
		return "[v sobotu o] LT"
	}
	return ""
}

// skCalendarLastWeek is calendar.lastWeek of sk.
func skCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[minulú nedeľu o] LT"
	case 1, 2, 4, 5:
		return "[minulý] dddd [o] LT"
	case 3:
		return "[minulú stredu o] LT"
	case 6:
		return "[minulú sobotu o] LT"
	}
	return ""
}

// skRelativeTimeSeconds is translate of sk.
func skRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	result := num(number) + " "
	plural := number > 1 && number < 5
	counted := func(few, many, past string) string {
		if withoutSuffix || isFuture {
			if plural {
				return result + few
			}
			return result + many
		}
		return result + past
	}
	futureOr := func(a, past string) string {
		if withoutSuffix || isFuture {
			return a
		}
		return past
	}
	pick := func(a, b, c string) string {
		switch {
		case withoutSuffix:
			return a
		case isFuture:
			return b
		}
		return c
	}
	switch key {
	case "s":
		return futureOr("pár sekúnd", "pár sekundami")
	case "ss":
		return counted("sekundy", "sekúnd", "sekundami")
	case "m":
		return pick("minúta", "minútu", "minútou")
	case "mm":
		return counted("minúty", "minút", "minútami")
	case "h":
		return pick("hodina", "hodinu", "hodinou")
	case "hh":
		return counted("hodiny", "hodín", "hodinami")
	case "d":
		return futureOr("deň", "dňom")
	case "dd":
		return counted("dni", "dní", "dňami")
	case "M":
		return futureOr("mesiac", "mesiacom")
	case "MM":
		return counted("mesiace", "mesiacov", "mesiacmi")
	case "y":
		return futureOr("rok", "rokom")
	case "yy":
		return counted("roky", "rokov", "rokmi")
	}
	return ""
}
