// moment.js locale configuration
// locale : Swedish [sv]
// author : Jens Alm : https://github.com/ulmus

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("sv", func() {
		moment.DefineLocale("sv", js.Obj(
			"months", []string{"januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"},
			"monthsShort", []string{"jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
			"weekdays", []string{"söndag", "måndag", "tisdag", "onsdag", "torsdag", "fredag", "lördag"},
			"weekdaysShort", []string{"sön", "mån", "tis", "ons", "tors", "fre", "lör"},
			"weekdaysMin", []string{"sö", "må", "ti", "on", "to", "fr", "lö"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "YYYY-MM-DD",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY [kl.] HH:mm",
				"LLLL", "dddd D MMMM YYYY [kl.] HH:mm",
				"lll", "D MMM YYYY HH:mm",
				"llll", "ddd D MMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Idag] LT",
				"nextDay", "[Imorgon] LT",
				"lastDay", "[Igår] LT",
				"nextWeek", "[På] dddd LT",
				"lastWeek", "[I] dddd[s] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "om %s",
				"past", "för %s sedan",
				"s", "några sekunder",
				"ss", "%d sekunder",
				"m", "en minut",
				"mm", "%d minuter",
				"h", "en timme",
				"hh", "%d timmar",
				"d", "en dag",
				"dd", "%d dagar",
				"M", "en månad",
				"MM", "%d månader",
				"y", "ett år",
				"yy", "%d år",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(\\:e|\\:a)", ""),
			"ordinal", svOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// svOrdinal is the ordinal callback of sv.
func svOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	n := js.ToNumber(number)
	b := js.Mod(n, 10)
	output := ":e"
	if js.ToInt32(js.Mod(n, 100)/10) != 1 && (b == 1 || b == 2) {
		output = ":a"
	}
	return str(number) + output
}
