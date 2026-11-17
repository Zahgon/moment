// moment.js locale configuration
// locale : Sindhi [sd]
// author : Narain Sagar : https://github.com/narainsagar

package locale

import (
	"github.com/moment/moment"
	"github.com/moment/moment/internal/js"
	"github.com/moment/moment/internal/jsre"
)

func init() {
	moment.RegisterLocale("sd", func() {
		moment.DefineLocale("sd", js.Obj(
			"months", []string{"جنوري", "فيبروري", "مارچ", "اپريل", "مئي", "جون", "جولاءِ", "آگسٽ", "سيپٽمبر", "آڪٽوبر", "نومبر", "ڊسمبر"},
			"monthsShort", []string{"جنوري", "فيبروري", "مارچ", "اپريل", "مئي", "جون", "جولاءِ", "آگسٽ", "سيپٽمبر", "آڪٽوبر", "نومبر", "ڊسمبر"},
			"weekdays", []string{"آچر", "سومر", "اڱارو", "اربع", "خميس", "جمع", "ڇنڇر"},
			"weekdaysShort", []string{"آچر", "سومر", "اڱارو", "اربع", "خميس", "جمع", "ڇنڇر"},
			"weekdaysMin", []string{"آچر", "سومر", "اڱارو", "اربع", "خميس", "جمع", "ڇنڇر"},
			"longDateFormat", js.Obj(
				"LT", "HH:mm",
				"LTS", "HH:mm:ss",
				"L", "DD/MM/YYYY",
				"LL", "D MMMM YYYY",
				"LLL", "D MMMM YYYY HH:mm",
				"LLLL", "dddd، D MMMM YYYY HH:mm",
			),
			"meridiemParse", jsre.MustCompile("صبح|شام", ""),
			"isPM", sdIsPM,
			"meridiem", sdMeridiem,
			"calendar", js.Obj(
				"sameDay", "[اڄ] LT",
				"nextDay", "[سڀاڻي] LT",
				"nextWeek", "dddd [اڳين هفتي تي] LT",
				"lastDay", "[ڪالهه] LT",
				"lastWeek", "[گزريل هفتي] dddd [تي] LT",
				"sameElse", "L",
			),
			"relativeTime", js.Obj(
				"future", "%s پوء",
				"past", "%s اڳ",
				"s", "چند سيڪنڊ",
				"ss", "%d سيڪنڊ",
				"m", "هڪ منٽ",
				"mm", "%d منٽ",
				"h", "هڪ ڪلاڪ",
				"hh", "%d ڪلاڪ",
				"d", "هڪ ڏينهن",
				"dd", "%d ڏينهن",
				"M", "هڪ مهينو",
				"MM", "%d مهينا",
				"y", "هڪ سال",
				"yy", "%d سال",
			),
			"preparse", sdPreparse,
			"postformat", sdPostformat,
			"week", js.Obj(
				"dow", 1.0,
				"doy", 4.0,
			),
		))
	})
}

// sdIsPM is the isPM callback of sd.
func sdIsPM(_ *moment.Locale, input string) bool { return input == "شام" }

// sdMeridiem is the meridiem callback of sd.
func sdMeridiem(_ *moment.Locale, hour, _ float64, _ bool) string {
	if hour < 12 {
		return "صبح"
	}
	return "شام"
}

var (
	sdPreparse   = dvPreparse
	sdPostformat = dvPostformat
)
