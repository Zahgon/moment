// moment.js locale configuration
// locale : Romanian [ro]
// author : Vlad Gurdiga : https://github.com/gurdiga
// author : Valentin Agachi : https://github.com/avaly
// author : Emanuel Cepoi : https://github.com/cepem

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("ro", func() {
		moment.DefineLocale("ro", js.Obj(
			"months", []string{"ianuarie", "februarie", "martie", "aprilie", "mai", "iunie", "iulie", "august", "septembrie", "octombrie", "noiembrie", "decembrie"},
			"monthsShort", []string{"ian.", "feb.", "mart.", "apr.", "mai", "iun.", "iul.", "aug.", "sept.", "oct.", "nov.", "dec."},
			"monthsParseExact", true,
			"weekdays", []string{"duminică", "luni", "marți", "miercuri", "joi", "vineri", "sâmbătă"},
			"weekdaysShort", []string{"Dum", "Lun", "Mar", "Mie", "Joi", "Vin", "Sâm"},
			"weekdaysMin", []string{"Du", "Lu", "Ma", "Mi", "Jo", "Vi", "Sâ"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY H:mm",
				"LLLL", "dddd, D MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[azi la] LT",
				"nextDay", "[mâine la] LT",
				"nextWeek", "dddd [la] LT",
				"lastDay", "[ieri la] LT",
				"lastWeek", "[fosta] dddd [la] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "peste %s",
				"past", "%s în urmă",
				"s", "câteva secunde",
				"ss", roRelativeTimeSecondsN,
				"m", "un minut",
				"mm", roRelativeTimeSecondsN,
				"h", "o oră",
				"hh", roRelativeTimeSecondsN,
				"d", "o zi",
				"dd", roRelativeTimeSecondsN,
				"w", "o săptămână",
				"ww", roRelativeTimeSecondsN,
				"M", "o lună",
				"MM", roRelativeTimeSecondsN,
				"y", "un an",
				"yy", roRelativeTimeSecondsN,
			),
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var roUnits = map[string]string{"ss": "secunde", "mm": "minute", "hh": "ore", "dd": "zile", "ww": "săptămâni",
	"MM": "luni", "yy": "ani"}

// roRelativeTimeSecondsN is relativeTimeWithPlural of ro: "de" before counts from 20.
func roRelativeTimeSecondsN(number float64, _ bool, key string, _ bool) string {
	separator := " "
	if js.Mod(number, 100) >= 20 || (number >= 100 && js.Mod(number, 100) == 0) {
		separator = " de "
	}
	return num(number) + separator + roUnits[key]
}
