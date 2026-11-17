// moment.js locale configuration
// locale : Welsh [cy]
// author : Robert Allen : https://github.com/robgallen
// author : https://github.com/ryangreaves

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("cy", func() {
		moment.DefineLocale("cy", js.Obj(
			"months", []string{"Ionawr", "Chwefror", "Mawrth", "Ebrill", "Mai", "Mehefin", "Gorffennaf", "Awst", "Medi", "Hydref", "Tachwedd", "Rhagfyr"},
			"monthsShort", []string{"Ion", "Chwe", "Maw", "Ebr", "Mai", "Meh", "Gor", "Aws", "Med", "Hyd", "Tach", "Rhag"},
			"weekdays", []string{"Dydd Sul", "Dydd Llun", "Dydd Mawrth", "Dydd Mercher", "Dydd Iau", "Dydd Gwener", "Dydd Sadwrn"},
			"weekdaysShort", []string{"Sul", "Llun", "Maw", "Mer", "Iau", "Gwe", "Sad"},
			"weekdaysMin", []string{"Su", "Ll", "Ma", "Me", "Ia", "Gw", "Sa"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Heddiw am] LT",
				"nextDay", "[Yfory am] LT",
				"nextWeek", "dddd [am] LT",
				"lastDay", "[Ddoe am] LT",
				"lastWeek", "dddd [diwethaf am] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "mewn %s",
				"past", "%s yn ôl",
				"s", "ychydig eiliadau",
				"ss", "%d eiliad",
				"m", "munud",
				"mm", "%d munud",
				"h", "awr",
				"hh", "%d awr",
				"d", "diwrnod",
				"dd", "%d diwrnod",
				"M", "mis",
				"MM", "%d mis",
				"y", "blwyddyn",
				"yy", "%d flynedd",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(fed|ain|af|il|ydd|ed|eg)", ""),
			"ordinal", cyOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// cyLookup is the suffix of 0 to 20, keyed like the indexes of the JavaScript array it replaces.
var cyLookup = map[string]string{"0": "", "1": "af", "2": "il", "3": "ydd", "4": "ydd", "5": "ed", "6": "ed",
	"7": "ed", "8": "fed", "9": "fed", "10": "fed", "11": "eg", "12": "fed", "13": "eg", "14": "eg", "15": "fed",
	"16": "eg", "17": "eg", "18": "fed", "19": "eg", "20": "fed"}

// cyOrdinal is the ordinal callback of cy.
func cyOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	n := js.ToNumber(number)
	var output moment.Value = ""
	switch {
	case n > 20:
		output = "ain"
		for _, k := range []float64{40, 50, 60, 80, 100} {
			if sameNumber(number, k) {
				output = "fed"
			}
		}
	case n > 0:
		output = js.Undef
		if s, ok := cyLookup[str(number)]; ok {
			output = s
		}
	}
	return js.Add(number, output)
}
