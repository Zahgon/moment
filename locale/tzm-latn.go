// moment.js locale configuration
// locale : Central Atlas Tamazight Latin [tzm-latn]
// author : Abdel Said : https://github.com/abdelsaid

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("tzm-latn", func() {
		moment.DefineLocale("tzm-latn", js.Obj(
			"months", []string{"innayr", "brˤayrˤ", "marˤsˤ", "ibrir", "mayyw", "ywnyw", "ywlywz", "ɣwšt", "šwtanbir", "ktˤwbrˤ", "nwwanbir", "dwjnbir"},
			"monthsShort", []string{"innayr", "brˤayrˤ", "marˤsˤ", "ibrir", "mayyw", "ywnyw", "ywlywz", "ɣwšt", "šwtanbir", "ktˤwbrˤ", "nwwanbir", "dwjnbir"},
			"weekdays", []string{"asamas", "aynas", "asinas", "akras", "akwas", "asimwas", "asiḍyas"},
			"weekdaysShort", []string{"asamas", "aynas", "asinas", "akras", "akwas", "asimwas", "asiḍyas"},
			"weekdaysMin", []string{"asamas", "aynas", "asinas", "akras", "akwas", "asimwas", "asiḍyas"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[asdkh g] LT",
				"nextDay", "[aska g] LT",
				"nextWeek", "dddd [g] LT",
				"lastDay", "[assant g] LT",
				"lastWeek", "dddd [g] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "dadkh s yan %s",
				"past", "yan %s",
				"s", "imik",
				"ss", "%d imik",
				"m", "minuḍ",
				"mm", "%d minuḍ",
				"h", "saɛa",
				"hh", "%d tassaɛin",
				"d", "ass",
				"dd", "%d ossan",
				"M", "ayowr",
				"MM", "%d iyyirn",
				"y", "asgas",
				"yy", "%d isgasn",
			),
			"week", js.Obj(
				"dow", 6.0,
				"doy", 12.0,
			),
		))
	})
}
