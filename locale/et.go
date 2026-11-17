// moment.js locale configuration
// locale : Estonian [et]
// author : Henry Kehlmann : https://github.com/madhenry
// improvements : Illimar Tambek : https://github.com/ragulka

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("et", func() {
		moment.DefineLocale("et", js.Obj(
			"months", []string{"jaanuar", "veebruar", "märts", "aprill", "mai", "juuni", "juuli", "august", "september", "oktoober", "november", "detsember"},
			"monthsShort", []string{"jaan", "veebr", "märts", "apr", "mai", "juuni", "juuli", "aug", "sept", "okt", "nov", "dets"},
			"weekdays", []string{"pühapäev", "esmaspäev", "teisipäev", "kolmapäev", "neljapäev", "reede", "laupäev"},
			"weekdaysShort", []string{"P", "E", "T", "K", "N", "R", "L"},
			"weekdaysMin", []string{"P", "E", "T", "K", "N", "R", "L"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY H:mm",
				"LLLL", "dddd, D. MMMM YYYY H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Täna,] LT",
				"nextDay", "[Homme,] LT",
				"nextWeek", "[Järgmine] dddd LT",
				"lastDay", "[Eile,] LT",
				"lastWeek", "[Eelmine] dddd LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s pärast",
				"past", "%s tagasi",
				"s", etRelativeTimeSeconds,
				"ss", etRelativeTimeSeconds,
				"m", etRelativeTimeSeconds,
				"mm", etRelativeTimeSeconds,
				"h", etRelativeTimeSeconds,
				"hh", etRelativeTimeSeconds,
				"d", etRelativeTimeSeconds,
				"dd", "%d päeva",
				"M", etRelativeTimeSeconds,
				"MM", etRelativeTimeSeconds,
				"y", etRelativeTimeSeconds,
				"yy", etRelativeTimeSeconds,
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

// etRelativeTimeSeconds is processRelativeTime of et.
func etRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	n := num(number)
	format := map[string][]string{
		"s": {"mõne sekundi", "mõni sekund", "paar sekundit"}, "ss": {n + "sekundi", n + "sekundit"},
		"m": {"ühe minuti", "üks minut"}, "mm": {n + " minuti", n + " minutit"},
		"h": {"ühe tunni", "tund aega", "üks tund"}, "hh": {n + " tunni", n + " tundi"},
		"d": {"ühe päeva", "üks päev"}, "M": {"kuu aja", "kuu aega", "üks kuu"}, "MM": {n + " kuu", n + " kuud"},
		"y": {"ühe aasta", "aasta", "üks aasta"}, "yy": {n + " aasta", n + " aastat"},
	}
	forms := format[key]
	if withoutSuffix {
		if len(forms) > 2 {
			return forms[2]
		}
		return forms[1]
	}
	if isFuture {
		return forms[0]
	}
	return forms[1]
}
