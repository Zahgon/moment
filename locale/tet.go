// moment.js locale configuration
// locale : Tetun Dili (East Timor) [tet]
// author : Joshua Brooks : https://github.com/joshbrooks
// author : Onorio De J. Afonso : https://github.com/marobo
// author : Sonia Simoes : https://github.com/soniasimoes

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("tet", func() {
		moment.DefineLocale("tet", js.Obj(
			"months", []string{"Janeiru", "Fevereiru", "Marsu", "Abril", "Maiu", "Juñu", "Jullu", "Agustu", "Setembru", "Outubru", "Novembru", "Dezembru"},
			"monthsShort", []string{"Jan", "Fev", "Mar", "Abr", "Mai", "Jun", "Jul", "Ago", "Set", "Out", "Nov", "Dez"},
			"weekdays", []string{"Domingu", "Segunda", "Tersa", "Kuarta", "Kinta", "Sesta", "Sabadu"},
			"weekdaysShort", []string{"Dom", "Seg", "Ters", "Kua", "Kint", "Sest", "Sab"},
			"weekdaysMin", []string{"Do", "Seg", "Te", "Ku", "Ki", "Ses", "Sa"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[Ohin iha] LT",
				"nextDay", "[Aban iha] LT",
				"nextWeek", "dddd [iha] LT",
				"lastDay", "[Horiseik iha] LT",
				"lastWeek", "dddd [semana kotuk] [iha] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "iha %s",
				"past", "%s liuba",
				"s", "segundu balun",
				"ss", "segundu %d",
				"m", "minutu ida",
				"mm", "minutu %d",
				"h", "oras ida",
				"hh", "oras %d",
				"d", "loron ida",
				"dd", "loron %d",
				"M", "fulan ida",
				"MM", "fulan %d",
				"y", "tinan ida",
				"yy", "tinan %d",
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(st|nd|rd|th)", ""),
			"ordinal", tetOrdinal,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

var tetOrdinal = englishOrdinal
