// moment.js locale configuration
// locale : German (Austria) [de-at]
// author : lluchs : https://github.com/lluchs
// author: Menelion Elensúle: https://github.com/Oire
// author : Martin Groller : https://github.com/MadMG
// author : Mikolaj Dadela : https://github.com/mik01aj

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("de-at", func() {
		moment.DefineLocale("de-at", js.Obj(
			"months", []string{"Jänner", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
			"monthsShort", []string{"Jän.", "Feb.", "März", "Apr.", "Mai", "Juni", "Juli", "Aug.", "Sep.", "Okt.", "Nov.", "Dez."},
			"monthsRegex", jsre.MustCompile("^(Jänner|Februar|März|April|Mai|Juni|Juli|August|September|Oktober|November|Dezember|Jän\\.?|Feb\\.?|Mär\\.?|Apr\\.?|Jun\\.?|Jul\\.?|Aug\\.?|Sep\\.?|Okt\\.?|Nov\\.?|Dez\\.?)", "i"),
			"monthsShortRegex", jsre.MustCompile("^(Jänner|Februar|März|April|Mai|Juni|Juli|August|September|Oktober|November|Dezember|Jän\\.?|Feb\\.?|Mär\\.?|Apr\\.?|Jun\\.?|Jul\\.?|Aug\\.?|Sep\\.?|Okt\\.?|Nov\\.?|Dez\\.?)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(Jänner|Februar|März|April|Mai|Juni|Juli|August|September|Oktober|November|Dezember)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("^(Jän\\.?|Feb\\.?|März|Mär\\.?|Apr\\.?|Mai|Juni|Jun\\.?|Juli|Jul\\.?|Aug\\.?|Sep\\.?|Okt\\.?|Nov\\.?|Dez\\.?)", "i"),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^jän", "i"),
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
				jsre.MustCompile("^jän", "i"),
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
				jsre.MustCompile("^jän", "i"),
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
				"m", deAtRelativeTimeMinute,
				"mm", "%d Minuten",
				"h", deAtRelativeTimeMinute,
				"hh", "%d Stunden",
				"d", deAtRelativeTimeMinute,
				"dd", deAtRelativeTimeMinute,
				"w", deAtRelativeTimeMinute,
				"ww", "%d Wochen",
				"M", deAtRelativeTimeMinute,
				"MM", deAtRelativeTimeMinute,
				"y", deAtRelativeTimeMinute,
				"yy", deAtRelativeTimeMinute,
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

var deAtRelativeTimeMinute = germanRelativeTime
