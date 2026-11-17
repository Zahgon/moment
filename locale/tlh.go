// moment.js locale configuration
// locale : Klingon [tlh]
// author : Dominika Kruk : https://github.com/amaranthrose

package locale

import (
	"math"

	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("tlh", func() {
		moment.DefineLocale("tlh", js.Obj(
			"months", []string{"tera’ jar wa’", "tera’ jar cha’", "tera’ jar wej", "tera’ jar loS", "tera’ jar vagh", "tera’ jar jav", "tera’ jar Soch", "tera’ jar chorgh", "tera’ jar Hut", "tera’ jar wa’maH", "tera’ jar wa’maH wa’", "tera’ jar wa’maH cha’"},
			"monthsShort", []string{"jar wa’", "jar cha’", "jar wej", "jar loS", "jar vagh", "jar jav", "jar Soch", "jar chorgh", "jar Hut", "jar wa’maH", "jar wa’maH wa’", "jar wa’maH cha’"},
			"monthsParseExact", true,
			"weekdays", []string{"lojmItjaj", "DaSjaj", "povjaj", "ghItlhjaj", "loghjaj", "buqjaj", "ghInjaj"},
			"weekdaysShort", []string{"lojmItjaj", "DaSjaj", "povjaj", "ghItlhjaj", "loghjaj", "buqjaj", "ghInjaj"},
			"weekdaysMin", []string{"lojmItjaj", "DaSjaj", "povjaj", "ghItlhjaj", "loghjaj", "buqjaj", "ghInjaj"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[DaHjaj] LT",
				"nextDay", "[wa’leS] LT",
				"nextWeek", "LLL",
				"lastDay", "[wa’Hu’] LT",
				"lastWeek", "LLL",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", tlhRelativeTimeFuture,
				"past", tlhRelativeTimePast,
				"s", "puS lup",
				"ss", tlhRelativeTimeSecondsN,
				"m", "wa’ tup",
				"mm", tlhRelativeTimeSecondsN,
				"h", "wa’ rep",
				"hh", tlhRelativeTimeSecondsN,
				"d", "wa’ jaj",
				"dd", tlhRelativeTimeSecondsN,
				"M", "wa’ jar",
				"MM", tlhRelativeTimeSecondsN,
				"y", "wa’ DIS",
				"yy", tlhRelativeTimeSecondsN,
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

var tlhNumberNouns = []string{"pagh", "wa’", "cha’", "wej", "loS", "vagh", "jav", "Soch", "chorgh", "Hut"}

// tlhNoun returns the Klingon word for digit d ("undefined" for a non-digit).
func tlhNoun(d float64) string {
	if d >= 0 && d < 10 && d == math.Trunc(d) {
		return tlhNumberNouns[int(d)]
	}
	return "undefined"
}

// tlhNumberAsNoun spells a number below 1000 with hundreds, tens and ones.
func tlhNumberAsNoun(number float64) string {
	hundred := math.Floor(js.Mod(number, 1000) / 100)
	ten := math.Floor(js.Mod(number, 100) / 10)
	one := js.Mod(number, 10)
	word := ""
	if hundred > 0 {
		word += tlhNoun(hundred) + "vatlh"
	}
	if ten > 0 {
		if word != "" {
			word += " "
		}
		word += tlhNoun(ten) + "maH"
	}
	if one > 0 {
		if word != "" {
			word += " "
		}
		word += tlhNoun(one)
	}
	if word == "" {
		return "pagh"
	}
	return word
}

// tlhTranslateEnding replaces the last three letters of a day, month or year phrase, or appends the suffix.
func tlhTranslateEnding(output, day, month, year, other string) string {
	cut := js.Slice16(output, 0, -3)
	switch {
	case js.IndexOf16(output, "jaj") != -1:
		return cut + day
	case js.IndexOf16(output, "jar") != -1:
		return cut + month
	case js.IndexOf16(output, "DIS") != -1:
		return cut + year
	}
	return output + other
}

// tlhRelativeTimeFuture is translateFuture of tlh.
func tlhRelativeTimeFuture(output string) string {
	return tlhTranslateEnding(output, "leS", "waQ", "nem", " pIq")
}

// tlhRelativeTimePast is translatePast of tlh.
func tlhRelativeTimePast(output string) string {
	return tlhTranslateEnding(output, "Hu’", "wen", "ben", " ret")
}

// tlhRelativeTimeSecondsN is translate of tlh.
func tlhRelativeTimeSecondsN(number float64, _ bool, key string, _ bool) string {
	units := map[string]string{"ss": " lup", "mm": " tup", "hh": " rep", "dd": " jaj", "MM": " jar", "yy": " DIS"}
	unit, ok := units[key]
	if !ok {
		return ""
	}
	return tlhNumberAsNoun(number) + unit
}
