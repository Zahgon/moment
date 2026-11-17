// moment.js locale configuration
// locale : Central Atlas Tamazight [tzm]
// author : Abdel Said : https://github.com/abdelsaid

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("tzm", func() {
		moment.DefineLocale("tzm", js.Obj(
			"months", []string{"ⵉⵏⵏⴰⵢⵔ", "ⴱⵕⴰⵢⵕ", "ⵎⴰⵕⵚ", "ⵉⴱⵔⵉⵔ", "ⵎⴰⵢⵢⵓ", "ⵢⵓⵏⵢⵓ", "ⵢⵓⵍⵢⵓⵣ", "ⵖⵓⵛⵜ", "ⵛⵓⵜⴰⵏⴱⵉⵔ", "ⴽⵟⵓⴱⵕ", "ⵏⵓⵡⴰⵏⴱⵉⵔ", "ⴷⵓⵊⵏⴱⵉⵔ"},
			"monthsShort", []string{"ⵉⵏⵏⴰⵢⵔ", "ⴱⵕⴰⵢⵕ", "ⵎⴰⵕⵚ", "ⵉⴱⵔⵉⵔ", "ⵎⴰⵢⵢⵓ", "ⵢⵓⵏⵢⵓ", "ⵢⵓⵍⵢⵓⵣ", "ⵖⵓⵛⵜ", "ⵛⵓⵜⴰⵏⴱⵉⵔ", "ⴽⵟⵓⴱⵕ", "ⵏⵓⵡⴰⵏⴱⵉⵔ", "ⴷⵓⵊⵏⴱⵉⵔ"},
			"weekdays", []string{"ⴰⵙⴰⵎⴰⵙ", "ⴰⵢⵏⴰⵙ", "ⴰⵙⵉⵏⴰⵙ", "ⴰⴽⵔⴰⵙ", "ⴰⴽⵡⴰⵙ", "ⴰⵙⵉⵎⵡⴰⵙ", "ⴰⵙⵉⴹⵢⴰⵙ"},
			"weekdaysShort", []string{"ⴰⵙⴰⵎⴰⵙ", "ⴰⵢⵏⴰⵙ", "ⴰⵙⵉⵏⴰⵙ", "ⴰⴽⵔⴰⵙ", "ⴰⴽⵡⴰⵙ", "ⴰⵙⵉⵎⵡⴰⵙ", "ⴰⵙⵉⴹⵢⴰⵙ"},
			"weekdaysMin", []string{"ⴰⵙⴰⵎⴰⵙ", "ⴰⵢⵏⴰⵙ", "ⴰⵙⵉⵏⴰⵙ", "ⴰⴽⵔⴰⵙ", "ⴰⴽⵡⴰⵙ", "ⴰⵙⵉⵎⵡⴰⵙ", "ⴰⵙⵉⴹⵢⴰⵙ"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[ⴰⵙⴷⵅ ⴴ] LT",
				"nextDay", "[ⴰⵙⴽⴰ ⴴ] LT",
				"nextWeek", "dddd [ⴴ] LT",
				"lastDay", "[ⴰⵚⴰⵏⵜ ⴴ] LT",
				"lastWeek", "dddd [ⴴ] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "ⴷⴰⴷⵅ ⵙ ⵢⴰⵏ %s",
				"past", "ⵢⴰⵏ %s",
				"s", "ⵉⵎⵉⴽ",
				"ss", "%d ⵉⵎⵉⴽ",
				"m", "ⵎⵉⵏⵓⴺ",
				"mm", "%d ⵎⵉⵏⵓⴺ",
				"h", "ⵙⴰⵄⴰ",
				"hh", "%d ⵜⴰⵙⵙⴰⵄⵉⵏ",
				"d", "ⴰⵙⵙ",
				"dd", "%d oⵙⵙⴰⵏ",
				"M", "ⴰⵢoⵓⵔ",
				"MM", "%d ⵉⵢⵢⵉⵔⵏ",
				"y", "ⴰⵙⴳⴰⵙ",
				"yy", "%d ⵉⵙⴳⴰⵙⵏ",
			),
			"week", js.Obj(
				"dow", 6.0,
				"doy", 12.0,
			),
		))
	})
}
