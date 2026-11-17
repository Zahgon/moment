// moment.js locale configuration
// locale : Albanian [sq]
// author : Flakërim Ismani : https://github.com/flakerimi
// author : Menelion Elensúle : https://github.com/Oire
// author : Oerd Cukalla : https://github.com/oerd

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("sq", func() {
		moment.DefineLocale("sq", js.Obj(
			"months", []string{"Janar", "Shkurt", "Mars", "Prill", "Maj", "Qershor", "Korrik", "Gusht", "Shtator", "Tetor", "Nëntor", "Dhjetor"},
			"monthsShort", []string{"Jan", "Shk", "Mar", "Pri", "Maj", "Qer", "Kor", "Gus", "Sht", "Tet", "Nën", "Dhj"},
			"weekdays", []string{"E Diel", "E Hënë", "E Martë", "E Mërkurë", "E Enjte", "E Premte", "E Shtunë"},
			"weekdaysShort", []string{"Die", "Hën", "Mar", "Mër", "Enj", "Pre", "Sht"},
			"weekdaysMin", []string{"D", "H", "Ma", "Më", "E", "P", "Sh"},
			"weekdaysParseExact", true,
			"meridiemParse", jsre.MustCompile("PD|MD", ""),
			"isPM", sqIsPM,
			"meridiem", sqMeridiem,
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Sot në] LT",
				"nextDay", "[Nesër në] LT",
				"nextWeek", "dddd [në] LT",
				"lastDay", "[Dje në] LT",
				"lastWeek", "dddd [e kaluar në] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "në %s",
				"past", "%s më parë",
				"s", "disa sekonda",
				"ss", "%d sekonda",
				"m", "një minutë",
				"mm", "%d minuta",
				"h", "një orë",
				"hh", "%d orë",
				"d", "një ditë",
				"dd", "%d ditë",
				"M", "një muaj",
				"MM", "%d muaj",
				"y", "një vit",
				"yy", "%d vite",
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

// sqIsPM is the isPM callback of sq.
func sqIsPM(_ *moment.Locale, input string) bool { return js.CharAt16(input, 0) == "M" }

// sqMeridiem is the meridiem callback of sq.
func sqMeridiem(_ *moment.Locale, hours, _ float64, _ bool) string {
	if hours < 12 {
		return "PD"
	}
	return "MD"
}
