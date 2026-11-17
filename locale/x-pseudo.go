// moment.js locale configuration
// locale : Pseudo [x-pseudo]
// author : Andrew Hood : https://github.com/andrewhood125

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("x-pseudo", func() {
		moment.DefineLocale("x-pseudo", js.Obj(
			"months", []string{"J~áñúá~rý", "F~ébrú~árý", "~Márc~h", "Áp~ríl", "~Máý", "~Júñé~", "Júl~ý", "Áú~gúst~", "Sép~témb~ér", "Ó~ctób~ér", "Ñ~óvém~bér", "~Décé~mbér"},
			"monthsShort", []string{"J~áñ", "~Féb", "~Már", "~Ápr", "~Máý", "~Júñ", "~Júl", "~Áúg", "~Sép", "~Óct", "~Ñóv", "~Déc"},
			"monthsParseExact", true,
			"weekdays", []string{"S~úñdá~ý", "Mó~ñdáý~", "Túé~sdáý~", "Wéd~ñésd~áý", "T~húrs~dáý", "~Fríd~áý", "S~átúr~dáý"},
			"weekdaysShort", []string{"S~úñ", "~Móñ", "~Túé", "~Wéd", "~Thú", "~Frí", "~Sát"},
			"weekdaysMin", []string{"S~ú", "Mó~", "Tú", "~Wé", "T~h", "Fr~", "Sá"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[T~ódá~ý át] LT",
				"nextDay", "[T~ómó~rró~w át] LT",
				"nextWeek", "dddd [át] LT",
				"lastDay", "[Ý~ést~érdá~ý át] LT",
				"lastWeek", "[L~ást] dddd [át] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "í~ñ %s",
				"past", "%s á~gó",
				"s", "á ~féw ~sécó~ñds",
				"ss", "%d s~écóñ~ds",
				"m", "á ~míñ~úté",
				"mm", "%d m~íñú~tés",
				"h", "á~ñ hó~úr",
				"hh", "%d h~óúrs",
				"d", "á ~dáý",
				"dd", "%d d~áýs",
				"M", "á ~móñ~th",
				"MM", "%d m~óñt~hs",
				"y", "á ~ýéár",
				"yy", "%d ý~éárs",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(th|st|nd|rd)", ""),
			"ordinal", xPseudoOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var xPseudoOrdinal = englishOrdinal
