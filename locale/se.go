// moment.js locale configuration
// locale : Northern Sami [se]
// authors : Bård Rolstad Henriksen : https://github.com/karamell

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("se", func() {
		moment.DefineLocale("se", js.Obj(
			"months", []string{"ođđajagemánnu", "guovvamánnu", "njukčamánnu", "cuoŋománnu", "miessemánnu", "geassemánnu", "suoidnemánnu", "borgemánnu", "čakčamánnu", "golggotmánnu", "skábmamánnu", "juovlamánnu"},
			"monthsShort", []string{"ođđj", "guov", "njuk", "cuo", "mies", "geas", "suoi", "borg", "čakč", "golg", "skáb", "juov"},
			"weekdays", []string{"sotnabeaivi", "vuossárga", "maŋŋebárga", "gaskavahkku", "duorastat", "bearjadat", "lávvardat"},
			"weekdaysShort", []string{"sotn", "vuos", "maŋ", "gask", "duor", "bear", "láv"},
			"weekdaysMin", []string{"s", "v", "m", "g", "d", "b", "L"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "MMMM D. [b.] YYYY",
				"LLL", "MMMM D. [b.] YYYY [ti.] HH:mm",
				"LLLL", "dddd, MMMM D. [b.] YYYY [ti.] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[otne ti] LT",
				"nextDay", "[ihttin ti] LT",
				"nextWeek", "dddd [ti] LT",
				"lastDay", "[ikte ti] LT",
				"lastWeek", "[ovddit] dddd [ti] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s geažes",
				"past", "maŋit %s",
				"s", "moadde sekunddat",
				"ss", "%d sekunddat",
				"m", "okta minuhta",
				"mm", "%d minuhtat",
				"h", "okta diimmu",
				"hh", "%d diimmut",
				"d", "okta beaivi",
				"dd", "%d beaivvit",
				"M", "okta mánnu",
				"MM", "%d mánut",
				"y", "okta jahki",
				"yy", "%d jagit",
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
