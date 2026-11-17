// moment.js locale configuration
// locale : Nynorsk [nn]
// authors : https://github.com/mechuwind
//           Stephen Ramthun : https://github.com/stephenramthun

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("nn", func() {
		moment.DefineLocale("nn", js.Obj(
			"months", []string{"januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
			"monthsShort", []string{"jan.", "feb.", "mars", "apr.", "mai", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "des."},
			"monthsParseExact", true,
			"weekdays", []string{"sundag", "måndag", "tysdag", "onsdag", "torsdag", "fredag", "laurdag"},
			"weekdaysShort", []string{"su.", "må.", "ty.", "on.", "to.", "fr.", "lau."},
			"weekdaysMin", []string{"su", "må", "ty", "on", "to", "fr", "la"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY [kl.] H:mm",
				"LLLL", "dddd D. MMMM YYYY [kl.] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[I dag klokka] LT",
				"nextDay", "[I morgon klokka] LT",
				"nextWeek", "dddd [klokka] LT",
				"lastDay", "[I går klokka] LT",
				"lastWeek", "[Føregåande] dddd [klokka] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "om %s",
				"past", "%s sidan",
				"s", "nokre sekund",
				"ss", "%d sekund",
				"m", "eit minutt",
				"mm", "%d minutt",
				"h", "ein time",
				"hh", "%d timar",
				"d", "ein dag",
				"dd", "%d dagar",
				"w", "ei veke",
				"ww", "%d veker",
				"M", "ein månad",
				"MM", "%d månader",
				"y", "eit år",
				"yy", "%d år",
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
