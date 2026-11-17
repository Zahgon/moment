// moment.js locale configuration
// locale : Malayalam [ml]
// author : Floyd Pink : https://github.com/floydpink

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("ml", func() {
		moment.DefineLocale("ml", js.Obj(
			"months", []string{"ജനുവരി", "ഫെബ്രുവരി", "മാർച്ച്", "ഏപ്രിൽ", "മേയ്", "ജൂൺ", "ജൂലൈ", "ഓഗസ്റ്റ്", "സെപ്റ്റംബർ", "ഒക്ടോബർ", "നവംബർ", "ഡിസംബർ"},
			"monthsShort", []string{"ജനു.", "ഫെബ്രു.", "മാർ.", "ഏപ്രി.", "മേയ്", "ജൂൺ", "ജൂലൈ.", "ഓഗ.", "സെപ്റ്റ.", "ഒക്ടോ.", "നവം.", "ഡിസം."},
			"monthsParseExact", true,
			"weekdays", []string{"ഞായറാഴ്ച", "തിങ്കളാഴ്ച", "ചൊവ്വാഴ്ച", "ബുധനാഴ്ച", "വ്യാഴാഴ്ച", "വെള്ളിയാഴ്ച", "ശനിയാഴ്ച"},
			"weekdaysShort", []string{"ഞായർ", "തിങ്കൾ", "ചൊവ്വ", "ബുധൻ", "വ്യാഴം", "വെള്ളി", "ശനി"},
			"weekdaysMin", []string{"ഞാ", "തി", "ചൊ", "ബു", "വ്യാ", "വെ", "ശ"},
			"longDateFormat", js.Obj(
				"LT", "A h:mm -നു",
				"LTS", "A h:mm:ss -നു",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY, A h:mm -നു",
				"LLLL", "dddd, D MMMM YYYY, A h:mm -നു",
			),
			"calendar", js.Obj(
				"sameDay", "[ഇന്ന്] LT",
				"nextDay", "[നാളെ] LT",
				"nextWeek", "dddd, LT",
				"lastDay", "[ഇന്നലെ] LT",
				"lastWeek", "[കഴിഞ്ഞ] dddd, LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s കഴിഞ്ഞ്",
				"past", "%s മുൻപ്",
				"s", "അൽപ നിമിഷങ്ങൾ",
				"ss", "%d സെക്കൻഡ്",
				"m", "ഒരു മിനിറ്റ്",
				"mm", "%d മിനിറ്റ്",
				"h", "ഒരു മണിക്കൂർ",
				"hh", "%d മണിക്കൂർ",
				"d", "ഒരു ദിവസം",
				"dd", "%d ദിവസം",
				"M", "ഒരു മാസം",
				"MM", "%d മാസം",
				"y", "ഒരു വർഷം",
				"yy", "%d വർഷം",
			),
			"meridiemParse", jsre.MustCompile("രാത്രി|രാവിലെ|ഉച്ച കഴിഞ്ഞ്|വൈകുന്നേരം|രാത്രി", "i"),
			"meridiemHour", mlMeridiemHour,
			"meridiem", mlMeridiem,
		))
	})
}

// mlMeridiemHour is the meridiemHour callback of ml.
func mlMeridiemHour(_ *moment.Locale, hour any, meridiem string) any {
	hour = noonToZero(hour)
	if (meridiem == "രാത്രി" && hourNumber(hour) >= 4) || meridiem == "ഉച്ച കഴിഞ്ഞ്" || meridiem == "വൈകുന്നേരം" {
		return plus12(hour)
	}
	return hour
}

// mlMeridiem is the meridiem callback of ml.
func mlMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	switch {
	case hour < 4:
		return "രാത്രി"
	case hour < 12:
		return "രാവിലെ"
	case hour < 17:
		return "ഉച്ച കഴിഞ്ഞ്"
	case hour < 20:
		return "വൈകുന്നേരം"
	}
	return "രാത്രി"
}
