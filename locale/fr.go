// moment.js locale configuration
// locale : French [fr]
// author : John Fischer : https://github.com/jfroffice

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("fr", func() {
		moment.DefineLocale("fr", js.Obj(
			"months", []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
			"monthsShort", []string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
			"monthsRegex", jsre.MustCompile("(janv\\.?|févr\\.?|mars|avr\\.?|mai|juin|juil\\.?|août|sept\\.?|oct\\.?|nov\\.?|déc\\.?|janvier|février|mars|avril|mai|juin|juillet|août|septembre|octobre|novembre|décembre)", "i"),
			"monthsShortRegex", jsre.MustCompile("(janv\\.?|févr\\.?|mars|avr\\.?|mai|juin|juil\\.?|août|sept\\.?|oct\\.?|nov\\.?|déc\\.?|janvier|février|mars|avril|mai|juin|juillet|août|septembre|octobre|novembre|décembre)", "i"),
			"monthsStrictRegex", jsre.MustCompile("^(janvier|février|mars|avril|mai|juin|juillet|août|septembre|octobre|novembre|décembre)", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("(janv\\.?|févr\\.?|mars|avr\\.?|mai|juin|juil\\.?|août|sept\\.?|oct\\.?|nov\\.?|déc\\.?)", "i"),
			"monthsParse", []*jsre.Regexp{
				jsre.MustCompile("^janv", "i"),
				jsre.MustCompile("^févr", "i"),
				jsre.MustCompile("^mars", "i"),
				jsre.MustCompile("^avr", "i"),
				jsre.MustCompile("^mai", "i"),
				jsre.MustCompile("^juin", "i"),
				jsre.MustCompile("^juil", "i"),
				jsre.MustCompile("^août", "i"),
				jsre.MustCompile("^sept", "i"),
				jsre.MustCompile("^oct", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^déc", "i"),
			},
			"longMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^janv", "i"),
				jsre.MustCompile("^févr", "i"),
				jsre.MustCompile("^mars", "i"),
				jsre.MustCompile("^avr", "i"),
				jsre.MustCompile("^mai", "i"),
				jsre.MustCompile("^juin", "i"),
				jsre.MustCompile("^juil", "i"),
				jsre.MustCompile("^août", "i"),
				jsre.MustCompile("^sept", "i"),
				jsre.MustCompile("^oct", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^déc", "i"),
			},
			"shortMonthsParse", []*jsre.Regexp{
				jsre.MustCompile("^janv", "i"),
				jsre.MustCompile("^févr", "i"),
				jsre.MustCompile("^mars", "i"),
				jsre.MustCompile("^avr", "i"),
				jsre.MustCompile("^mai", "i"),
				jsre.MustCompile("^juin", "i"),
				jsre.MustCompile("^juil", "i"),
				jsre.MustCompile("^août", "i"),
				jsre.MustCompile("^sept", "i"),
				jsre.MustCompile("^oct", "i"),
				jsre.MustCompile("^nov", "i"),
				jsre.MustCompile("^déc", "i"),
			},
			"weekdays", []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
			"weekdaysShort", []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
			"weekdaysMin", []string{"di", "lu", "ma", "me", "je", "ve", "sa"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Aujourd’hui à] LT",
				"nextDay", "[Demain à] LT",
				"nextWeek", "dddd [à] LT",
				"lastDay", "[Hier à] LT",
				"lastWeek", "dddd [dernier à] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "dans %s",
				"past", "il y a %s",
				"s", "quelques secondes",
				"ss", "%d secondes",
				"m", "une minute",
				"mm", "%d minutes",
				"h", "une heure",
				"hh", "%d heures",
				"d", "un jour",
				"dd", "%d jours",
				"w", "une semaine",
				"ww", "%d semaines",
				"M", "un mois",
				"MM", "%d mois",
				"y", "un an",
				"yy", "%d ans",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(er|)", ""),
			"ordinal", frOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// frOrdinal is the ordinal callback of fr: no suffix for days of the month after the first.
func frOrdinal(l *moment.Locale, number moment.Value, period string) moment.Value {
	if period == "D" {
		if sameNumber(number, 1) {
			return str(number) + "er"
		}
		return str(number)
	}
	return frCaOrdinal(l, number, period)
}
