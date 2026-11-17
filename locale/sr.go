// moment.js locale configuration
// locale : Serbian [sr]
// author : Milan Janačković<milanjanackovic@gmail.com> : https://github.com/milan-j
// author : Stefan Crnjaković <stefan@hotmail.rs> : https://github.com/crnjakovic

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("sr", func() {
		moment.DefineLocale("sr", js.Obj(
			"months", []string{"januar", "februar", "mart", "april", "maj", "jun", "jul", "avgust", "septembar", "oktobar", "novembar", "decembar"},
			"monthsShort", []string{"jan.", "feb.", "mar.", "apr.", "maj", "jun", "jul", "avg.", "sep.", "okt.", "nov.", "dec."},
			"monthsParseExact", true,
			"weekdays", []string{"nedelja", "ponedeljak", "utorak", "sreda", "četvrtak", "petak", "subota"},
			"weekdaysShort", []string{"ned.", "pon.", "uto.", "sre.", "čet.", "pet.", "sub."},
			"weekdaysMin", []string{"ne", "po", "ut", "sr", "če", "pe", "su"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "D. M. YYYY.",
				"LL", "D. MMMM YYYY.",
				"LLL", "D. MMMM YYYY. H:mm",
				"LLLL", "dddd, D. MMMM YYYY. H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[danas u] LT",
				"nextDay", "[sutra u] LT",
				"nextWeek", srCalendarNextWeek,
				"lastDay", "[juče u] LT",
				"lastWeek", srCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "za %s",
				"past", "pre %s",
				"s", "nekoliko sekundi",
				"ss", srRelativeTimeSecondsN,
				"m", srRelativeTimeSecondsN,
				"mm", srRelativeTimeSecondsN,
				"h", srRelativeTimeSecondsN,
				"hh", srRelativeTimeSecondsN,
				"d", srRelativeTimeSecondsN,
				"dd", srRelativeTimeSecondsN,
				"M", srRelativeTimeSecondsN,
				"MM", srRelativeTimeSecondsN,
				"y", srRelativeTimeSecondsN,
				"yy", srRelativeTimeSecondsN,
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

var srWords = map[string][]string{
	"ss": {"sekunda", "sekunde", "sekundi"}, "m": {"jedan minut", "jednog minuta"}, "mm": {"minut", "minuta", "minuta"},
	"h": {"jedan sat", "jednog sata"}, "hh": {"sat", "sata", "sati"}, "d": {"jedan dan", "jednog dana"},
	"dd": {"dan", "dana", "dana"}, "M": {"jedan mesec", "jednog meseca"}, "MM": {"mesec", "meseca", "meseci"},
	"y": {"jednu godinu", "jedne godine"}, "yy": {"godinu", "godine", "godina"},
}

var srLastWeekDays = []string{"[prošle] [nedelje] [u] LT", "[prošlog] [ponedeljka] [u] LT",
	"[prošlog] [utorka] [u] LT", "[prošle] [srede] [u] LT", "[prošlog] [četvrtka] [u] LT",
	"[prošlog] [petka] [u] LT", "[prošle] [subote] [u] LT"}

// srCalendarNextWeek is calendar.nextWeek of sr.
func srCalendarNextWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[u] [nedelju] [u] LT"
	case 3:
		return "[u] [sredu] [u] LT"
	case 6:
		return "[u] [subotu] [u] LT"
	case 1, 2, 4, 5:
		return "[u] dddd [u] LT"
	}
	return ""
}

// srCalendarLastWeek is calendar.lastWeek of sr.
func srCalendarLastWeek(m, _ *moment.Moment) string { return serbianLastWeek(m, srLastWeekDays) }

// serbianLastWeek picks the last-week calendar text by weekday.
func serbianLastWeek(m *moment.Moment, days []string) string {
	if d := m.Day(); d >= 0 && d < 7 {
		return days[int(d)]
	}
	return ""
}

// serbianCase picks the grammatical case of a counted word.
func serbianCase(number float64, words []string) string {
	if d := js.Mod(number, 10); d >= 1 && d <= 4 && (js.Mod(number, 100) < 10 || js.Mod(number, 100) >= 20) {
		if d == 1 {
			return words[0]
		}
		return words[1]
	}
	return words[2]
}

// serbianTranslate is translator.translate of sr and sr-cyrl.
func serbianTranslate(words map[string][]string, oneYear, yearAccusative, yearsNominative string, number float64,
	withoutSuffix bool, key string, isFuture bool) string {
	forms := words[key]
	if js.Len16(key) == 1 {
		if key == "y" && withoutSuffix {
			return oneYear
		}
		if isFuture || withoutSuffix {
			return forms[0]
		}
		return forms[1]
	}
	word := serbianCase(number, forms)
	if key == "yy" && withoutSuffix && word == yearAccusative {
		return num(number) + yearsNominative
	}
	return num(number) + " " + word
}

// srRelativeTimeSecondsN is translator.translate of sr.
func srRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, isFuture bool) string {
	return serbianTranslate(srWords, "jedna godina", "godinu", " godina", number, withoutSuffix, key, isFuture)
}
