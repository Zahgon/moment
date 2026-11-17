// moment.js locale configuration
// locale : Irish or Irish Gaelic [ga]
// author : André Silva : https://github.com/askpt

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ga", func() {
		moment.DefineLocale("ga", js.Obj(
			"months", []string{"Eanáir", "Feabhra", "Márta", "Aibreán", "Bealtaine", "Meitheamh", "Iúil", "Lúnasa", "Meán Fómhair", "Deireadh Fómhair", "Samhain", "Nollaig"},
			"monthsShort", []string{"Ean", "Feabh", "Márt", "Aib", "Beal", "Meith", "Iúil", "Lún", "M.F.", "D.F.", "Samh", "Noll"},
			"monthsParseExact", true,
			"weekdays", []string{"Dé Domhnaigh", "Dé Luain", "Dé Máirt", "Dé Céadaoin", "Déardaoin", "Dé hAoine", "Dé Sathairn"},
			"weekdaysShort", []string{"Domh", "Luan", "Máirt", "Céad", "Déar", "Aoine", "Sath"},
			"weekdaysMin", []string{"Do", "Lu", "Má", "Cé", "Dé", "A", "Sa"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Inniu ag] LT",
				"nextDay", "[Amárach ag] LT",
				"nextWeek", "dddd [ag] LT",
				"lastDay", "[Inné ag] LT",
				"lastWeek", "dddd [seo caite] [ag] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "i %s",
				"past", "%s ó shin",
				"s", "cúpla soicind",
				"ss", "%d soicind",
				"m", "nóiméad",
				"mm", "%d nóiméad",
				"h", "uair an chloig",
				"hh", "%d uair an chloig",
				"d", "lá",
				"dd", "%d lá",
				"M", "mí",
				"MM", "%d míonna",
				"y", "bliain",
				"yy", "%d bliain",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(d|na|mh)", ""),
			"ordinal", gaOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// gaOrdinal is the ordinal callback of ga (and gd).
func gaOrdinal(_ *moment.Locale, number moment.Value, _ string) moment.Value {
	switch {
	case sameNumber(number, 1):
		return str(number) + "d"
	case js.Mod(js.ToNumber(number), 10) == 2:
		return str(number) + "na"
	}
	return str(number) + "mh"
}
