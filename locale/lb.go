// moment.js locale configuration
// locale : Luxembourgish [lb]
// author : mweimerskirch : https://github.com/mweimerskirch
// author : David Raison : https://github.com/kwisatz

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("lb", func() {
		moment.DefineLocale("lb", js.Obj(
			"months", []string{"Januar", "Februar", "Mäerz", "Abrëll", "Mee", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
			"monthsShort", []string{"Jan.", "Febr.", "Mrz.", "Abr.", "Mee", "Jun.", "Jul.", "Aug.", "Sept.", "Okt.", "Nov.", "Dez."},
			"monthsParseExact", true,
			"weekdays", []string{"Sonndeg", "Méindeg", "Dënschdeg", "Mëttwoch", "Donneschdeg", "Freideg", "Samschdeg"},
			"weekdaysShort", []string{"So.", "Mé.", "Dë.", "Më.", "Do.", "Fr.", "Sa."},
			"weekdaysMin", []string{"So", "Mé", "Dë", "Më", "Do", "Fr", "Sa"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm [Auer]",
				"LTS", "H:mm:ss [Auer]",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY H:mm [Auer]",
				"LLLL", "dddd, D. MMMM YYYY H:mm [Auer]",
			),
			"calendar", js.Obj(
				"sameDay", "[Haut um] LT",
				"sameElse", "L",
				"nextDay", "[Muer um] LT",
				"nextWeek", "dddd [um] LT",
				"lastDay", "[Gëschter um] LT",
				"lastWeek", lbCalendarLastWeek,
			),
			"relativeTime", js.Obj(
				"future", lbRelativeTimeFuture,
				"past", lbRelativeTimePast,
				"s", "e puer Sekonnen",
				"ss", "%d Sekonnen",
				"m", lbRelativeTimeMinute,
				"mm", "%d Minutten",
				"h", lbRelativeTimeMinute,
				"hh", "%d Stonnen",
				"d", lbRelativeTimeMinute,
				"dd", "%d Deeg",
				"M", lbRelativeTimeMinute,
				"MM", "%d Méint",
				"y", lbRelativeTimeMinute,
				"yy", "%d Joer",
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

// lbEifeler reports whether the word before number loses its final n (the Eifeler rule).
func lbEifeler(number float64) bool {
	number = js.ParseInt(number)
	switch {
	case number != number:
		return false
	case number < 0:
		return true
	case number < 10:
		return 4 <= number && number <= 7
	case number < 100:
		lastDigit := js.Mod(number, 10)
		if lastDigit == 0 {
			return lbEifeler(number / 10)
		}
		return lbEifeler(lastDigit)
	case number < 10000:
		for number >= 10 {
			number /= 10
		}
		return lbEifeler(number)
	}
	return lbEifeler(number / 1000)
}

// lbLeadingNumber is the text before the first space as parseInt reads it.
func lbLeadingNumber(s string) float64 {
	return js.ParseInt(js.Substr16(s, 0, js.IndexOf16(s, " ")))
}

// lbCalendarLastWeek is calendar.lastWeek of lb.
func lbCalendarLastWeek(m, _ *moment.Moment) string {
	switch m.Day() {
	case 2, 4:
		return "[Leschten] dddd [um] LT"
	}
	return "[Leschte] dddd [um] LT"
}

// lbRelativeTimeFuture is processFutureTime of lb.
func lbRelativeTimeFuture(s string) string {
	if lbEifeler(lbLeadingNumber(s)) {
		return "a " + s
	}
	return "an " + s
}

// lbRelativeTimePast is processPastTime of lb.
func lbRelativeTimePast(s string) string {
	if lbEifeler(lbLeadingNumber(s)) {
		return "viru " + s
	}
	return "virun " + s
}

// lbRelativeTimeMinute is processRelativeTime of lb.
func lbRelativeTimeMinute(_ float64, withoutSuffix bool, key string, _ bool) string {
	format := map[string][2]string{"m": {"eng Minutt", "enger Minutt"}, "h": {"eng Stonn", "enger Stonn"},
		"d": {"een Dag", "engem Dag"}, "M": {"ee Mount", "engem Mount"}, "y": {"ee Joer", "engem Joer"}}
	if withoutSuffix {
		return format[key][0]
	}
	return format[key][1]
}
