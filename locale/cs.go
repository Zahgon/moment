// moment.js locale configuration
// locale : Czech [cs]
// author : petrbela : https://github.com/petrbela

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("cs", func() {
		moment.DefineLocale("cs", js.Obj(
			"months", js.Obj(
				"standalone", []string{"leden", "únor", "březen", "duben", "květen", "červen", "červenec", "srpen", "září", "říjen", "listopad", "prosinec"},
				"format", []string{"ledna", "února", "března", "dubna", "května", "června", "července", "srpna", "září", "října", "listopadu", "prosince"},
				"isFormat", jsre.MustCompile("DD?[o.]?(\\[[^\\[\\]]*\\]|\\s)+MMMM", ""),
			),
			"monthsShort", []string{"led", "úno", "bře", "dub", "kvě", "čvn", "čvc", "srp", "zář", "říj", "lis", "pro"},
			"monthsRegex", jsre.MustCompile("^(leden|únor|březen|duben|květen|červenec|července|červen|června|srpen|září|říjen|listopad|prosinec|led|úno|bře|dub|kvě|čvn|čvc|srp|zář|říj|lis|pro)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(leden|únor|březen|duben|květen|červenec|července|červen|června|srpen|září|říjen|listopad|prosinec|led|úno|bře|dub|kvě|čvn|čvc|srp|zář|říj|lis|pro)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(leden|ledna|února|únor|březen|března|duben|dubna|květen|května|červenec|července|červen|června|srpen|srpna|září|říjen|října|listopadu|listopad|prosinec|prosince)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(led|úno|bře|dub|kvě|čvn|čvc|srp|zář|říj|lis|pro)", "i"),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^led", "i"),
				jsre.MustCompile("^úno", "i"),
				jsre.MustCompile("^bře", "i"),
				jsre.MustCompile("^dub", "i"),
				jsre.MustCompile("^kvě", "i"),
				jsre.MustCompile("^(čvn|červen$|června)", "i"),
				jsre.MustCompile("^(čvc|červenec|července)", "i"),
				jsre.MustCompile("^srp", "i"),
				jsre.MustCompile("^zář", "i"),
				jsre.MustCompile("^říj", "i"),
				jsre.MustCompile("^lis", "i"),
				jsre.MustCompile("^pro", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^led", "i"),
				jsre.MustCompile("^úno", "i"),
				jsre.MustCompile("^bře", "i"),
				jsre.MustCompile("^dub", "i"),
				jsre.MustCompile("^kvě", "i"),
				jsre.MustCompile("^(čvn|červen$|června)", "i"),
				jsre.MustCompile("^(čvc|červenec|července)", "i"),
				jsre.MustCompile("^srp", "i"),
				jsre.MustCompile("^zář", "i"),
				jsre.MustCompile("^říj", "i"),
				jsre.MustCompile("^lis", "i"),
				jsre.MustCompile("^pro", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^led", "i"),
				jsre.MustCompile("^úno", "i"),
				jsre.MustCompile("^bře", "i"),
				jsre.MustCompile("^dub", "i"),
				jsre.MustCompile("^kvě", "i"),
				jsre.MustCompile("^(čvn|červen$|června)", "i"),
				jsre.MustCompile("^(čvc|červenec|července)", "i"),
				jsre.MustCompile("^srp", "i"),
				jsre.MustCompile("^zář", "i"),
				jsre.MustCompile("^říj", "i"),
				jsre.MustCompile("^lis", "i"),
				jsre.MustCompile("^pro", "i"),
			},
			"weekdays", []string{"neděle", "pondělí", "úterý", "středa", "čtvrtek", "pátek", "sobota"},
			"weekdaysShort", []string{"ne", "po", "út", "st", "čt", "pá", "so"},
			"weekdaysMin", []string{"ne", "po", "út", "st", "čt", "pá", "so"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY H:mm",
				"LLLL", "dddd D. MMMM YYYY H:mm",
				"l", "D. M. YYYY",
			),
			"calendar", js.Obj(
				"sameDay", "[dnes v] LT",
				"nextDay", "[zítra v] LT",
				"nextWeek", csCalendarNextWeek,
				"lastDay", "[včera v] LT",
				"lastWeek", csCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "za %s",
				"past", "před %s",
				"s", csRelativeTimeSeconds,
				"ss", csRelativeTimeSeconds,
				"m", csRelativeTimeSeconds,
				"mm", csRelativeTimeSeconds,
				"h", csRelativeTimeSeconds,
				"hh", csRelativeTimeSeconds,
				"d", csRelativeTimeSeconds,
				"dd", csRelativeTimeSeconds,
				"M", csRelativeTimeSeconds,
				"MM", csRelativeTimeSeconds,
				"y", csRelativeTimeSeconds,
				"yy", csRelativeTimeSeconds,
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

// csCalendarNextWeek is calendar.nextWeek of cs.
func csCalendarNextWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[v neděli v] LT"
	case 1, 2:
		return "[v] dddd [v] LT"
	case 3:
		return "[ve středu v] LT"
	case 4:
		return "[ve čtvrtek v] LT"
	case 5:
		return "[v pátek v] LT"
	case 6:
		return "[v sobotu v] LT"
	}
	return ""
}

// csCalendarLastWeek is calendar.lastWeek of cs.
func csCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[minulou neděli v] LT"
	case 1, 2:
		return "[minulé] dddd [v] LT"
	case 3:
		return "[minulou středu v] LT"
	case 4, 5:
		return "[minulý] dddd [v] LT"
	case 6:
		return "[minulou sobotu v] LT"
	}
	return ""
}

// csPlural is the Czech paucal test (2 to 4).
func csPlural(n float64) bool { return n > 1 && n < 5 && js.ToInt32(n/10) != 1 }

// csRelativeTimeSeconds is translate of cs.
func csRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	result := num(number) + " "
	counted := func(few, many, past string) string {
		if withoutSuffix || isFuture {
			if csPlural(number) {
				return result + few
			}
			return result + many
		}
		return result + past
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
	futureOr := func(a, past string) string {
		if withoutSuffix || isFuture {
			return a
		}
		return past
	}
	switch key {
	case "s":
		return futureOr("pár sekund", "pár sekundami")
	case "ss":
		return counted("sekundy", "sekund", "sekundami")
	case "m":
		return pick("minuta", "minutu", "minutou")
	case "mm":
		return counted("minuty", "minut", "minutami")
	case "h":
		return pick("hodina", "hodinu", "hodinou")
	case "hh":
		return counted("hodiny", "hodin", "hodinami")
	case "d":
		return futureOr("den", "dnem")
	case "dd":
		return counted("dny", "dní", "dny")
	case "M":
		return futureOr("měsíc", "měsícem")
	case "MM":
		return counted("měsíce", "měsíců", "měsíci")
	case "y":
		return futureOr("rok", "rokem")
	case "yy":
		return counted("roky", "let", "lety")
	}
	return ""
}
