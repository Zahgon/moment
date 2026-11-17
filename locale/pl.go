// moment.js locale configuration
// locale : Polish [pl]
// author : Rafal Hirsz : https://github.com/evoL

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("pl", func() {
		moment.DefineLocale("pl", js.Obj(
			"months", plMonths,
			"monthsShort", []string{"sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^sty", "i"),
				jsre.MustCompile("^lut", "i"),
				jsre.MustCompile("^mar", "i"),
				jsre.MustCompile("^kwi", "i"),
				jsre.MustCompile("^maj", "i"),
				jsre.MustCompile("^cze", "i"),
				jsre.MustCompile("^lip", "i"),
				jsre.MustCompile("^sie", "i"),
				jsre.MustCompile("^wrz", "i"),
				jsre.MustCompile("^paź", "i"),
				jsre.MustCompile("^lis", "i"),
				jsre.MustCompile("^gru", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^sty", "i"),
				jsre.MustCompile("^lut", "i"),
				jsre.MustCompile("^mar", "i"),
				jsre.MustCompile("^kwi", "i"),
				jsre.MustCompile("^maj", "i"),
				jsre.MustCompile("^cze", "i"),
				jsre.MustCompile("^lip", "i"),
				jsre.MustCompile("^sie", "i"),
				jsre.MustCompile("^wrz", "i"),
				jsre.MustCompile("^paź", "i"),
				jsre.MustCompile("^lis", "i"),
				jsre.MustCompile("^gru", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^sty", "i"),
				jsre.MustCompile("^lut", "i"),
				jsre.MustCompile("^mar", "i"),
				jsre.MustCompile("^kwi", "i"),
				jsre.MustCompile("^maj", "i"),
				jsre.MustCompile("^cze", "i"),
				jsre.MustCompile("^lip", "i"),
				jsre.MustCompile("^sie", "i"),
				jsre.MustCompile("^wrz", "i"),
				jsre.MustCompile("^paź", "i"),
				jsre.MustCompile("^lis", "i"),
				jsre.MustCompile("^gru", "i"),
			},
			"weekdays", []string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
			"weekdaysShort", []string{"ndz", "pon", "wt", "śr", "czw", "pt", "sob"},
			"weekdaysMin", []string{"Nd", "Pn", "Wt", "Śr", "Cz", "Pt", "So"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Dziś o] LT",
				"nextDay", "[Jutro o] LT",
				"nextWeek", plCalendarNextWeek,
				"lastDay", "[Wczoraj o] LT",
				"lastWeek", plCalendarLastWeek,
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "za %s",
				"past", "%s temu",
				"s", "kilka sekund",
				"ss", plRelativeTimeSecondsN,
				"m", plRelativeTimeSecondsN,
				"mm", plRelativeTimeSecondsN,
				"h", plRelativeTimeSecondsN,
				"hh", plRelativeTimeSecondsN,
				"d", "1 dzień",
				"dd", "%d dni",
				"w", "tydzień",
				"ww", plRelativeTimeSecondsN,
				"M", "miesiąc",
				"MM", plRelativeTimeSecondsN,
				"y", "rok",
				"yy", plRelativeTimeSecondsN,
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

var (
	plMonthsNominative = []string{"styczeń", "luty", "marzec", "kwiecień", "maj", "czerwiec", "lipiec", "sierpień",
		"wrzesień", "październik", "listopad", "grudzień"}
	plMonthsSubjective = []string{"stycznia", "lutego", "marca", "kwietnia", "maja", "czerwca", "lipca", "sierpnia",
		"września", "października", "listopada", "grudnia"}
	plDayBeforeMonth = jsre.MustCompile(`D\.? MMMM`, "")
)

// plMonths returns the genitive month name after a day number.
func plMonths(_ *moment.Locale, m *moment.Moment, format any) any {
	if m == nil {
		return plMonthsNominative
	}
	list := plMonthsNominative
	if plDayBeforeMonth.Test(js.ToString(format)) {
		list = plMonthsSubjective
	}
	if i := m.Month(); i >= 0 && int(i) < len(list) {
		return list[int(i)]
	}
	return js.Undef
}

// plCalendarNextWeek is calendar.nextWeek of pl.
func plCalendarNextWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[W niedzielę o] LT"
	case 2:
		return "[We wtorek o] LT"
	case 3:
		return "[W środę o] LT"
	case 6:
		return "[W sobotę o] LT"
	}
	return "[W] dddd [o] LT"
}

// plCalendarLastWeek is calendar.lastWeek of pl.
func plCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 0:
		return "[W zeszłą niedzielę o] LT"
	case 3:
		return "[W zeszłą środę o] LT"
	case 6:
		return "[W zeszłą sobotę o] LT"
	}
	return "[W zeszły] dddd [o] LT"
}

// plPlural is the Polish paucal test.
func plPlural(n float64) bool {
	return js.Mod(n, 10) < 5 && js.Mod(n, 10) > 1 && js.Mod(js.ToInt32(n/10), 10) != 1
}

// plRelativeTimeSecondsN is translate of pl.
func plRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, _ bool) string {
	result := num(number) + " "
	counted := func(few, many string) string {
		if plPlural(number) {
			return result + few
		}
		return result + many
	}
	switch key {
	case "ss":
		return counted("sekundy", "sekund")
	case "m":
		if withoutSuffix {
			return "minuta"
		}
		return "minutę"
	case "mm":
		return counted("minuty", "minut")
	case "h":
		if withoutSuffix {
			return "godzina"
		}
		return "godzinę"
	case "hh":
		return counted("godziny", "godzin")
	case "ww":
		return counted("tygodnie", "tygodni")
	case "MM":
		return counted("miesiące", "miesięcy")
	case "yy":
		return counted("lata", "lat")
	}
	return ""
}
