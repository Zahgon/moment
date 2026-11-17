// moment.js locale configuration
// locale : siSwati [ss]
// author : Nicolai Davies<mail@nicolai.io> : https://github.com/nicolaidavies

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ss", func() {
		moment.DefineLocale("ss", js.Obj(
			"months", []string{"Bhimbidvwane", "Indlovana", "Indlov'lenkhulu", "Mabasa", "Inkhwekhweti", "Inhlaba", "Kholwane", "Ingci", "Inyoni", "Imphala", "Lweti", "Ingongoni"},
			"monthsShort", []string{"Bhi", "Ina", "Inu", "Mab", "Ink", "Inh", "Kho", "Igc", "Iny", "Imp", "Lwe", "Igo"},
			"weekdays", []string{"Lisontfo", "Umsombuluko", "Lesibili", "Lesitsatfu", "Lesine", "Lesihlanu", "Umgcibelo"},
			"weekdaysShort", []string{"Lis", "Umb", "Lsb", "Les", "Lsi", "Lsh", "Umg"},
			"weekdaysMin", []string{"Li", "Us", "Lb", "Lt", "Ls", "Lh", "Ug"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "h:mm A",
				"LTS", "h:mm:ss A",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY h:mm A",
				"LLLL", "dddd, D MMMM YYYY h:mm A",
			),
			"calendar", js.Obj(
				"sameDay", "[Namuhla nga] LT",
				"nextDay", "[Kusasa nga] LT",
				"nextWeek", "dddd [nga] LT",
				"lastDay", "[Itolo nga] LT",
				"lastWeek", "dddd [leliphelile] [nga] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "nga %s",
				"past", "wenteka nga %s",
				"s", "emizuzwana lomcane",
				"ss", "%d mzuzwana",
				"m", "umzuzu",
				"mm", "%d emizuzu",
				"h", "lihora",
				"hh", "%d emahora",
				"d", "lilanga",
				"dd", "%d emalanga",
				"M", "inyanga",
				"MM", "%d tinyanga",
				"y", "umnyaka",
				"yy", "%d iminyaka",
			),
			"meridiemParse", jsre.MustCompile("ekuseni|emini|entsambama|ebusuku", ""),
			"meridiem", ssMeridiem,
			"meridiemHour", ssMeridiemHour,
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}", ""),
			"ordinal", "%d",
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// ssMeridiem is the meridiem callback of ss.
func ssMeridiem(_ *moment.Locale, hours, _ float64, _ bool) string {
	switch {
	case hours < 11:
		return "ekuseni"
	case hours < 15:
		return "emini"
	case hours < 19:
		return "entsambama"
	}
	return "ebusuku"
}

// ssMeridiemHour is the meridiemHour callback of ss.
func ssMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	switch meridiem {
	case "ekuseni":
		return hour
	case "emini":
		if hourNumber(hour) >= 11 {
			return hour
		}
		return plus12(hour)
	case "entsambama", "ebusuku":
		if js.StrictEquals(hour, 0.0) {
			return 0.0
		}
		return plus12(hour)
	}
	return js.Undef
}
