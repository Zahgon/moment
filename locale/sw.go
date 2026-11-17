// moment.js locale configuration
// locale : Swahili [sw]
// author : Fahad Kassim : https://github.com/fadsel

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
)

func init() {
	moment.RegisterLocale("sw", func() {
		moment.DefineLocale("sw", js.Obj(
			"months", []string{"Januari", "Februari", "Machi", "Aprili", "Mei", "Juni", "Julai", "Agosti", "Septemba", "Oktoba", "Novemba", "Desemba"},
			"monthsShort", []string{"Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ago", "Sep", "Okt", "Nov", "Des"},
			"weekdays", []string{"Jumapili", "Jumatatu", "Jumanne", "Jumatano", "Alhamisi", "Ijumaa", "Jumamosi"},
			"weekdaysShort", []string{"Jpl", "Jtat", "Jnne", "Jtan", "Alh", "Ijm", "Jmos"},
			"weekdaysMin", []string{"J2", "J3", "J4", "J5", "Al", "Ij", "J1"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "hh:mm A",
				"LTS", "HH:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd, D MMMM YYYY HH:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[leo saa] LT",
				"nextDay", "[kesho saa] LT",
				"nextWeek", "dddd [ijayo saa] LT",
				"lastDay", "[jana] LT",
				"lastWeek", "dddd [iliyopita saa] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s baadaye",
				"past", swRelativeTimePast,
				"s", "hivi punde",
				"ss", swRelativeTimeSecondsN,
				"m", swRelativeTimeSecondsN,
				"mm", swRelativeTimeSecondsN,
				"h", swRelativeTimeSecondsN,
				"hh", swRelativeTimeSecondsN,
				"d", swRelativeTimeSecondsN,
				"dd", swRelativeTimeSecondsN,
				"M", swRelativeTimeSecondsN,
				"MM", swRelativeTimeSecondsN,
				"y", swRelativeTimeSecondsN,
				"yy", swRelativeTimeSecondsN,
			),
			"week", js.Obj(
				"dow", 1.0,
				"doy", 7.0,
			),
		))
	})
}

var swForms = map[string]string{"ss": "sekunde %d", "m": "dakika moja", "mm": "dakika %d", "h": "saa moja",
	"hh": "saa %d", "d": "siku moja", "dd": "siku %d", "M": "mwezi mmoja", "MM": "miezi %d", "y": "mwaka mmoja",
	"yy": "miaka %d"}

// swRelativeTimePast is relativeTime.past of sw.
func swRelativeTimePast(output string) string {
	if output == "hivi punde" {
		return "tokea " + output
	}
	return output
}

// swRelativeTimeSecondsN is relativeTimeWithSuffix of sw.
func swRelativeTimeSecondsN(number float64, withoutSuffix bool, key string, isFuture bool) string {
	output := replaceCount(swForms[key], number)
	if withoutSuffix || isFuture {
		return output
	}
	switch key {
	case "ss", "mm", "hh", "dd":
		if number == 1 {
			return output + " iliyopita"
		}
		return output + " zilizopita"
	case "m", "h", "d", "MM", "yy":
		return output + " iliyopita"
	case "M", "y":
		return output + " uliopita"
	}
	return ""
}
