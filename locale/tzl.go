// moment.js locale configuration
// locale : Talossan [tzl]
// author : Robin van der Vliet : https://github.com/robin0van0der0v
// author : Iustì Canun

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("tzl", func() {
		moment.DefineLocale("tzl", js.Obj(
			"months", []string{"Januar", "Fevraglh", "Març", "Avrïu", "Mai", "Gün", "Julia", "Guscht", "Setemvar", "Listopäts", "Noemvar", "Zecemvar"},
			"monthsShort", []string{"Jan", "Fev", "Mar", "Avr", "Mai", "Gün", "Jul", "Gus", "Set", "Lis", "Noe", "Zec"},
			"weekdays", []string{"Súladi", "Lúneçi", "Maitzi", "Márcuri", "Xhúadi", "Viénerçi", "Sáturi"},
			"weekdaysShort", []string{"Súl", "Lún", "Mai", "Már", "Xhú", "Vié", "Sát"},
			"weekdaysMin", []string{"Sú", "Lú", "Ma", "Má", "Xh", "Vi", "Sá"},
			"longDateFormat", js.Obj(
				"LT", "HH.mm",
				"LTS", "HH.mm.ss",
				"L", "DD.MM.YYYY",
				"LL", "D. MMMM [dallas] YYYY",
				"LLL", "D. MMMM [dallas] YYYY HH.mm",
				"LLLL", "dddd, [li] D. MMMM [dallas] YYYY HH.mm",
			),
			"meridiemParse", jsre.MustCompile("d\\'o|d\\'a", "i"),
			"isPM", tzlIsPM,
			"meridiem", tzlMeridiem,
			"calendar", js.Obj(
				"sameDay", "[oxhi à] LT",
				"nextDay", "[demà à] LT",
				"nextWeek", "dddd [à] LT",
				"lastDay", "[ieiri à] LT",
				"lastWeek", "[sür el] dddd [lasteu à] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "osprei %s",
				"past", "ja%s",
				"s", tzlRelativeTimeSeconds,
				"ss", tzlRelativeTimeSeconds,
				"m", tzlRelativeTimeSeconds,
				"mm", tzlRelativeTimeSeconds,
				"h", tzlRelativeTimeSeconds,
				"hh", tzlRelativeTimeSeconds,
				"d", tzlRelativeTimeSeconds,
				"dd", tzlRelativeTimeSeconds,
				"M", tzlRelativeTimeSeconds,
				"MM", tzlRelativeTimeSeconds,
				"y", tzlRelativeTimeSeconds,
				"yy", tzlRelativeTimeSeconds,
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

// tzlIsPM is the isPM callback of tzl.
func tzlIsPM(_ *moment.Locale, input string) bool { return js.ToLower(input) == "d'o" }

// tzlMeridiem is the meridiem callback of tzl.
func tzlMeridiem(_ *moment.Locale, hours, _ float64, isLower bool) string {
	if hours > 11 {
		if isLower {
			return "d'o"
		}
		return "D'O"
	}
	if isLower {
		return "d'a"
	}
	return "D'A"
}

// tzlRelativeTimeSeconds is processRelativeTime of tzl.
func tzlRelativeTimeSeconds(number float64, withoutSuffix bool, key string, isFuture bool) string {
	n := num(number)
	format := map[string][2]string{
		"s": {"viensas secunds", "'iensas secunds"}, "ss": {n + " secunds", n + " secunds"},
		"m": {"'n míut", "'iens míut"}, "mm": {n + " míuts", n + " míuts"},
		"h": {"'n þora", "'iensa þora"}, "hh": {n + " þoras", n + " þoras"},
		"d": {"'n ziua", "'iensa ziua"}, "dd": {n + " ziuas", n + " ziuas"},
		"M": {"'n mes", "'iens mes"}, "MM": {n + " mesen", n + " mesen"},
		"y": {"'n ar", "'iens ar"}, "yy": {n + " ars", n + " ars"},
	}
	if isFuture || withoutSuffix {
		return format[key][0]
	}
	return format[key][1]
}
