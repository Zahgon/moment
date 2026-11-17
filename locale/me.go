// moment.js locale configuration
// locale : Montenegrin [me]
// author : Miodrag Nikač <miodrag@restartit.me> : https://github.com/miodragnikac

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("me", func() {
		moment.DefineLocale("me", js.Obj(
			"months", []string{"januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"},
			"monthsShort", []string{"jan.", "feb.", "mar.", "apr.", "maj", "jun", "jul", "avg.", "sep.", "okt.", "nov.", "dec."},
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
				"nextDay", "[sjutra u] LT",
				"nextWeek", meCalendarNextWeek,
				"lastDay", "[juče u] LT",
				"lastWeek", meCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "za %s",
				"past", "prije %s",
				"s", "nekoliko sekundi",
				"ss", meRelativeTimeSecondsN,
				"m", meRelativeTimeSecondsN,
				"mm", meRelativeTimeSecondsN,
				"h", meRelativeTimeSecondsN,
				"hh", meRelativeTimeSecondsN,
				"d", "dan",
				"dd", meRelativeTimeSecondsN,
				"M", "mjesec",
				"MM", meRelativeTimeSecondsN,
				"y", "godinu",
				"yy", meRelativeTimeSecondsN,
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

var meWords = map[string][]string{
	"ss": {"sekund", "sekunda", "sekundi"}, "m": {"jedan minut", "jednog minuta"}, "mm": {"minut", "minuta", "minuta"},
	"h": {"jedan sat", "jednog sata"}, "hh": {"sat", "sata", "sati"}, "dd": {"dan", "dana", "dana"},
	"MM": {"mjesec", "mjeseca", "mjeseci"}, "yy": {"godina", "godine", "godina"},
}

// meCalendarNextWeek is calendar.nextWeek of me.
func meCalendarNextWeek(m, _ *moment.Moment) string { return bsCalendarNextWeek(m, nil) }

var meLastWeekDays = []string{"[prošle] [nedjelje] [u] LT", "[prošlog] [ponedjeljka] [u] LT",
	"[prošlog] [utorka] [u] LT", "[prošle] [srijede] [u] LT", "[prošlog] [četvrtka] [u] LT",
	"[prošlog] [petka] [u] LT", "[prošle] [subote] [u] LT"}

// meCalendarLastWeek is calendar.lastWeek of me.
func meCalendarLastWeek(m, _ *moment.Moment) string {
	if d := m.Day(); d >= 0 && d < 7 {
		return meLastWeekDays[int(d)]
	}
	return ""
}

// meCase picks the grammatical case of a counted word: singular, paucal (2 to 4) or plural.
func meCase(number float64, words []string) string {
	switch {
	case number == 1:
		return words[0]
	case number >= 2 && number <= 4:
		return words[1]
	}
	return words[2]
}

// meRelativeTimeSecondsN is translator.translate of me.
func meRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, _ bool) string {
	words := meWords[key]
	if js.Len16(key) == 1 {
		if withoutSuffix {
			return words[0]
		}
		return words[1]
	}
	return num(number) + " " + meCase(number, words)
}
