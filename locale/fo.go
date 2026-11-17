// moment.js locale configuration
// locale : Faroese [fo]
// author : Ragnar Johannesen : https://github.com/ragnar123
// author : Kristian Sakarisson : https://github.com/sakarisson

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("fo", func() {
		moment.DefineLocale("fo", js.Obj(
			"months", []string{"januar", "februar", "mars", "apríl", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
			"monthsShort", []string{"jan", "feb", "mar", "apr", "mai", "jun", "jul", "aug", "sep", "okt", "nov", "des"},
			"weekdays", []string{"sunnudagur", "mánadagur", "týsdagur", "mikudagur", "hósdagur", "fríggjadagur", "leygardagur"},
			"weekdaysShort", []string{"sun", "mán", "týs", "mik", "hós", "frí", "ley"},
			"weekdaysMin", []string{"su", "má", "tý", "mi", "hó", "fr", "le"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd D. MMMM, YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Í dag kl.] LT",
				"nextDay", "[Í morgin kl.] LT",
				"nextWeek", "dddd [kl.] LT",
				"lastDay", "[Í gjár kl.] LT",
				"lastWeek", "[síðstu] dddd [kl] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "um %s",
				"past", "%s síðani",
				"s", "fá sekund",
				"ss", "%d sekundir",
				"m", "ein minuttur",
				"mm", "%d minuttir",
				"h", "ein tími",
				"hh", "%d tímar",
				"d", "ein dagur",
				"dd", "%d dagar",
				"M", "ein mánaður",
				"MM", "%d mánaðir",
				"y", "eitt ár",
				"yy", "%d ár",
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
