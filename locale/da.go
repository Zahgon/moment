// moment.js locale configuration
// locale : Danish [da]
// author : Ulrik Nielsen : https://github.com/mrbase

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("da", func() {
		moment.DefineLocale("da", js.Obj(
			"months", []string{"januar", "februar", "marts", "april", "maj", "juni", "juli", "august", "september", "oktober", "november", "december"},
			"monthsShort", []string{"jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
			"weekdays", []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"},
			"weekdaysShort", []string{"søn", "man", "tir", "ons", "tor", "fre", "lør"},
			"weekdaysMin", []string{"sø", "ma", "ti", "on", "to", "fr", "lø"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY HH:mm",
				"LLLL", "dddd [d.] D. MMMM YYYY [kl.] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[i dag kl.] LT",
				"nextDay", "[i morgen kl.] LT",
				"nextWeek", "på dddd [kl.] LT",
				"lastDay", "[i går kl.] LT",
				"lastWeek", "[i] dddd[s kl.] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "om %s",
				"past", "%s siden",
				"s", "få sekunder",
				"ss", "%d sekunder",
				"m", "et minut",
				"mm", "%d minutter",
				"h", "en time",
				"hh", "%d timer",
				"d", "en dag",
				"dd", "%d dage",
				"M", "en måned",
				"MM", "%d måneder",
				"y", "et år",
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
