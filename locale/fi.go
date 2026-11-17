// moment.js locale configuration
// locale : Finnish [fi]
// author : Tarmo Aidantausta : https://github.com/bleadof

package locale

import (
	"math"

	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("fi", func() {
		moment.DefineLocale("fi", js.Obj(
			"months", []string{"tammikuu", "helmikuu", "maaliskuu", "huhtikuu", "toukokuu", "kesäkuu", "heinäkuu", "elokuu", "syyskuu", "lokakuu", "marraskuu", "joulukuu"},
			"monthsShort", []string{"tammi", "helmi", "maalis", "huhti", "touko", "kesä", "heinä", "elo", "syys", "loka", "marras", "joulu"},
			"weekdays", []string{"sunnuntai", "maanantai", "tiistai", "keskiviikko", "torstai", "perjantai", "lauantai"},
			"weekdaysShort", []string{"su", "ma", "ti", "ke", "to", "pe", "la"},
			"weekdaysMin", []string{"su", "ma", "ti", "ke", "to", "pe", "la"},
			"longDateFormat", js.Obj(
				"LT", "HH.mm",
				"LTS", "HH.mm.ss",
				"L", "DD.MM.YYYY",
				"LL", "Do MMMM[ta] YYYY",
				"LLL", "Do MMMM[ta] YYYY, [klo] HH.mm",
				"LLLL", "dddd, Do MMMM[ta] YYYY, [klo] HH.mm",
				"l", "D.M.YYYY",
				"ll", "Do MMM YYYY",
				"lll", "Do MMM YYYY, [klo] HH.mm",
				"llll", "ddd, Do MMM YYYY, [klo] HH.mm",
			),
			"calendar", js.Obj(
				"sameDay", "[tänään] [klo] LT",
				"nextDay", "[huomenna] [klo] LT",
				"nextWeek", "dddd [klo] LT",
				"lastDay", "[eilen] [klo] LT",
				"lastWeek", "[viime] dddd[na] [klo] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s päästä",
				"past", "%s sitten",
				"s", fiRelativeTimeSeconds,
				"ss", fiRelativeTimeSeconds,
				"m", fiRelativeTimeSeconds,
				"mm", fiRelativeTimeSeconds,
				"h", fiRelativeTimeSeconds,
				"hh", fiRelativeTimeSeconds,
				"d", fiRelativeTimeSeconds,
				"dd", fiRelativeTimeSeconds,
				"M", fiRelativeTimeSeconds,
				"MM", fiRelativeTimeSeconds,
				"y", fiRelativeTimeSeconds,
				"yy", fiRelativeTimeSeconds,
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

var (
	fiNumbersPast   = []string{"nolla", "yksi", "kaksi", "kolme", "neljä", "viisi", "kuusi", "seitsemän", "kahdeksan", "yhdeksän"}
	fiNumbersFuture = []string{"nolla", "yhden", "kahden", "kolmen", "neljän", "viiden", "kuuden", "seitsemän", "kahdeksan", "yhdeksän"}
)

// fiVerbalNumber spells numbers below ten.
func fiVerbalNumber(number float64, isFuture bool) string {
	if number >= 10 || number != number {
		return num(number)
	}
	if number < 0 || number != math.Trunc(number) {
		return "undefined"
	}
	if isFuture {
		return fiNumbersFuture[int(number)]
	}
	return fiNumbersPast[int(number)]
}

// fiRelativeTimeSeconds is translate of fi.
func fiRelativeTimeSeconds(number float64, _ bool, key string, isFuture bool) string {
	pick := func(future, past string) string {
		if isFuture {
			return future
		}
		return past
	}
	result := ""
	switch key {
	case "s":
		return pick("muutaman sekunnin", "muutama sekunti")
	case "ss":
		result = pick("sekunnin", "sekuntia")
	case "m":
		return pick("minuutin", "minuutti")
	case "mm":
		result = pick("minuutin", "minuuttia")
	case "h":
		return pick("tunnin", "tunti")
	case "hh":
		result = pick("tunnin", "tuntia")
	case "d":
		return pick("päivän", "päivä")
	case "dd":
		result = pick("päivän", "päivää")
	case "M":
		return pick("kuukauden", "kuukausi")
	case "MM":
		result = pick("kuukauden", "kuukautta")
	case "y":
		return pick("vuoden", "vuosi")
	case "yy":
		result = pick("vuoden", "vuotta")
	}
	return fiVerbalNumber(number, isFuture) + " " + result
}
