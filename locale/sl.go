// moment.js locale configuration
// locale : Slovenian [sl]
// author : Robert Sedovšek : https://github.com/sedovsek

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("sl", func() {
		moment.DefineLocale("sl", js.Obj(
			"months", []string{"januar", "februar", "marec", "april", "maj", "junij", "julij", "avgust", "september", "oktober", "november", "december"},
			"monthsShort", []string{"jan.", "feb.", "mar.", "apr.", "maj.", "jun.", "jul.", "avg.", "sep.", "okt.", "nov.", "dec."},
			"monthsParseExact", true,
			"weekdays", []string{"nedelja", "ponedeljek", "torek", "sreda", "četrtek", "petek", "sobota"},
			"weekdaysShort", []string{"ned.", "pon.", "tor.", "sre.", "čet.", "pet.", "sob."},
			"weekdaysMin", []string{"ne", "po", "to", "sr", "če", "pe", "so"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD. MM. YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY H:mm",
				"LLLL", "dddd, D. MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[danes ob] LT",
				"nextDay", "[jutri ob] LT",
				"nextWeek", slCalendarNextWeek,
				"lastDay", "[včeraj ob] LT",
				"lastWeek", slCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "čez %s",
				"past", "pred %s",
				"s", slRelativeTimeSeconds,
				"ss", slRelativeTimeSeconds,
				"m", slRelativeTimeSeconds,
				"mm", slRelativeTimeSeconds,
				"h", slRelativeTimeSeconds,
				"hh", slRelativeTimeSeconds,
				"d", slRelativeTimeSeconds,
				"dd", slRelativeTimeSeconds,
				"M", slRelativeTimeSeconds,
				"MM", slRelativeTimeSeconds,
				"y", slRelativeTimeSeconds,
				"yy", slRelativeTimeSeconds,
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

// slCalendarNextWeek is calendar.nextWeek of sl.
func slCalendarNextWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[v] [nedeljo] [ob] LT"
	case 3:
		return "[v] [sredo] [ob] LT"
	case 6:
		return "[v] [soboto] [ob] LT"
	case 1, 2, 4, 5:
		return "[v] dddd [ob] LT"
	}
	return ""
}

// slCalendarLastWeek is calendar.lastWeek of sl.
func slCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[prejšnjo] [nedeljo] [ob] LT"
	case 3:
		return "[prejšnjo] [sredo] [ob] LT"
	case 6:
		return "[prejšnjo] [soboto] [ob] LT"
	case 1, 2, 4, 5:
		return "[prejšnji] dddd [ob] LT"
	}
	return ""
}

// slCount picks the form for one, two, three or four, and five or more; each form has a plain variant (length
// of time or future) and a past variant.
func slCount(number float64, plain bool, forms [4][2]string) string {
	i := 3
	switch {
	case number == 1:
		i = 0
	case number == 2:
		i = 1
	case number < 5:
		i = 2
	}
	if plain {
		return num(number) + " " + forms[i][0]
	}
	return num(number) + " " + forms[i][1]
}

// slRelativeTimeSeconds is processRelativeTime of sl.
func slRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	plain := withoutSuffix || isFuture
	pick := func(a, b string) string {
		if plain {
			return a
		}
		return b
	}
	switch key {
	case "s":
		return pick("nekaj sekund", "nekaj sekundami")
	case "ss":
		if number == 1 {
			if withoutSuffix {
				return num(number) + " sekundo"
			}
			return num(number) + " sekundi"
		}
		return slCount(number, plain, [4][2]string{{}, {"sekundi", "sekundah"}, {"sekunde", "sekundah"}, {"sekund", "sekund"}})
	case "m":
		if withoutSuffix {
			return "ena minuta"
		}
		return "eno minuto"
	case "mm":
		if number == 1 {
			if withoutSuffix {
				return num(number) + " minuta"
			}
			return num(number) + " minuto"
		}
		return slCount(number, plain, [4][2]string{{}, {"minuti", "minutama"}, {"minute", "minutami"}, {"minut", "minutami"}})
	case "h":
		if withoutSuffix {
			return "ena ura"
		}
		return "eno uro"
	case "hh":
		if number == 1 {
			if withoutSuffix {
				return num(number) + " ura"
			}
			return num(number) + " uro"
		}
		return slCount(number, plain, [4][2]string{{}, {"uri", "urama"}, {"ure", "urami"}, {"ur", "urami"}})
	case "d":
		return pick("en dan", "enim dnem")
	case "dd":
		switch number {
		case 1:
			return num(number) + " " + pick("dan", "dnem")
		case 2:
			return num(number) + " " + pick("dni", "dnevoma")
		}
		return num(number) + " " + pick("dni", "dnevi")
	case "M":
		return pick("en mesec", "enim mesecem")
	case "MM":
		return slCount(number, plain, [4][2]string{{"mesec", "mesecem"}, {"meseca", "mesecema"}, {"mesece", "meseci"}, {"mesecev", "meseci"}})
	case "y":
		return pick("eno leto", "enim letom")
	case "yy":
		return slCount(number, plain, [4][2]string{{"leto", "letom"}, {"leti", "letoma"}, {"leta", "leti"}, {"let", "leti"}})
	}
	return ""
}
