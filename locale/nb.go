// moment.js locale configuration
// locale : Norwegian Bokmål [nb]
// authors : Espen Hovlandsdal : https://github.com/rexxars
//           Sigurd Gartmann : https://github.com/sigurdga
//           Stephen Ramthun : https://github.com/stephenramthun

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("nb", func() {
		moment.DefineLocale("nb", js.Obj(
			"months", []string{"januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
			"monthsShort", []string{"jan.", "feb.", "mars", "apr.", "mai", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "des."},
			"monthsParseExact", true,
			"weekdays", []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"},
			"weekdaysShort", []string{"sø.", "ma.", "ti.", "on.", "to.", "fr.", "lø."},
			"weekdaysMin", []string{"sø", "ma", "ti", "on", "to", "fr", "lø"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY [kl.] HH:mm",
				"LLLL", "dddd D. MMMM YYYY [kl.] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[i dag kl.] LT",
				"nextDay", "[i morgen kl.] LT",
				"nextWeek", "dddd [kl.] LT",
				"lastDay", "[i går kl.] LT",
				"lastWeek", "[forrige] dddd [kl.] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "om %s",
				"past", "%s siden",
				"s", "noen sekunder",
				"ss", "%d sekunder",
				"m", "ett minutt",
				"mm", "%d minutter",
				"h", "én time",
				"hh", "%d timer",
				"d", "én dag",
				"dd", "%d dager",
				"w", "én uke",
				"ww", "%d uker",
				"M", "én måned",
				"MM", "%d måneder",
				"y", "ett år",
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
