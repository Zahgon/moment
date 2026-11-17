// moment.js locale configuration
// locale : Maori [mi]
// author : John Corrigan <robbiecloset@gmail.com> : https://github.com/johnideal

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("mi", func() {
		moment.DefineLocale("mi", js.Obj(
			"months", []string{"Kohi-tāte", "Hui-tanguru", "Poutū-te-rangi", "Paenga-whāwhā", "Haratua", "Pipiri", "Hōngoingoi", "Here-turi-kōkā", "Mahuru", "Whiringa-ā-nuku", "Whiringa-ā-rangi", "Hakihea"},
			"monthsShort", []string{"Kohi", "Hui", "Pou", "Pae", "Hara", "Pipi", "Hōngoi", "Here", "Mahu", "Whi-nu", "Whi-ra", "Haki"},
			"monthsRegex", jsre.MustCompile("(?:['a-z\\u0101\\u014D\\u016B]+\\-?){1,3}", "i"),
			"monthsStrictRegex", jsre.MustCompile("(?:['a-z\\u0101\\u014D\\u016B]+\\-?){1,3}", "i"),
			"monthsShortRegex", jsre.MustCompile("(?:['a-z\\u0101\\u014D\\u016B]+\\-?){1,3}", "i"),
			"monthsShortStrictRegex", jsre.MustCompile("(?:['a-z\\u0101\\u014D\\u016B]+\\-?){1,2}", "i"),
			"weekdays", []string{"Rātapu", "Mane", "Tūrei", "Wenerei", "Tāite", "Paraire", "Hātarei"},
			"weekdaysShort", []string{"Ta", "Ma", "Tū", "We", "Tāi", "Pa", "Hā"},
			"weekdaysMin", []string{"Ta", "Ma", "Tū", "We", "Tāi", "Pa", "Hā"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY [i] HH:mm",
				"LLLL", "dddd, D MMMM YYYY [i] HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[i teie mahana, i] LT",
				"nextDay", "[apopo i] LT",
				"nextWeek", "dddd [i] LT",
				"lastDay", "[inanahi i] LT",
				"lastWeek", "dddd [whakamutunga i] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "i roto i %s",
				"past", "%s i mua",
				"s", "te hēkona ruarua",
				"ss", "%d hēkona",
				"m", "he meneti",
				"mm", "%d meneti",
				"h", "te haora",
				"hh", "%d haora",
				"d", "he ra",
				"dd", "%d ra",
				"M", "he marama",
				"MM", "%d marama",
				"y", "he tau",
				"yy", "%d tau",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}º", ""),
			"ordinal", "%dº",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}
