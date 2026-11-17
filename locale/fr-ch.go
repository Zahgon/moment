// moment.js locale configuration
// locale : French (Switzerland) [fr-ch]
// author : Gaspard Bucher : https://github.com/gaspard

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("fr-ch", func() {
		moment.DefineLocale("fr-ch", js.Obj(
			"months", []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
			"monthsShort", []string{"janv.", "févr.", "mars", "avr.", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
			"monthsParseExact", true,
			"weekdays", []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
			"weekdaysShort", []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
			"weekdaysMin", []string{"di", "lu", "ma", "me", "je", "ve", "sa"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
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
				"M", "un mois",
				"MM", "%d mois",
				"y", "un an",
				"yy", "%d ans",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(er|e)", ""),
			"ordinal", frChOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var frChOrdinal = frCaOrdinal
