// moment.js locale configuration
// locale : German [de]
// author : lluchs : https://github.com/lluchs
// author: Menelion Elensúle: https://github.com/Oire
// author : Mikolaj Dadela : https://github.com/mik01aj

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("de", func() {
		moment.DefineLocale("de", js.Obj(
			"months", []string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
			"monthsShort", []string{"Jan.", "Feb.", "März", "Apr.", "Mai", "Juni", "Juli", "Aug.", "Sep.", "Okt.", "Nov.", "Dez."},
			"monthsRegex", jsre.MustCompile("^(Januar|Februar|März|April|Mai|Juni|Juli|August|September|Oktober|November|Dezember|Jan\\.?|Feb\\.?|Mär\\.?|Apr\\.?|Jun\\.?|Jul\\.?|Aug\\.?|Sep\\.?|Okt\\.?|Nov\\.?|Dez\\.?)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(Januar|Februar|März|April|Mai|Juni|Juli|August|September|Oktober|November|Dezember|Jan\\.?|Feb\\.?|Mär\\.?|Apr\\.?|Jun\\.?|Jul\\.?|Aug\\.?|Sep\\.?|Okt\\.?|Nov\\.?|Dez\\.?)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(Januar|Februar|März|April|Mai|Juni|Juli|August|September|Oktober|November|Dezember)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(Jan\\.?|Feb\\.?|März|Mär\\.?|Apr\\.?|Mai|Juni|Jun\\.?|Juli|Jul\\.?|Aug\\.?|Sep\\.?|Okt\\.?|Nov\\.?|Dez\\.?)", "i"),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^jan", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^mär", "i"),
				jsre.MustCompile("^apr", "i"),
				jsre.MustCompile("^mai", "i"),
				jsre.MustCompile("^jun", "i"),
				jsre.MustCompile("^jul", "i"),
				jsre.MustCompile("^aug", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^okt", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dez", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^jan", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^mär", "i"),
				jsre.MustCompile("^apr", "i"),
				jsre.MustCompile("^mai", "i"),
				jsre.MustCompile("^jun", "i"),
				jsre.MustCompile("^jul", "i"),
				jsre.MustCompile("^aug", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^okt", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dez", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^jan", "i"),
				jsre.MustCompile("^feb", "i"),
				jsre.MustCompile("^mär", "i"),
				jsre.MustCompile("^apr", "i"),
				jsre.MustCompile("^mai", "i"),
				jsre.MustCompile("^jun", "i"),
				jsre.MustCompile("^jul", "i"),
				jsre.MustCompile("^aug", "i"),
				jsre.MustCompile("^sep", "i"),
				jsre.MustCompile("^okt", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^dez", "i"),
			},
			"weekdays", []string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
			"weekdaysShort", []string{"So.", "Mo.", "Di.", "Mi.", "Do.", "Fr.", "Sa."},
			"weekdaysMin", []string{"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY HH:mm",
				"LLLL", "dddd, D. MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[heute um] LT [Uhr]",
				"sameElse", "L",
				"nextDay", "[morgen um] LT [Uhr]",
				"nextWeek", "dddd [um] LT [Uhr]",
				"lastDay", "[gestern um] LT [Uhr]",
				"lastWeek", "[letzten] dddd [um] LT [Uhr]",
			),
			"relativeTime", js.Obj(
				"future", "in %s",
				"past", "vor %s",
				"s", "ein paar Sekunden",
				"ss", "%d Sekunden",
				"m", deRelativeTimeMinute,
				"mm", "%d Minuten",
				"h", deRelativeTimeMinute,
				"hh", "%d Stunden",
				"d", deRelativeTimeMinute,
				"dd", deRelativeTimeMinute,
				"w", deRelativeTimeMinute,
				"ww", "%d Wochen",
				"M", deRelativeTimeMinute,
				"MM", deRelativeTimeMinute,
				"y", deRelativeTimeMinute,
				"yy", deRelativeTimeMinute,
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

var deRelativeTimeMinute = germanRelativeTime

// germanRelativeTime is processRelativeTime of de, de-at and de-ch.
func germanRelativeTime(number float64, withoutSuffix bool, key string, _ bool) string {
	n := num(number)
	format := map[string][2]string{
		"m": {"eine Minute", "einer Minute"}, "h": {"eine Stunde", "einer Stunde"}, "d": {"ein Tag", "einem Tag"},
		"dd": {n + " Tage", n + " Tagen"}, "w": {"eine Woche", "einer Woche"}, "M": {"ein Monat", "einem Monat"},
		"MM": {n + " Monate", n + " Monaten"}, "y": {"ein Jahr", "einem Jahr"}, "yy": {n + " Jahre", n + " Jahren"},
	}
	if withoutSuffix {
		return format[key][0]
	}
	return format[key][1]
}
