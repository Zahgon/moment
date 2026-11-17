// moment.js locale configuration
// locale : Occitan, lengadocian dialecte [oc-lnc]
// author : Quentin PAGÈS : https://github.com/Quenty31

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("oc-lnc", func() {
		moment.DefineLocale("oc-lnc", js.Obj(
			"months", js.Obj(
				"standalone", []string{"genièr", "febrièr", "març", "abril", "mai", "junh", "julhet", "agost", "setembre", "octòbre", "novembre", "decembre"},
				"format", []string{"de genièr", "de febrièr", "de març", "d'abril", "de mai", "de junh", "de julhet", "d'agost", "de setembre", "d'octòbre", "de novembre", "de decembre"},
				"isFormat", jsre.MustCompile("D[oD]?(\\s)+MMMM", ""),
			),
			"monthsShort", []string{"gen.", "febr.", "març", "abr.", "mai", "junh", "julh.", "ago.", "set.", "oct.", "nov.", "dec."},
			"monthsParseExact", true,
			"weekdays", []string{"dimenge", "diluns", "dimars", "dimècres", "dijòus", "divendres", "dissabte"},
			"weekdaysShort", []string{"dg.", "dl.", "dm.", "dc.", "dj.", "dv.", "ds."},
			"weekdaysMin", []string{"dg", "dl", "dm", "dc", "dj", "dv", "ds"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM [de] YYYY",
				"ll", "D MMM YYYY",
				"LLL", "D MMMM [de] YYYY [a] H:mm",
				"lll", "D MMM YYYY, H:mm",
				"LLLL", "dddd D MMMM [de] YYYY [a] H:mm",
				"llll", "ddd D MMM YYYY, H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[uèi a] LT",
				"nextDay", "[deman a] LT",
				"nextWeek", "dddd [a] LT",
				"lastDay", "[ièr a] LT",
				"lastWeek", "dddd [passat a] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "d'aquí %s",
				"past", "fa %s",
				"s", "unas segondas",
				"ss", "%d segondas",
				"m", "una minuta",
				"mm", "%d minutas",
				"h", "una ora",
				"hh", "%d oras",
				"d", "un jorn",
				"dd", "%d jorns",
				"M", "un mes",
				"MM", "%d meses",
				"y", "un an",
				"yy", "%d ans",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(r|n|t|è|a)", ""),
			"ordinal", ocLncOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// ocLncOrdinal is the ordinal callback of oc-lnc (the rules of ca).
func ocLncOrdinal(l *moment.Locale, number moment.Value, period string) moment.Value {
	return caOrdinal(l, number, period)
}
