// moment.js locale configuration
// locale : Basque [eu]
// author : Eneko Illarramendi : https://github.com/eillarra

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("eu", func() {
		moment.DefineLocale("eu", js.Obj(
			"months", []string{"urtarrila", "otsaila", "martxoa", "apirila", "maiatza", "ekaina", "uztaila", "abuztua", "iraila", "urria", "azaroa", "abendua"},
			"monthsShort", []string{"urt.", "ots.", "mar.", "api.", "mai.", "eka.", "uzt.", "abu.", "ira.", "urr.", "aza.", "abe."},
			"monthsParseExact", true,
			"weekdays", []string{"igandea", "astelehena", "asteartea", "asteazkena", "osteguna", "ostirala", "larunbata"},
			"weekdaysShort", []string{"ig.", "al.", "ar.", "az.", "og.", "ol.", "lr."},
			"weekdaysMin", []string{"ig", "al", "ar", "az", "og", "ol", "lr"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY-MM-DD",
				"LL", "YYYY[ko] MMMM[ren] D[a]",
				"LLL", "YYYY[ko] MMMM[ren] D[a] HH:mm",
				"LLLL", "dddd, YYYY[ko] MMMM[ren] D[a] HH:mm",
				"l", "YYYY-M-D",
				"ll", "YYYY[ko] MMM D[a]",
				"lll", "YYYY[ko] MMM D[a] HH:mm",
				"llll", "ddd, YYYY[ko] MMM D[a] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[gaur] LT[etan]",
				"nextDay", "[bihar] LT[etan]",
				"nextWeek", "dddd LT[etan]",
				"lastDay", "[atzo] LT[etan]",
				"lastWeek", "[aurreko] dddd LT[etan]",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s barru",
				"past", "duela %s",
				"s", "segundo batzuk",
				"ss", "%d segundo",
				"m", "minutu bat",
				"mm", "%d minutu",
				"h", "ordu bat",
				"hh", "%d ordu",
				"d", "egun bat",
				"dd", "%d egun",
				"M", "hilabete bat",
				"MM", "%d hilabete",
				"y", "urte bat",
				"yy", "%d urte",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}\\.", ""),
			"ordinal", "%d.",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}
