// moment.js locale configuration
// locale : Icelandic [is]
// author : Hinrik Örn Sigurðsson : https://github.com/hinrik

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("is", func() {
		moment.DefineLocale("is", js.Obj(
			"months", []string{"janúar", "febrúar", "mars", "apríl", "maí", "júní", "júlí", "ágúst", "september", "október", "nóvember", "desember"},
			"monthsShort", []string{"jan", "feb", "mar", "apr", "maí", "jún", "júl", "ágú", "sep", "okt", "nóv", "des"},
			"weekdays", []string{"sunnudagur", "mánudagur", "þriðjudagur", "miðvikudagur", "fimmtudagur", "föstudagur", "laugardagur"},
			"weekdaysShort", []string{"sun", "mán", "þri", "mið", "fim", "fös", "lau"},
			"weekdaysMin", []string{"Su", "Má", "Þr", "Mi", "Fi", "Fö", "La"},
			"longDateFormat", js.Obj(
				"LT", "H:mm",
				"LTS", "H:mm:ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM YYYY",
				"LLL", "D. MMMM YYYY [kl.] H:mm",
				"LLLL", "dddd, D. MMMM YYYY [kl.] H:mm",
			),
			"calendar", js.Obj(
				"sameDay", "[í dag kl.] LT",
				"nextDay", "[á morgun kl.] LT",
				"nextWeek", "dddd [kl.] LT",
				"lastDay", "[í gær kl.] LT",
				"lastWeek", "[síðasta] dddd [kl.] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "eftir %s",
				"past", "fyrir %s síðan",
				"s", isRelativeTimeSeconds,
				"ss", isRelativeTimeSeconds,
				"m", isRelativeTimeSeconds,
				"mm", isRelativeTimeSeconds,
				"h", "klukkustund",
				"hh", isRelativeTimeSeconds,
				"d", isRelativeTimeSeconds,
				"dd", isRelativeTimeSeconds,
				"M", isRelativeTimeSeconds,
				"MM", isRelativeTimeSeconds,
				"y", isRelativeTimeSeconds,
				"yy", isRelativeTimeSeconds,
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

// isPlural is false for numbers ending in 1 (except 11).
func isPlural(n float64) bool {
	if js.Mod(n, 100) == 11 {
		return true
	}
	return js.Mod(n, 10) != 1
}

// isRelativeTimeSeconds is translate of is.
func isRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	result := num(number) + " "
	either := func(a, b string) string {
		if withoutSuffix || isFuture {
			return a
		}
		return b
	}
	future := func(a, b string) string {
		if isFuture {
			return a
		}
		return b
	}
	plural := isPlural(number)
	switch key {
	case "s":
		return either("nokkrar sekúndur", "nokkrum sekúndum")
	case "ss":
		if plural {
			return result + either("sekúndur", "sekúndum")
		}
		return result + "sekúnda"
	case "m":
		if withoutSuffix {
			return "mínúta"
		}
		return "mínútu"
	case "mm":
		switch {
		case plural:
			return result + either("mínútur", "mínútum")
		case withoutSuffix:
			return result + "mínúta"
		}
		return result + "mínútu"
	case "hh":
		if plural {
			return result + either("klukkustundir", "klukkustundum")
		}
		return result + "klukkustund"
	case "d":
		if withoutSuffix {
			return "dagur"
		}
		return future("dag", "degi")
	case "dd":
		switch {
		case plural && withoutSuffix:
			return result + "dagar"
		case plural:
			return result + future("daga", "dögum")
		case withoutSuffix:
			return result + "dagur"
		}
		return result + future("dag", "degi")
	case "M":
		if withoutSuffix {
			return "mánuður"
		}
		return future("mánuð", "mánuði")
	case "MM":
		switch {
		case plural && withoutSuffix:
			return result + "mánuðir"
		case plural:
			return result + future("mánuði", "mánuðum")
		case withoutSuffix:
			return result + "mánuður"
		}
		return result + future("mánuð", "mánuði")
	case "y":
		return either("ár", "ári")
	case "yy":
		if plural {
			return result + either("ár", "árum")
		}
		return result + either("ár", "ári")
	}
	return ""
}
