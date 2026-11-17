// moment.js locale configuration
// locale : Konkani Latin script [gom-latn]
// author : The Discoverer : https://github.com/WikiDiscoverer

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("gom-latn", func() {
		moment.DefineLocale("gom-latn", js.Obj(
			"months", js.Obj(
				"standalone", []string{"Janer", "Febrer", "Mars", "Abril", "Mai", "Jun", "Julai", "Agost", "Setembr", "Otubr", "Novembr", "Dezembr"},
				"format", []string{"Janerachea", "Febrerachea", "Marsachea", "Abrilachea", "Maiachea", "Junachea", "Julaiachea", "Agostachea", "Setembrachea", "Otubrachea", "Novembrachea", "Dezembrachea"},
				"isFormat", jsre.MustCompile("MMMM(\\s)+D[oD]?", ""),
			),
			"monthsShort", []string{"Jan.", "Feb.", "Mars", "Abr.", "Mai", "Jun", "Jul.", "Ago.", "Set.", "Otu.", "Nov.", "Dez."},
			"monthsParseExact", true,
			"weekdays", []string{"Aitar", "Somar", "Mongllar", "Budhvar", "Birestar", "Sukrar", "Son'var"},
			"weekdaysShort", []string{"Ait.", "Som.", "Mon.", "Bud.", "Bre.", "Suk.", "Son."},
			"weekdaysMin", []string{"Ai", "Sm", "Mo", "Bu", "Br", "Su", "Sn"},
			"weekdaysParseExact", true,
			"longDateFormat", js.Obj(
				"LT", "A h:mm [vazta]",
				"LTS", "A h:mm:ss [vazta]",
				"L", "DD-MM-YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY A h:mm [vazta]",
				"LLLL", "dddd, MMMM Do, YYYY, A h:mm [vazta]",
				"llll", "ddd, D MMM YYYY, A h:mm [vazta]",
			),
			"calendar", js.Obj(
				"sameDay", "[Aiz] LT",
				"nextDay", "[Faleam] LT",
				"nextWeek", "[Fuddlo] dddd[,] LT",
				"lastDay", "[Kal] LT",
				"lastWeek", "[Fattlo] dddd[,] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s",
				"past", "%s adim",
				"s", gomLatnRelativeTimeSeconds,
				"ss", gomLatnRelativeTimeSeconds,
				"m", gomLatnRelativeTimeSeconds,
				"mm", gomLatnRelativeTimeSeconds,
				"h", gomLatnRelativeTimeSeconds,
				"hh", gomLatnRelativeTimeSeconds,
				"d", gomLatnRelativeTimeSeconds,
				"dd", gomLatnRelativeTimeSeconds,
				"M", gomLatnRelativeTimeSeconds,
				"MM", gomLatnRelativeTimeSeconds,
				"y", gomLatnRelativeTimeSeconds,
				"yy", gomLatnRelativeTimeSeconds,
			),
			"dayOfMonthOrdinalParse", jsre.MustCompile("\\d{1,2}(er)", ""),
			"ordinal", gomLatnOrdinal,
			"week", js.Obj(
				"dow", 0.0,
				"doy", 3.0,
			),
			"meridiemParse", jsre.MustCompile("rati|sokallim|donparam|sanje", ""),
			"meridiemHour", gomLatnMeridiemHour,
			"meridiem", gomLatnMeridiem,
		))
	})
}

// gomLatnRelativeTimeSeconds is processRelativeTime of gom-latn.
func gomLatnRelativeTimeSeconds(number float64, _ bool, key string, isFuture bool) string {
	n := num(number)
	format := map[string][2]string{
		"s": {"thoddea sekondamni", "thodde sekond"}, "ss": {n + " sekondamni", n + " sekond"},
		"m": {"eka mintan", "ek minut"}, "mm": {n + " mintamni", n + " mintam"},
		"h": {"eka voran", "ek vor"}, "hh": {n + " voramni", n + " voram"},
		"d": {"eka disan", "ek dis"}, "dd": {n + " disamni", n + " dis"},
		"M": {"eka mhoinean", "ek mhoino"}, "MM": {n + " mhoineamni", n + " mhoine"},
		"y": {"eka vorsan", "ek voros"}, "yy": {n + " vorsamni", n + " vorsam"},
	}
	if isFuture {
		return format[key][0]
	}
	return format[key][1]
}

// gomLatnOrdinal is the ordinal callback of gom-latn.
func gomLatnOrdinal(_ *moment.Locale, number moment.Value, period string) moment.Value {
	if period == "D" {
		return str(number) + "er"
	}
	return number
}

// gomLatnMeridiemHour is the meridiemHour callback of gom-latn.
func gomLatnMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	return konkaniMeridiemHour(hour, meridiem, "rati", "sokallim", "donparam", "sanje")
}

// gomLatnMeridiem is the meridiem callback of gom-latn.
func gomLatnMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "rati"
	case hour < 12:
		return "sokallim"
	case hour < 16:
		return "donparam"
	case hour < 20:
		return "sanje"
	}
	return "rati"
}
